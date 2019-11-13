package handler

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateDate return Year Month Day Hour Minitues Seconds
func GenerateDate() string {
	const layout = "2006-01-02 15:04:05"
	t := time.Now()
	s := ""
	s = t.Format(layout)
	return s
}

// GenerateID return UUID
func GenerateID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	uu := u.String()
	return uu
}
