package models

type Task struct {
    ID          int    `json:"id" bson:"id"`
    Title       string `json:"title" bson:"title"`
    Description string `json:"description" bson:"description"`
    DueDate     string `json:"due_date" bson:"due_date"`
    Status      string `json:"status" bson:"status"`
}