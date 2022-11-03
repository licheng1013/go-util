package test

import (
	"gitee.com/licheng1013/go-util/common"
	"log"
	"testing"
	"time"
)

func TestTaskUtil(t *testing.T) {
	common.TaskUtil.RunTask(PrintTest, 1*time.Second)

	go func() {
		common.TaskUtil.CustomLoopTask(PrintTest, 1*time.Second, 1)
	}()

	go func() {
		common.TaskUtil.LoopTask(PrintTest, 1*time.Second)
	}()
	time.Sleep(2 * time.Second)
}

func PrintTest() {
	log.Println("HelloWorld")
}
