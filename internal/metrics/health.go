package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *metrics) HealthStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		health := pingData{
			Status: setStatus(HealthDbRandom),
		}

		if health.Status == "fail" {
			m.ErrorsTotal.Inc()
			m.RequestTotal.Inc()
		} else {
			m.RequestTotal.Inc()
		}

		if err := json.NewEncoder(w).Encode(health); err != nil {
			fmt.Println("Error in encoding to the stream: ", err)
		}
	}
}
