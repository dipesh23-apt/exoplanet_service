// controllers/exoplanet_controller.go
package controllers

import (
    "encoding/json"
    "exoplanet-service/services"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
    var data struct {
        ID          string  `json:"id"`
        Name        string  `json:"name"`
        Description string  `json:"description"`
        Distance    float64 `json:"distance"`
        Radius      float64 `json:"radius"`
        Mass        float64 `json:"mass,omitempty"`
        Type        string  `json:"type"`
    }
    _ = json.NewDecoder(r.Body).Decode(&data)
    err := services.AddExoplanet(data.ID, data.Name, data.Description, data.Type, data.Distance, data.Radius, data.Mass)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
    exoplanets := services.ListExoplanets()
    json.NewEncoder(w).Encode(exoplanets)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    exoplanet, err := services.GetExoplanetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var data struct {
        Name        string  `json:"name"`
        Description string  `json:"description"`
        Distance    float64 `json:"distance"`
        Radius      float64 `json:"radius"`
        Mass        float64 `json:"mass,omitempty"`
        Type        string  `json:"type"`
    }
    _ = json.NewDecoder(r.Body).Decode(&data)
    err := services.UpdateExoplanet(id, data.Name, data.Description, data.Type, data.Distance, data.Radius, data.Mass)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    services.DeleteExoplanet(id)
    w.WriteHeader(http.StatusNoContent)
}

func EstimateFuel(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    crewCapacity, _ := strconv.Atoi(r.URL.Query().Get("crew_capacity"))
    fuel, err := services.EstimateFuel(id, crewCapacity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
}
