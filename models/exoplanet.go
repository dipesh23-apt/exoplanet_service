// models/exoplanet.go
package models

type Exoplanet struct {
    ID          string        `json:"id"`
    Name        string        `json:"name" validate:"required"`
    Description string        `json:"description" validate:"required"`
    Distance    int           `json:"distance" validate:"required,gte=10,lte=1000"`
    Radius      float64       `json:"radius" validate:"required,gt=0.1,lt=10"`
    Type        string        `json:"type" validate:"required"`
}


type GasGiant struct {
    Exoplanet
}

type Terrestrial struct {
    Exoplanet
    Mass        float64       `json:"mass,omitempty" validate:"omitempty,gt=0.1,lt=10"`
}
