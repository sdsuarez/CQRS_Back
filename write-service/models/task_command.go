
package models

type CreateTaskCommand struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

type UpdateTaskCommand struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`

}
type CreateTaskCommentCommand struct {
    TaskID      uint      `json:"task_id"`
    CommentText string    `json:"comment_text"`
}

type UpdateTaskCommentCommand struct {
    ID          uint   `json:"id"`
    CommentText string `json:"comment_text"`
}

type CreateTaskAttachmentCommand struct {
    TaskID    uint   `json:"task_id"`
    FileName  string `json:"file_name"`
    FileURL   string `json:"file_url"`
    FileType  string `json:"file_type"`
    FileSize  uint64 `json:"file_size"`
}

type UpdateTaskAttachmentCommand struct {
    ID        uint   `json:"id"`
    FileName  string `json:"file_name"`
    FileURL   string `json:"file_url"`
    FileType  string `json:"file_type"`
    FileSize  uint64 `json:"file_size"`
}
