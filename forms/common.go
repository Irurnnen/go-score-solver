package forms

type (
	HTTPError struct {
		Message string `json:"message" example:"An error has occurred"`
		Error   string `json:"error" example:"unkown error"`
	}
)
