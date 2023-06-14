package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
)

// this is the incoming JSON payload's schema
type EmployeeData struct {
	Name     	string `json:"name"`
	Age     	int    `json:"age"`
	JobTitle	string `json:"jobTitle"`
	BadgeNumber int    `json:"badgeNumber"`
}


func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=data.csv")

	var employees []EmployeeData
	err := json.NewDecoder(r.Body).Decode(&employees)
	if err != nil {
		http.Error(w, "Could not parse the JSON payload.", http.StatusBadRequest)
		return
	}

	writer := csv.NewWriter(w)
	writer.Write([]string{"Name", "Age", "Job Title", "Badge Number"})

	for _, employee := range employees {
		row := []string{employee.Name, fmt.Sprintf("%d", employee.Age), employee.JobTitle, fmt.Sprintf("%d", employee.BadgeNumber)}
		err := writer.Write(row)
		if err != nil {
			http.Error(w, "Could not write CSV data.", http.StatusInternalServerError)
			return
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		http.Error(w, "Could not complete writing CSV data", http.StatusInternalServerError)
		return
	}
}