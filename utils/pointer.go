package utils

import "time"

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func NowTH() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}
