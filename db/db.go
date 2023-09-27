package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	"proyect_modules/models"
)

var db *gorm.DB

// Inicializa la conexión a la base de datos y crea la tabla si no existe
func InitDB(connectionString string) error {
    var err error
    db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        return err
    }

    // AutoMigrate creará la tabla TaskQueryResult si no existe
    err = db.AutoMigrate(&models.TaskQueryResult{})
    if err != nil {
        return err
    }

    log.Println("Conexión exitosa a la base de datos y tabla creada si no existía")
    return nil
}


// CreateTask creates a new task in the database
func CreateTask(task models.TaskQueryResult) error {
    result := db.Create(&task)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// GetTaskByID retrieves a task by its ID from the database
func GetTaskByID(id int) (models.TaskQueryResult, error) {
    var task models.TaskQueryResult
    result := db.First(&task, id)
    if result.Error != nil {
        return task, result.Error
    }
    return task, nil
}

// UpdateTask updates an existing task in the database
func UpdateTask(task models.TaskQueryResult) error {
    result := db.Save(&task)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteTask deletes a task by its ID from the database
func DeleteTask(id int) error {
    result := db.Delete(&models.TaskQueryResult{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]models.TaskQueryResult, error) {
    var tasks []models.TaskQueryResult
    result := db.Find(&tasks)
    if result.Error != nil {
        return nil, result.Error
    }
    return tasks, nil
}


// GetCompletedTasks retrieves all completed tasks from the database
func GetCompletedTasks() ([]models.TaskQueryResult, error) {
    var tasks []models.TaskQueryResult
    result := db.Where("completed = ?", true).Find(&tasks)
    if result.Error != nil {
        return nil, result.Error
    }
    return tasks, nil
}
