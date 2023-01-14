package model

type Clipboard struct {
	Content string `json:"content"`
}

type Result struct {
	Code      int       `json:"code"`
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Clipboard Clipboard `json:"clipboard"`
}
