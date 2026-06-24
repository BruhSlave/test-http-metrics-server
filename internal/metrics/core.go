package metrics


import (
	"fmt"
	"net/http"
)

func (m *metrics) Core() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.RequestTotal.Inc()
		if _, err := fmt.Fprintf(w, `Server is working
Avaliable endpoints:
1. /health
2. /db
3. /random
4. /metrics`); err != nil {
			fmt.Println("Error in / endpoint", err)
		}
	}
}
