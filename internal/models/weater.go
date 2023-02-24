package models

type Weather struct {
	City        string
	Temperature Temperature
	Description WeatherDescription
}

type WeatherDescription struct {
	ID   int64
	Main string
	Text string
}

type Temperature struct {
	Current          float64
	CurrentFeelsLike float64
	Min              float64
	Max              float64
}
