// services/exoplanet_service.go
package services

import (
    "exoplanet-service/factory"
    "exoplanet-service/models"
    "exoplanet-service/repository"
    "errors"
    // "fmt"
)

func AddExoplanet(id, name, description, exoType string, distance, radius, mass float64) error {
    exoplanet, err := factory.CreateExoplanet(exoType, id, name, description, distance, radius, mass)
    if err != nil {
        return err
    }
    repository.AddExoplanet(exoplanet)
    return nil
}

func ListExoplanets() map[string]interface{} {
    return repository.GetExoplanets()
}

func GetExoplanetByID(id string) (interface{}, error) {
    exoplanet, exists := repository.GetExoplanetByID(id)
    if !exists {
        return nil, errors.New("exoplanet not found")
    }
    return exoplanet, nil
}

func UpdateExoplanet(id, name, description, exoType string, distance, radius, mass float64) error {
    exoplanet, err := factory.CreateExoplanet(exoType, id, name, description, distance, radius, mass)
    if err != nil {
        return err
    }
    repository.UpdateExoplanet(id, exoplanet)
    return nil
}

func DeleteExoplanet(id string) {
    repository.DeleteExoplanet(id)
}

func EstimateFuel(id string, crewCapacity int) (float64, error) {
    exoplanet, exists := repository.GetExoplanetByID(id)
    if !exists {
        return 0, errors.New("exoplanet not found")
    }

    distance := exoplanet.(models.Exoplanet).Distance
    radius := exoplanet.(models.Exoplanet).Radius
    var gravity float64

    switch exo := exoplanet.(type) {
    case *models.GasGiant:
        gravity = 0.5 / (radius * radius)
    case *models.Terrestrial:
        gravity = exo.Mass / (radius * radius)
    }

    fuel := (distance / (gravity * gravity)) * float64(crewCapacity)
    return fuel, nil
}
