// factory/exoplanet_factory.go
package factory

import (
    "errors"
    "exoplanet-service/models"
)

func CreateExoplanet(exoType string, id, name, description string, distance, radius, mass float64) (interface{}, error) {
    switch exoType {
    case "GasGiant":
        return &models.GasGiant{
            Exoplanet: models.Exoplanet{
                ID:          id,
                Name:        name,
                Description: description,
                Distance:    distance,
                Radius:      radius,
                Type:        exoType,
            },
        }, nil
    case "Terrestrial":
        return &models.Terrestrial{
            Exoplanet: models.Exoplanet{
                ID:          id,
                Name:        name,
                Description: description,
                Distance:    distance,
                Radius:      radius,
                Type:        exoType,
            },
            Mass: mass,
        }, nil
    default:
        return nil, errors.New("invalid exoplanet type")
    }
}
