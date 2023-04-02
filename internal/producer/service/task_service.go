package service

import (
	"go-scrapy/proto/task"
)

type TaskService struct {
}

// NewTask 创建新任务
func (s *TaskService) NewTask() {
	// implement the logic of creating new task
}

// TaskPeek 查询任务状态 & 进度
func (s *TaskService) TaskPeek(req *task.TaskPeekRequest) (resp *task.TaskPeekResponse, err error) {
	// implement the logic of peeking task
	return resp, err
}
