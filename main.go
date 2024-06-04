// main.go
package main

import (
    "exoplanet-service/controllers"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/exoplanets", controllers.AddExoplanet).Methods("POST")
    r.HandleFunc("/exoplanets", controllers.ListExoplanets).Methods("GET")
    r.HandleFunc("/exoplanets/{id}", controllers.GetExoplanetByID).Methods("GET")
    r.HandleFunc("/exoplanets/{id}", controllers.UpdateExoplanet).Methods("PUT")
    r.HandleFunc("/exoplanets/{id}", controllers.DeleteExoplanet).Methods("DELETE")
    r.HandleFunc("/exoplanets/{id}/estimate-fuel", controllers.EstimateFuel).Methods("GET")

    http.ListenAndServe(":8080", r)
    fmt.Println("Listening on Port :8080")
}
