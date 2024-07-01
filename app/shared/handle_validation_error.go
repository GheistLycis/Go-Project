package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
HandleValidationError builds a map from the field errors and invokes c.JSON(http.StatusBadRequest) with it
*/
func HandleValidationError(e validator.ValidationErrors, c *gin.Context) {
	fieldErrors := map[string]string{}

	for _, err := range e {
		fieldErrors[err.Field()] = err.Tag()
	}

	c.JSON(http.StatusBadRequest, fieldErrors)
}
