
package models

type TaskQueryResult struct {
    ID          int    `json:"id" gorm:"column:tasks"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
