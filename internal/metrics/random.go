package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var Random = []status{
	StatusOk,
	StatusWarning,
	StatusError,
}

func (m *metrics) RandomStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		random := pingData{
			Status: setStatus(Random),
		}

		if random.Status == "error" {
			m.ErrorsTotal.Inc()
			m.RequestTotal.Inc()
		} else {
			m.RequestTotal.Inc()
		}

		if err := json.NewEncoder(w).Encode(random); err != nil {
			fmt.Println("Error in encoding to the stream", err)
		}
	}
}
