// En replicadb/replicadb.go
package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "proyect_modules/models"
)

var dbr *gorm.DB

// Inicializa la conexión a la base de datos réplica y crea la tabla si no existe
func InitReplicaDB(connectionString string) error {
    var err error
    dbr, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        return err
    }

    // AutoMigrate creará la tabla TaskQueryResult si no existe
    err = dbr.AutoMigrate(&models.TaskQueryResult{})
    if err != nil {
        return err
    }
    
    
    log.Println("Conexión exitosa a la base de datos sec y tabla creada si no existía")
    return nil
}


// GetTaskByID retrieves a task by its ID from the database
func GetTaskByID(id int) (models.TaskQueryResult, error) {
    var task models.TaskQueryResult
    result := dbr.First(&task, id)
    if result.Error != nil {
        return task, result.Error
    }
    return task, nil
}


// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]models.TaskQueryResult, error) {
    var tasks []models.TaskQueryResult
    result := dbr.Find(&tasks)
    if result.Error != nil {
        return nil, result.Error
    }
    return tasks, nil
}


// GetCompletedTasks retrieves all completed tasks from the database
func GetCompletedTasks() ([]models.TaskQueryResult, error) {
    var tasks []models.TaskQueryResult
    result := dbr.Where("completed = ?", true).Find(&tasks)
    if result.Error != nil {
        return nil, result.Error
    }
    return tasks, nil
}
// GetAllComments recupera todos los comentarios de la base de datos
func GetAllComments() ([]models.TaskQueryComment, error) {
    var comments []models.TaskQueryComment
    result := dbr.Find(&comments)
    if result.Error != nil {
        return nil, result.Error
    }
    return comments, nil
}

// GetCommentByID recupera un comentario por su ID de la base de datos
func GetCommentByID(id uint) (models.TaskQueryComment, error) {
    var comment models.TaskQueryComment
    result := dbr.First(&comment, id)
    if result.Error != nil {
        return comment, result.Error
    }
    return comment, nil
}

// GetAllAttachments recupera todos los archivos adjuntos de la base de datos
func GetAllAttachments() ([]models.TaskQueryAttachment, error) {
    var attachments []models.TaskQueryAttachment
    result := dbr.Find(&attachments)
    if result.Error != nil {
        return nil, result.Error
    }
    return attachments, nil
}

// GetAttachmentByID recupera un archivo adjunto por su ID de la base de datos
func GetAttachmentByID(id uint) (models.TaskQueryAttachment, error) {
    var attachment models.TaskQueryAttachment
    result := dbr.First(&attachment, id)
    if result.Error != nil {
        return attachment, result.Error
    }
    return attachment, nil
}