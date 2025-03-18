package http

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/core/usecase"
)

type WeatherHandler struct {
	usecase *usecase.WeatherUseCase
}

func NewWeatherHandler(uc *usecase.WeatherUseCase) *WeatherHandler {
	return &WeatherHandler{usecase: uc}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.usecase.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
