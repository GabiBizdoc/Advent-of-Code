package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Box[T any] struct {
	Key   string
	Value T
}

func NewBox[T any](key string, value T) Box[T] {
	return Box[T]{Key: key, Value: value}
}

func (x *Box[T]) SetUsage(usage string) {
	switch v := any(&x.Value).(type) {
	case *int:
		flag.IntVar(v, x.Key, *v, usage)
	case *string:
		flag.StringVar(v, x.Key, *v, usage)
	case *bool:
		flag.BoolVar(v, x.Key, *v, usage)
	default:
		panic(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (x *Box[T]) SetUsagef(usage string, a ...any) {
	x.SetUsage(fmt.Sprintf(usage, a...))
}

func PreviewArguments(c any) string {
	var sb strings.Builder
	sb.WriteString(os.Args[0])
	sb.WriteString(" ")

	v := reflect.ValueOf(c)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println(v.Kind(), v.Elem())
		panic("v is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sb.WriteString(fmt.Sprintf("--%s %v ",
			field.FieldByName("Key").Interface(),
			field.FieldByName("Value").Interface()),
		)
	}

	return sb.String()
}

type Config struct {
	OutputFile        Box[string]
	Lines             Box[int]
	Generators        Box[int]
	WriterChannelSize Box[int]
	Help              Box[bool]
	DryRun            Box[bool]
}

func NewConfig() *Config {
	return &Config{
		Lines:             NewBox("lines", 1_000_000_000),
		OutputFile:        NewBox("output", ""),
		Generators:        NewBox("generators", 10),
		WriterChannelSize: NewBox("writer-channel-size", 10),
		Help:              NewBox("help", false),
		DryRun:            NewBox("dry-run", false),
	}
}

func parseArgs() (*Config, error) {
	args := NewConfig()

	args.Help.SetUsage("Help!")
	args.DryRun.SetUsage("Dri run")
	args.Lines.SetUsagef("Number of goroutines used to generated data: --generators %d", args.Lines.Value)
	args.OutputFile.SetUsage("Output file. Skip for stdout: Example --output ./weather_stations_data.csv")
	args.Generators.SetUsagef("Number of goroutines used to generated data: --generators %d", args.Generators.Value)
	args.WriterChannelSize.SetUsagef("Number of chunks buffered. Must be bigger than 0: --writer-channel-size %d", args.WriterChannelSize.Value)

	flag.Parse()

	if args.Help.Value {
		flag.Usage()
		os.Exit(0)
	}
	if args.Lines.Value <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	var err error
	if args.OutputFile.Value != "" {
		args.OutputFile.Value, err = filepath.Abs(args.OutputFile.Value)
		if err != nil {
			return nil, err
		}
	}

	if args.Generators.Value <= 0 {
		const defaultValue = 10
		fmt.Printf("generators must be at least 1. Changing it to %d. it was %d\n",
			defaultValue, args.Generators.Value)
		args.Generators.Value = defaultValue
	}

	if args.Lines.Value <= 0 {
		fmt.Println("lines must be at least 1")
		os.Exit(1)
	}

	if args.WriterChannelSize.Value < 0 {
		const defaultValue = 5
		fmt.Printf("writer-channel-size must be at least 1. Changing it to %d. it was %d\n",
			defaultValue, args.WriterChannelSize.Value)
		args.WriterChannelSize.Value = defaultValue
	}

	return args, nil
}
