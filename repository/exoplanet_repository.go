// repository/exoplanet_repository.go
package repository

import (
    "exoplanet-service/models"
    "sync"
)

var (
    exoplanets = make(map[string]interface{})
    mu         sync.Mutex
)

func AddExoplanet(exoplanet interface{}) {
    mu.Lock()
    defer mu.Unlock()
    id := exoplanet.(models.Exoplanet).ID
    exoplanets[id] = exoplanet
}

func GetExoplanets() map[string]interface{} {
    return exoplanets
}

func GetExoplanetByID(id string) (interface{}, bool) {
    exoplanet, exists := exoplanets[id]
    return exoplanet, exists
}

func UpdateExoplanet(id string, exoplanet interface{}) {
    mu.Lock()
    defer mu.Unlock()
    exoplanets[id] = exoplanet
}

func DeleteExoplanet(id string) {
    mu.Lock()
    defer mu.Unlock()
    delete(exoplanets, id)
}
