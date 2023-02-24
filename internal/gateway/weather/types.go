package weather

type weather struct {
	Description []weatherDescription `json:"weather"`
	Temperature temperature          `json:"main"`
	City        string               `json:"name"`
}

type weatherDescription struct {
	ID   int64  `json:"id"`
	Main string `json:"main"`
	Text string `json:"description"`
}

type temperature struct {
	Current          float64 `json:"temp"`
	CurrentFeelsLike float64 `json:"feels_like"`
	Min              float64 `json:"temp_min"`
	Max              float64 `json:"temp_max"`
}
