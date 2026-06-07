package random

import (
	"math/rand"
	"time"
)

/* 470. 用 Rand7() 实现 Rand10() */
func rand7() int {
	time.Sleep(time.Duration(1) * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func rand10() int {
	var m int
	for {
		m = (rand7()-1)*7 + rand7() - 1
		if m <= 39 {
			break
		}
	}
	return m%10 + 1
}
