package models

import (
    "time"
)

type TaskQueryResult struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
 
}

type TaskQueryComment struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    CommentText string    `json:"comment_text"`
    CreatedAt   time.Time `json:"created_at"`
    Author      string    `json:"author"`
}

type TaskQueryAttachment struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    FileName  string    `json:"file_name"`
    FileURL   string    `json:"file_url"`
    FileType  string    `json:"file_type"`
    FileSize  uint64    `json:"file_size"`
    UploadDate time.Time `json:"upload_date"`
}
