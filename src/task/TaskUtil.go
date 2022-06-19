package task

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/task"
)

var (
	iTask = task.NewTask("iTask", "0 12 * * * *", func(ctx context.Context) error {
		fmt.Print("task init")
		return nil
	})
)

func AddTask(newTask task.Task) {
	task.AddTask(newTask.Taskname, &newTask)
}

func RunTask(ctx context.Context) {
	iTask.Run(ctx)
}
func StopTask() {
	task.StopTask()
}
