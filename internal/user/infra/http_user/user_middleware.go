package http_user

import (
	"ms-auth/internal/user/domain"
	"ms-auth/internal/user/usecase"
	"ms-auth/pkg/http_exception"

	"github.com/gin-gonic/gin"
)

func RequestLimit(us usecase.UseCaseUser) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := us.RequestLimit(domain.FormatKeyRequestLimit(c.ClientIP()), 20, 60)

		if err != nil {

			panic(http_exception.BadRequestException(err))

		}

		c.Next()

	}
}
