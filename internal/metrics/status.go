package metrics

import (
	"math/rand/v2"
)

type status string

const (
	StatusOk status = "ok"
	StatusWarning status = "warning"
	StatusError status = "error"
	StatusFail status = "fail"
)

var HealthDBRandom = []status{
	StatusOk,
	StatusFail,
}

type pingData struct {
	Status  status `json:"status"`
	Service string `json:"service,omitempty"`
}

func setStatus(statusList []status) status {
	randomIndex := rand.IntN(len(statusList))
	return statusList[randomIndex]
}


