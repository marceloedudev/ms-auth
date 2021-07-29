package middleware

import (
	"fmt"
	"ms-auth/pkg/http_exception"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {

			if err := recover(); err != nil {
				switch err.(type) {
				case *http_exception.HttpException:
					{
						errorsTemplate := err.(*http_exception.HttpException)
						errorsTemplate.Error = http.StatusText(errorsTemplate.Status)
						errorsTemplate.Timestamp = time.Now()
						errorsTemplate.Path = c.Request.URL.Path

						c.JSON(errorsTemplate.Status, errorsTemplate)
						c.Abort()
						return
					}
				default:
					{
						c.JSON(http.StatusInternalServerError, &http_exception.HttpException{
							Message:   fmt.Sprintf("%v", err),
							Status:    http.StatusInternalServerError,
							Error:     http.StatusText(http.StatusInternalServerError),
							Timestamp: time.Now(),
							Path:      c.Request.URL.Path,
						})
						c.Abort()
						return
					}
				}
			}
		}()

		c.Next()

	}
}
