package metrics

import (
	"math/rand/v2"
)

type Status string

const (
	StatusOk Status = "ok"
	StatusWarning Status = "warning"
	StatusError Status = "error"
	StatusFail Status = "fail"
)

var HealthDbRandom = []Status{
	StatusOk,
	StatusFail,
}

type pingData struct {
	Status  Status `json:"status"`
	Service string `json:"service,omitempty"`
}

func setStatus(statusList []Status) Status {
	randomIndex := rand.IntN(len(statusList))
	return statusList[randomIndex]
}


