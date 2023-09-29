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
    err = db.AutoMigrate(&models.TaskQueryComment{})
    if err != nil {
        return err
    }
    err = db.AutoMigrate(&models.TaskQueryAttachment{})
    if err != nil {
        return err
    }
    log.Println("Conexión exitosa a la base de datos prim y tabla creada si no existía")
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


func CreateTaskComment(comment models.TaskQueryComment) error {
    result := db.Create(&comment)
    if result.Error != nil {
        return result.Error
    }

    

    return nil
}


// UpdateTaskComment actualiza un comentario existente en la base de datos
func UpdateTaskComment(comment models.TaskQueryComment) error {
    result := db.Save(&comment)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteTaskComment elimina un comentario por su ID de la base de datos
func DeleteTaskComment(id uint) error {
    result := db.Delete(&models.TaskQueryComment{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// CreateTaskAttachment crea un nuevo archivo adjunto en la base de datos
func CreateTaskAttachment(attachment models.TaskQueryAttachment) error {
    result := db.Create(&attachment)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// UpdateTaskAttachment actualiza un archivo adjunto existente en la base de datos
func UpdateTaskAttachment(attachment models.TaskQueryAttachment) error {
    result := db.Save(&attachment)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// DeleteTaskAttachment elimina un archivo adjunto por su ID de la base de datos
func DeleteTaskAttachment(id uint) error {
    result := db.Delete(&models.TaskQueryAttachment{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


