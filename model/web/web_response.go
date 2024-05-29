package web

type WebResponse struct {
	Code   int16  `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
