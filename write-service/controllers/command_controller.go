package controllers

import (
    "net/http"
    "encoding/json"
    "proyect_modules/models"
    "proyect_modules/db"
    "strconv"
    "github.com/gorilla/mux"
    "time"
    "log"
    "io/ioutil"

)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var cmd models.CreateTaskCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de creación
    if cmd.Title == "" || cmd.Description == "" {
        http.Error(w, "El título y la descripción son obligatorios", http.StatusBadRequest)
        return
    }

    task := models.TaskQueryResult{
        Title:       cmd.Title,
        Description: cmd.Description,
        Completed:   false, // Por defecto, una nueva tarea se crea como no completada
    }

    err := db.CreateTask(task)
    if err != nil {
        http.Error(w, "Error al crear la tarea", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
        log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
        return
    }

    // Log para verificar el contenido del cuerpo de la solicitud
    log.Printf("Cuerpo de la solicitud: %s", body)

    var cmd models.UpdateTaskCommand
    err = json.Unmarshal(body, &cmd)
    if err != nil {
        http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
        log.Printf("Error al decodificar la solicitud JSON: %v", err)
        return
    }

    // Log para verificar el contenido de cmd
    log.Printf("Solicitud recibida: %+v", cmd)

    // Valida y procesa el comando de actualización
    if cmd.Title == "" || cmd.Description == "" {
        http.Error(w, "El título y la descripción son obligatorios", http.StatusBadRequest)
        return
    }

    // Convierte cmd.ID a uint
    taskID := uint(cmd.ID)

    // Crear una estructura TaskQueryResult para actualizar la tarea
    task := models.TaskQueryResult{
        ID:          taskID,
        Title:       cmd.Title,
        Description: cmd.Description,
        Completed:   cmd.Completed,
    }

    // Actualizar la tarea en la base de datos
    err = db.UpdateTask(task)
    if err != nil {
        http.Error(w, "Error al actualizar la tarea", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}



func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
        return
    }

    // Valida y procesa el comando de eliminación
    err = db.DeleteTask(taskID)
    if err != nil {
        http.Error(w, "Error al eliminar la tarea", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func CreateTaskComment(w http.ResponseWriter, r *http.Request) {
    var cmd models.CreateTaskCommentCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de creación de comentario
    if cmd.TaskID == 0 || cmd.CommentText == "" {
        http.Error(w, "Task ID y texto del comentario son obligatorios", http.StatusBadRequest)
        return
    }

    comment := models.TaskQueryComment{
        CommentText: cmd.CommentText,
        CreatedAt:   time.Now(),
        Author:      "Admin", // Puedes especificar el autor aquí
    }

    // Crea el comentario en la base de datos
    err := db.CreateTaskComment(comment)
    if err != nil {
        http.Error(w, "Error al crear el comentario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}




func CreateTaskAttachment(w http.ResponseWriter, r *http.Request) {
    var cmd models.CreateTaskAttachmentCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de creación de adjunto
    if cmd.TaskID == 0 || cmd.FileName == "" || cmd.FileURL == "" {
        http.Error(w, "Task ID, nombre de archivo y URL son obligatorios", http.StatusBadRequest)
        return
    }

    attachment := models.TaskQueryAttachment{
        FileName:  cmd.FileName,
        FileURL:   cmd.FileURL,
        FileType:  cmd.FileType,
        FileSize:  cmd.FileSize,
        UploadDate: time.Now(),
    }

    err := db.CreateTaskAttachment(attachment)
    if err != nil {
        http.Error(w, "Error al crear el adjunto", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    
}
// Actualiza un comentario de tarea
func UpdateTaskComment(w http.ResponseWriter, r *http.Request) {
    var cmd models.UpdateTaskCommentCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de actualización de comentario
    if cmd.ID == 0 || cmd.CommentText == "" {
        http.Error(w, "ID de comentario y texto del comentario son obligatorios", http.StatusBadRequest)
        return
    }

    comment := models.TaskQueryComment{
        ID:          cmd.ID,
        CommentText: cmd.CommentText,
    }

    err := db.UpdateTaskComment(comment)
    if err != nil {
        http.Error(w, "Error al actualizar el comentario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// Elimina un comentario de tarea
func DeleteTaskComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    commentID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de comentario no válido", http.StatusBadRequest)
        return
    }

    // Valida y procesa el comando de eliminación de comentario
    err = db.DeleteTaskComment(uint(commentID))
    if err != nil {
        http.Error(w, "Error al eliminar el comentario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// Actualiza un archivo adjunto de tarea
func UpdateTaskAttachment(w http.ResponseWriter, r *http.Request) {
    var cmd models.UpdateTaskAttachmentCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de actualización de adjunto
    if cmd.ID == 0 || cmd.FileName == "" || cmd.FileURL == "" {
        http.Error(w, "ID de adjunto, nombre de archivo y URL son obligatorios", http.StatusBadRequest)
        return
    }

    attachment := models.TaskQueryAttachment{
        FileName:  cmd.FileName,
        FileURL:   cmd.FileURL,
        FileType:  cmd.FileType,
        FileSize:  cmd.FileSize,
    }

    err := db.UpdateTaskAttachment(attachment)
    if err != nil {
        http.Error(w, "Error al actualizar el adjunto", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// Elimina un archivo adjunto de tarea
func DeleteTaskAttachment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    attachmentID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de adjunto no válido", http.StatusBadRequest)
        return
    }

    // Valida y procesa el comando de eliminación de adjunto
    err = db.DeleteTaskAttachment(uint(attachmentID))
    if err != nil {
        http.Error(w, "Error al eliminar el adjunto", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
