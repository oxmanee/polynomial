package model

type CalculateRequest struct {
	X int `json:"x" validate:"required"`
	Y int `json:"y" validate:"required"`
	Z int `json:"z" validate:"required"`
}

type CalculateResponse struct {
	Message string `json:"message,omitempty"`
	Dataset []int  `json:"dataset,omitempty"`
}
