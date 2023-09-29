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



func GetAllComments(w http.ResponseWriter, r *http.Request) {
    // Realiza la consulta de todos los comentarios y devuelve los resultados
    comments, err := db.GetAllComments()
    if err != nil {
        http.Error(w, "Error al obtener todos los comentarios", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(comments)
}

func GetCommentByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    commentID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de comentario no válido", http.StatusBadRequest)
        return
    }

    // Realiza la consulta de un comentario por ID y devuelve el resultado
    comment, err := db.GetCommentByID(uint(commentID))
    if err != nil {
        http.Error(w, "El comentario no existe", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(comment)
}

// Nuevas funciones para archivos adjuntos

func GetAllAttachments(w http.ResponseWriter, r *http.Request) {
    // Realiza la consulta de todos los archivos adjuntos y devuelve los resultados
    attachments, err := db.GetAllAttachments()
    if err != nil {
        http.Error(w, "Error al obtener todos los archivos adjuntos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(attachments)
}

func GetAttachmentByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    attachmentID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID de archivo adjunto no válido", http.StatusBadRequest)
        return
    }

    // Realiza la consulta de un archivo adjunto por ID y devuelve el resultado
    attachment, err := db.GetAttachmentByID(uint(attachmentID))
    if err != nil {
        http.Error(w, "El archivo adjunto no existe", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(attachment)
}
