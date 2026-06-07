package random

import (
	"fmt"
	"testing"
	"time"
)

func Test_rand7(t *testing.T) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(1) * time.Nanosecond)
		fmt.Println(rand10())
	}
}
