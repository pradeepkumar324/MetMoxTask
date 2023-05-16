package responses

type CurrentWeather struct {
	Temperature   float32 `json:"temperature"`
	WindSpeed     float32 `json:"windspeed"`
	WindDirection float32 `json:"winddirection"`
	WeatheCcode   int     `json:"weathecode"`
	IsDay         int     `json:"is_day"`
	Time          string  `json:"time"`
}

type GetCurrentWeatherResponse struct {
	Latitude             float32        `json:"latitude"`
	Longitude            float32        `json:"longitude"`
	GenerationTimeMs     float32        `json:"generation_time_ms"`
	UTCOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             string         `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float32        `json:"elevation"`
	CurrentWeather       CurrentWeather `json:"current_weather"`
}
