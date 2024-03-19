package solver

import (
	env "aoc/server/config"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func buildCommandArgs(day, part int) (string, []string) {
	if env.Config.IsDev {
		return "go", []string{"run", "cmd/server/main.go", "-day=" + strconv.Itoa(day), "-part=" + strconv.Itoa(part)}
	}
	return "/app/data/main", []string{"-day=" + strconv.Itoa(day), "-part=" + strconv.Itoa(part)}
}

func execute(timeout time.Duration, day, part int, fileContent string) (string, error) {
	start := time.Now()
	defer log.Info("execute done: ", time.Since(start))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmdName, cmdArgs := buildCommandArgs(day, part)
	log.Debug("running: ", cmdName, cmdArgs)
	cmd := exec.CommandContext(ctx, cmdName, cmdArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	solutionFinder := &SolutionFinder{}
	cmd.Stdout = solutionFinder

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Info(err)
	}

	if err = cmd.Start(); err != nil {
		log.Error("An error occurred when starting child process: ", err)
		return "", err
	}

	writeString, err := io.WriteString(stdin, fileContent)
	if err != nil {
		return "", err
	}

	log.Info("len(stdin) = ", writeString)
	err = stdin.Close()
	if err != nil {
		log.Error(err)
		return "", err
	}

	done := make(chan error)
	go func() {
		err = cmd.Wait()
		if err != nil {
			log.Error(err)
			done <- err
		} else {
			done <- nil
		}
	}()

	select {
	case <-ctx.Done():
		log.Info("failed to resolve in time. killing the child")
		err := cmd.Process.Kill()
		if err != nil {
			log.Error(err)
		}
		err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			log.Error(err)
			return "", newDDLError(timeout)
		}

	case err := <-done:
		if err != nil {
			log.Error("challenge failed to resolve: ", err)
		}
		log.Info("challenge resolved ", cmd.ProcessState.Success())
		if cmd.ProcessState.Success() {
			if solutionFinder.ResultLine == "" {
				return "", fmt.Errorf("reslut line not found")
			}
			return solutionFinder.ResultLine, nil
		} else {
			return "", fmt.Errorf("failed to solve challenge")
		}
	}
	return "", fmt.Errorf("failed to solve challenge")
}

func solveProblem(timeout time.Duration, day, part int, fileContent string, r *ExecuteResult) *ExecuteResult {
	// we need this because cmd.Process.Kill() may panic if the process hasn't been started yet
	defer func(x *ExecuteResult) {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
			//fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
			log.Errorf("failed to solve problem in: %s, total_time: %s", timeout.String(), time.Since(x.Start).String())
			x.Err = errors.New("something went wrong")
			//x.Err = newDDLError(timeout)
		}
	}(r)
	response, err := execute(timeout, day, part, fileContent)
	if err != nil {
		r.Err = err
		return r
	}
	r.ParseResponse(response)
	return r
}

func newDDLError(timeout time.Duration) error {
	return fmt.Errorf("failed to spawn worker and solve challenge in %s", timeout.String())
}

func SolveProblem(timeout time.Duration, day, part int, fileContent string) *ExecuteResult {
	start := time.Now()
	r := NewExecutorResult()

	solveProblem(timeout, day, part, fileContent, r)
	defer func() {
		r.RealTime = time.Since(start)
	}()
	if r == nil {
		log.Debug("solution is nil!")
		r := NewExecutorResult()
		if time.Since(start) > timeout {
			r.Err = newDDLError(timeout)
		} else {
			r.Err = errors.New("something went wrong")
		}
		return r
	}
	return r
}

type ExecuteResult struct {
	Start         time.Time
	RawResult     string
	fields        []string
	ExecutionTime string
	RealTime      time.Duration
	Solution      int
	Err           error
}

func NewExecutorResult() *ExecuteResult {
	return &ExecuteResult{Start: time.Now()}
}

func (r *ExecuteResult) ParseResponse(rawResult string) {
	r.Err = r.parseResponse(rawResult)
}

func (r *ExecuteResult) parseResponse(rawResult string) error {
	r.fields = strings.Fields(rawResult)

	if len(r.fields) < 4 {
		log.Info("rawResult: ", rawResult)
		return errors.New("failed to parse raw result")
	}

	if r.fields[0] != "result:" {
		return fmt.Errorf("faild to parse solution. expected `result:` found: %s", r.fields[0])
	}
	if r.fields[2] != "took:" {
		return fmt.Errorf("faild to parse solution. expected `took:` found: %s", r.fields[2])
	}
	solution, err := strconv.Atoi(r.fields[1])
	if err != nil {
		return err
	}
	r.Solution = solution
	r.ExecutionTime = r.fields[3]
	return nil
}
