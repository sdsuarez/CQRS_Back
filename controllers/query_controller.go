// controllers/query_controller.go
package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "proyect_modules/db"
    "strconv"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
    // Realiza la consulta de todas las tareas y devuelve los resultados
    tasks, err := db.GetAllTasks()
    if err != nil {
        http.Error(w, "Error al obtener todas las tareas", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de tarea no válido", http.StatusBadRequest)
        return
    }

    // Realiza la consulta de una tarea por ID y devuelve el resultado
    task, err := db.GetTaskByID(taskID)
    if err != nil {
        http.Error(w, "La tarea no existe", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}

func GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
    // Realiza la consulta de tareas completadas y devuelve los resultados
    completedTasks, err := db.GetCompletedTasks()
    if err != nil {
        http.Error(w, "Error al obtener tareas completadas", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(completedTasks)
}
