package main

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type TaskInfo struct {
	Priority int
	TaskID   int
	UserID   int
}

type TaskManager struct {
	Tasklist     treeset.Set
	TaskIDToInfo map[int]*TaskInfo
}

func TaskInfoComparer(a, b interface{}) int {
	ap, bp := a.(*TaskInfo), b.(*TaskInfo)
	if ap.Priority == bp.Priority {
		return ap.TaskID - bp.TaskID
	}
	return ap.Priority - bp.Priority
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		Tasklist:     *treeset.NewWith(TaskInfoComparer),
		TaskIDToInfo: make(map[int]*TaskInfo),
	}
	for _, arr := range tasks {
		tm.Add(arr[0], arr[1], arr[2])
	}
	return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	taskInfo := &TaskInfo{Priority: priority, TaskID: taskId, UserID: userId}
	tm.Tasklist.Add(taskInfo)
	tm.TaskIDToInfo[taskId] = taskInfo
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	info := &TaskInfo{Priority: newPriority, TaskID: taskId, UserID: tm.TaskIDToInfo[taskId].UserID}

	tm.Tasklist.Add(info)
	tm.TaskIDToInfo[taskId] = info
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.TaskIDToInfo, taskId)
}

func (tm *TaskManager) ExecTop() int {
	it := tm.Tasklist.Iterator()
	it.Last()
	for {
		info := it.Value().(*TaskInfo)
		if infoGot, _ := tm.TaskIDToInfo[info.TaskID]; info == infoGot {
			tm.Rmv(info.TaskID)
			return info.UserID
		}

		if !it.Prev() {
			return -1
		}
	}
}

func main() {
	obj := Constructor(tasks)
	obj.Add(userId, taskId, priority)
	obj.Edit(taskId, newPriority)
	obj.Rmv(taskId)
	param_4 := obj.ExecTop()
}
