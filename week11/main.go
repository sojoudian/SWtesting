package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const worldTimeAPIURL = "http://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct {
	Datetime string `json:"datetime"`
}

func getTorontoTime() (string, error) {
	resp, err := http.Get(worldTimeAPIURL)
	if err != nil {
		return "Error retriving data", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading data", err
	}
	var TimeInfo TimeInfo
	err = json.Unmarshal(body, &TimeInfo)
	if err != nil {
		return "Error parsing data", err
	}
	return TimeInfo.Datetime, nil

}

func torontoTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime, err := getTorontoTime()
	if err != nil {
		http.Error(w, "Error fetching Toronto time", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Toronto time is %s", torontoTime)
	resp := map[string]string{"curremt_Time_Toronto": torontoTime}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
func main() {
	http.HandleFunc("/api/torontotime", torontoTimeHandler)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
