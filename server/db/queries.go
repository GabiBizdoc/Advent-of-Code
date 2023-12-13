package db

import (
	"context"
	"fmt"
	"time"
)

type RequestLog struct {
	ID            int
	IP            string
	CreatedAt     time.Time
	Day           int
	Part          int
	CorrectAnswer bool
	Valid         bool
	Message       string
}

func (r *RequestLog) Insert(ctx context.Context) error {
	const q = `INSERT INTO requests (ip, created_at, day, part, correct_answer, valid) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	err := db.QueryRowContext(ctx, q, r.IP, r.CreatedAt, r.Day, r.Part, r.CorrectAnswer, r.Valid).Scan(&r.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *RequestLog) Update(ctx context.Context) error {
	if r.ID == 0 {
		err := fmt.Errorf("can't update entity with id 0")
		fmt.Println(err)
		return err
	}
	const q = `UPDATE requests SET ip = $1, created_at = $2, day = $3, part = $4, correct_answer = $5, valid = $6
                WHERE id = $7;`
	_, err := db.ExecContext(ctx, q, r.IP, r.CreatedAt, r.Day, r.Part, r.CorrectAnswer, r.Valid, r.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
