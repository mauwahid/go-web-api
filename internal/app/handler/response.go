package handler

type DefaultResponse struct {
	Status
	Data interface{} `json:"data"`
}

type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	StatusSuccess = Status{"00", "Success"}
	StatusFailed  = Status{"99", "Failed"}
)
