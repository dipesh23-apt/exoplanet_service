// models/exoplanet.go
package models

type Exoplanet struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Distance    float64 `json:"distance"`
    Radius      float64 `json:"radius"`
    Type        string  `json:"type"`
}

type GasGiant struct {
    Exoplanet
}

type Terrestrial struct {
    Exoplanet
    Mass float64 `json:"mass"`
}
