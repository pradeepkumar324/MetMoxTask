package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/metmoxtask/responses"
	"example.com/metmoxtask/utils"
)

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		lattitude := r.URL.Query().Get("lat")
		longitude := r.URL.Query().Get("long")

		url := fmt.Sprintf("%s?latitude=%s&longitude=%s&current_weather=true", utils.BaseUrl, lattitude, longitude)
		response := responses.GetCurrentWeatherResponse{}
		err := utils.GetJson(url, &response)
		if err != nil {
			log.Printf("Error in getting Current Weather, error:%v", err)
			json.NewEncoder(w).Encode(responses.ErrorResponse{
				Code:    500,
				Message: "Internal Server error",
				Detail:  fmt.Sprintf("%v", err),
			})
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}
