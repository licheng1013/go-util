package util

import "time"

// TaskUtil 定时任务工具
type TaskUtil struct {
}

func NewTaskUtil() *TaskUtil {
	return &TaskUtil{}
}

// RunTask 延迟执行任务,1秒后执行
func (u TaskUtil) RunTask(f func(), d time.Duration) {
	time.AfterFunc(d, f)
}

// LoopTask 休眠执行任务 -> 你需要在外部启用一个线程来运行他
func (u TaskUtil) LoopTask(f func(), d time.Duration) {
	for {
		f()
		time.Sleep(d)
	}
}

// CustomLoopTask 自定义次数循环任务 -> 你需要在外部启用一个线程来运行他
func (u TaskUtil) CustomLoopTask(f func(), d time.Duration, count int) {
	for i := 0; i < count; i++ {
		f()
		time.Sleep(d)
	}
}
