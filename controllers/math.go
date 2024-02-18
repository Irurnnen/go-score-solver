package controllers

import (
	"net/http"

	"github.com/Irurnnen/go-score/forms"
	"github.com/Irurnnen/go-score/models"
	"github.com/gin-gonic/gin"
)

type MathController struct{}

var mm models.Maths

// SolveAddition godoc
// @Summary     Solve subtraction (+)
// @Description Solve subtraction (+)
// @Tags        Math
// @Param       a  query integer true "Number A"
// @Param       b  query integer true "Number B"
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Answer
// @Failure     400 {object} forms.HTTPError
// @Failure     500 {object} forms.HTTPError
// @Router      /addition [get]
func (mc MathController) SolveAddition(ctx *gin.Context) {
	mathForm := new(forms.MathOperationForm)
	err := ctx.BindQuery(mathForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, forms.HTTPError{Message: "Provided data is invalid", Error: ""})
		return
	}

	answer, err := mm.SolveAddition(mathForm.NumA, mathForm.NumB)

	switch err {
	case nil:
		break
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, forms.HTTPError{Message: "Unkown error while getting addiction", Error: "unkown error"})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

// SolveAddiction godoc
// @Summary     Solve addiction (-)
// @Description Solve addiction (-)
// @Tags        Math
// @Param       a  query integer true "Number A"
// @Param       b  query integer true "Number B"
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Answer
// @Failure     400 {object} forms.HTTPError
// @Failure     500 {object} forms.HTTPError
// @Router      /subtraction [get]
func (mc MathController) SolveSubtraction(ctx *gin.Context) {
	mathForm := new(forms.MathOperationForm)
	err := ctx.BindQuery(mathForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, forms.HTTPError{Message: "Provided data is invalid", Error: ""})
		return
	}

	answer, err := mm.SolveSubstraction(mathForm.NumA, mathForm.NumB)

	switch err {
	case nil:
		break
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, forms.HTTPError{Message: "Unkown error while getting addiction", Error: "unkown error"})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

// SolveMultiplication godoc
// @Summary     Solve multiplication (*)
// @Description Solve multiplication (*)
// @Tags        Math
// @Param       a  query integer true "Number A"
// @Param       b  query integer true "Number B"
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Answer
// @Failure     400 {object} forms.HTTPError
// @Failure     500 {object} forms.HTTPError
// @Router      /multiplication [get]
func (mc MathController) SolveMultiplication(ctx *gin.Context) {
	mathForm := new(forms.MathOperationForm)
	err := ctx.BindQuery(mathForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, forms.HTTPError{Message: "Provided data is invalid", Error: ""})
		return
	}

	answer, err := mm.SolveMultiplication(mathForm.NumA, mathForm.NumB)

	switch err {
	case nil:
		break
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, forms.HTTPError{Message: "Unkown error while getting addiction", Error: "unkown error"})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

// SolveDivision godoc
// @Summary     Solve division (/)
// @Description Solve division (/)
// @Tags        Math
// @Param       a  query integer true "Number A"
// @Param       b  query integer true "Number B"
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Answer
// @Failure     400 {object} forms.HTTPError
// @Failure     500 {object} forms.HTTPError
// @Router      /division [get]
func (mc MathController) SolveDivision(ctx *gin.Context) {
	mathForm := new(forms.MathOperationForm)
	err := ctx.BindQuery(mathForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, forms.HTTPError{Message: "Provided data is invalid", Error: ""})
		return
	}

	answer, err := mm.SolveDivision(mathForm.NumA, mathForm.NumB)

	switch err {
	case nil:
		break
	case models.ErrDivisionByZero:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, forms.HTTPError{Message: "Division by zero is illegal", Error: "Division by zero"})
		return
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, forms.HTTPError{Message: "Unkown error while getting addiction", Error: "unkown error"})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}
