package models

// TaskItem - объект задачи
type TaskItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

// TaskItemSlice - массив задач
type TaskItemSlice []TaskItem
