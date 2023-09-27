
package controllers

import (
    "net/http"
    "encoding/json"
    "proyect_modules/models"
    "proyect_modules/db"
    "strconv"
    "github.com/gorilla/mux"
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
    var cmd models.UpdateTaskCommand
    _ = json.NewDecoder(r.Body).Decode(&cmd)

    // Valida y procesa el comando de actualización
    if cmd.Title == "" || cmd.Description == "" {
        http.Error(w, "El título y la descripción son obligatorios", http.StatusBadRequest)
        return
    }

    existingTask, err := db.GetTaskByID(cmd.ID)
    if err != nil {
        http.Error(w, "La tarea no existe", http.StatusNotFound)
        return
    }

    existingTask.Title = cmd.Title
    existingTask.Description = cmd.Description
    existingTask.Completed = cmd.Completed

    err = db.UpdateTask(existingTask)
    if err != nil {
        http.Error(w, "Error al actualizar la tarea", http.StatusInternalServerError)
        return
    }

    
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
