package http_client

import (
	"ms-auth/internal/client/usecase"
	"ms-auth/pkg/http_exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tokenUser godoc
// @Summary Register Client
// @Description
// @Tags clients
// @Accept json
// @Param input body ClientRegisterRequest false "Create Client"
// @Success 200 {object} domain.Client
// @Router /client/register [post]
func registerClient(us usecase.UseCaseClient) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input ClientRegisterRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		client, err := us.RegisterClient(c, input.ClientID, input.ClientSecret, input.ManagingID, input.UserID)

		if err != nil {
			panic(http_exception.BadRequestException(err))
		}

		c.JSON(http.StatusOK, client)

	}

}
