package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	saveTimeToDatabase(torontoTime)

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) {
	db, err := sql.Open("mysql", "username:password@/GO_API")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO time_table (time) VALUES (?)", time)
	if err != nil {
		panic(err)
	}
}
