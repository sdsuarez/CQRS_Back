package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "proyect_modules/controllers"
    "proyect_modules/db"
)

func main() {
    r := mux.NewRouter()

    // Define las conexiones a la base de datos principal y la réplica
    err := db.InitReplicaDB("host=write-db port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatalf("Error al inicializar la base de datos principal: %v", err)
    }
    // Grupo de rutas para operaciones de lectura de tareas (usando la base de datos réplica)
    readGroup := r.PathPrefix("/read").Subrouter()
    readGroup.HandleFunc("/tasks/get_all", controllers.GetAllTasks).Methods("GET")
    readGroup.HandleFunc("/tasks/get_by_id/{id}", controllers.GetTaskByID).Methods("GET")
    readGroup.HandleFunc("/tasks/get_completed", controllers.GetCompletedTasks).Methods("GET")


    // Grupo de rutas para operaciones de lectura de comentarios (usando la base de datos réplica)
    readGroup.HandleFunc("/comments/get_all", controllers.GetAllComments).Methods("GET")
    readGroup.HandleFunc("/comments/get_by_id/{id}", controllers.GetCommentByID).Methods("GET")


    // Grupo de rutas para operaciones de lectura de archivos adjuntos (usando la base de datos réplica)
    readGroup.HandleFunc("/attachments/get_all", controllers.GetAllAttachments).Methods("GET")
    readGroup.HandleFunc("/attachments/get_by_id/{id}", controllers.GetAttachmentByID).Methods("GET")

    port := ":8081"
    log.Printf("Servidor en ejecución en el puerto %s", port)
    log.Fatal(http.ListenAndServe(port, r))
}
