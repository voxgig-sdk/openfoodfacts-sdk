package core

type OpenfoodfactsError struct {
	IsOpenfoodfactsError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewOpenfoodfactsError(code string, msg string, ctx *Context) *OpenfoodfactsError {
	return &OpenfoodfactsError{
		IsOpenfoodfactsError: true,
		Sdk:              "Openfoodfacts",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *OpenfoodfactsError) Error() string {
	return e.Msg
}
