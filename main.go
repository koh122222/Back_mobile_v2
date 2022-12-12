package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AutoGenerated struct {
	Current struct {
		TempC     int `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Day struct {
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
			} `json:"day"`
			Hour []struct {
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func workWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start work Weather")
	city := mux.Vars(r)["city"] //"moscow"
	fmt.Println("work with city '" + city + "'")
	resp, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=db90acd3651a4b82b96112114221112&q=" + city + "&days=3&lang=ru")
	if err != nil {
		fmt.Println("error get weather")
	}
	var rest_data AutoGenerated

	err = json.NewDecoder(resp.Body).Decode(&rest_data)
	if err != nil {
		fmt.Println(err)
	}

	//rest_data.Forecast.Forecastday[0].Hour = append(rest_data.Forecast.Forecastday[0].Hour[7:], rest_data.Forecast.Forecastday[0].Hour[9:]...)

	answerJson, err := json.Marshal(rest_data)

	fmt.Fprint(w, string(answerJson))
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func connectRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/city/{city}", workWeather).Methods("GET")
	http.ListenAndServe(":8082", router)
}

func main() {
	fmt.Println("Start main")
	connectRoutes()
}
