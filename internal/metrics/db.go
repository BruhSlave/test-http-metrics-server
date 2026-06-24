package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *metrics) DBStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		db := pingData{
			Status:  setStatus(HealthDBRandom),
			Service: "database",
		}

		if db.Status == "fail" {
			m.ErrorsTotal.Inc()
			m.RequestTotal.Inc()
		} else {
			m.RequestTotal.Inc()
		}

		if err := json.NewEncoder(w).Encode(db); err != nil {
			fmt.Println("Error in encoding to the stream: ", err)
		}
	}
}
