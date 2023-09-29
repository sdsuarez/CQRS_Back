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
    err := db.InitDB("host=write-db port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatalf("Error al inicializar la base de datos principal: %v", err)
    }
/*
    err = db.InitReplicaDB("host=postgres-primary port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatalf("Error al inicializar la base de datos réplica: %v", err)
    }
*/

    // Grupo de rutas para operaciones de escritura de tareas (usando la base de datos principal)
    writeGroup := r.PathPrefix("/write").Subrouter()
    writeGroup.HandleFunc("/tasks/create", controllers.CreateTask).Methods("POST")
    writeGroup.HandleFunc("/tasks/update/{id}", controllers.UpdateTask).Methods("PUT")
    writeGroup.HandleFunc("/tasks/delete/{id}", controllers.DeleteTask).Methods("DELETE")

   
    // Grupo de rutas para operaciones de escritura de comentarios (usando la base de datos principal)
    writeGroup.HandleFunc("/comments/create", controllers.CreateTaskComment).Methods("POST")
    writeGroup.HandleFunc("/comments/update/{id}", controllers.UpdateTaskComment).Methods("PUT")
    writeGroup.HandleFunc("/comments/delete/{id}", controllers.DeleteTaskComment).Methods("DELETE")

   

    // Grupo de rutas para operaciones de escritura de archivos adjuntos (usando la base de datos principal)
    writeGroup.HandleFunc("/attachments/create", controllers.CreateTaskAttachment).Methods("POST")
    writeGroup.HandleFunc("/attachments/update/{id}", controllers.UpdateTaskAttachment).Methods("PUT")
    writeGroup.HandleFunc("/attachments/delete/{id}", controllers.DeleteTaskAttachment).Methods("DELETE")



    port := ":8080"
    log.Printf("Servidor en ejecución en el puerto %s", port)
    log.Fatal(http.ListenAndServe(port, r))
}
