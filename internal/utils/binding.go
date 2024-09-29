package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "piyo-engine/api/v1"
)

type BindingError struct {
	DetailedError string `json:"detailed_error"`
}

func BindJSONOrBadRequest(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(
			ctx,
			http.StatusBadRequest,
			v1.ErrBadRequest,
			BindingError{
				DetailedError: err.Error(),
			},
		)
		return err
	}
	return nil
}
