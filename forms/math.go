package forms

type (
	MathOperationForm struct {
		NumA int64 `form:"a" binding:"required,numeric"`
		NumB int64 `form:"b" binding:"required,numeric"`
	}
)
