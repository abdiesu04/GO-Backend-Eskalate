package data

import (
    "task_manager/models"
    "errors"
)

var tasks []models.Task = []models.Task{
    {ID: 1, Title: "Task 1", Description: "Description 1", DueDate: "2023-10-01", Status: "pending"},
    {ID: 2, Title: "Task 2", Description: "Description 2", DueDate: "2023-10-02", Status: "in-progress"},
    {ID: 3, Title: "Task 3", Description: "Description 3", DueDate: "2023-10-03", Status: "completed"},
}
var nextID int = 4

func GetTasks() []models.Task {
    return tasks
}

func GetTaskByID(id int) (*models.Task, error) {
    for _, task := range tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
    task.ID = nextID
    nextID++
    tasks = append(tasks, task)
    return task
}

func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            tasks[i].ID = id
            return &tasks[i], nil
        }
    }
    return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}