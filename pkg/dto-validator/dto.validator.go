package dtovalidator

import (
	errorhandler "github.com/DaniaLD/EyeOn/pkg/error-handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func BindBodyAndValidate(c *gin.Context, dto interface{}) bool {
	if !checkNilValue(c, dto) {
		return false
	}

	if err := c.ShouldBindBodyWithJSON(dto); err != nil {
		formattedErrors := errorhandler.GlobalErrorWrapper(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"payload":    gin.H{"errors": formattedErrors},
		})
		return false
	}

	return true
}

func BindUriAndValidate(c *gin.Context, dto interface{}) bool {
	if !checkNilValue(c, dto) {
		return false
	}

	if err := c.ShouldBindUri(dto); err != nil {
		formattedErrors := errorhandler.GlobalErrorWrapper(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"payload":    gin.H{"errors": formattedErrors},
		})
		return false
	}

	return true
}

func BindQueryAndValidate(c *gin.Context, dto interface{}) bool {
	if !checkNilValue(c, dto) {
		return false
	}

	if err := c.ShouldBindQuery(dto); err != nil {
		formattedErrors := errorhandler.GlobalErrorWrapper(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"payload":    gin.H{"errors": formattedErrors},
		})
		return false
	}

	return true
}

func checkNilValue(c *gin.Context, dto interface{}) bool {
	// Use reflection to make sure we're dealing with a pointer
	val := reflect.ValueOf(dto)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"payload":    gin.H{"errors": "Invalid data type provided. Must be a non-nil value."},
		})
		return false
	}
	return true
}
