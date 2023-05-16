package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/metmoxtask/handler"
)

func main() {

	http.HandleFunc("/currentweather", handler.GetCurrentWeather)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}
