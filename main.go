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

    err := db.InitDB("host=postgres port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatalf("Error al inicializar la base de datos: %v", err)
    }


    // Rutas para comandos
    r.HandleFunc("/commands/create", controllers.CreateTask).Methods("POST")
    r.HandleFunc("/commands/update/{id}", controllers.UpdateTask).Methods("PUT")
    r.HandleFunc("/commands/delete/{id}", controllers.DeleteTask).Methods("DELETE")

    // Rutas para consultas
    r.HandleFunc("/queries/get_all", controllers.GetAllTasks).Methods("GET")
    r.HandleFunc("/queries/get_by_id/{id}", controllers.GetTaskByID).Methods("GET")
    r.HandleFunc("/queries/get_completed", controllers.GetCompletedTasks).Methods("GET")

    port := ":8080"
    log.Printf("Servidor en ejecución en el puerto %s", port)
    log.Fatal(http.ListenAndServe(port, r))
}
