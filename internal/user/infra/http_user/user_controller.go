package http_user

import (
	"errors"
	clientUsecase "ms-auth/internal/client/usecase"
	"ms-auth/internal/user/domain"
	userUsecase "ms-auth/internal/user/usecase"
	"ms-auth/pkg/http_exception"
	"ms-auth/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tokenUser godoc
// @Summary Token user
// @Description
// @Tags users
// @Accept json
// @Param login body UserTokenLoginRequest false "Data User Login"
// @Param refresh token body UserTokenRefreshRequest false "Data Refresh Token"
// @Success 200 {object} domain.TokenResponse
// @Router /user/tokens [post]
func tokenUser(us userUsecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		var inputType UserTokenRequest

		if err := c.ShouldBindJSON(&inputType); err != nil {
			panic(err)
		}

		grantTypes := map[string]func(){
			"password": func() {
				inputLogin := &UserTokenLoginRequest{
					GrantType: inputType.GrantType,
					Username:  inputType.Username,
					Password:  inputType.Password,
				}

				info, err := us.LoginUser(c, inputLogin.Username, inputLogin.Password)

				if err != nil {
					panic(http_exception.BadRequestException(err))
				}

				c.JSON(http.StatusOK, info)
			},
			"refresh_token": func() {
				inputRefresh := &UserTokenRefreshRequest{
					GrantType:    inputType.GrantType,
					RefreshToken: inputType.RefreshToken,
				}

				info, err := us.RefreshUserToken(c, inputRefresh.RefreshToken)

				if err != nil {
					panic(http_exception.BadRequestException(err))
				}

				c.JSON(http.StatusOK, info)
			},
		}

		if grantTypes[inputType.GrantType] == nil {
			panic(http_exception.BadRequestException(errors.New("missing or unknown grant type")))
		}

		grantTypes[inputType.GrantType]()

	}

}

// revokeUser godoc
// @Summary Revoke user
// @Description
// @Tags users
// @Accept json
// @Param input body UserRevokeTokenRequest false "Data"
// @Success 200
// @Router /user/revoke [post]
func revokeUser(us userUsecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input UserRevokeTokenRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		err := us.RevokeUser(input.Token, input.GrantType)

		if err != nil {
			panic(http_exception.BadRequestException(err))
		}

		c.Status(http.StatusOK)

	}

}

// authorizeUser godoc
// @Summary Authorize User
// @Description
// @Tags users
// @Accept json
// @Param input body UserInfoTokenRequest false "Authorize User"
// @Produce json
// @Success 200 {object} domain.UserDataToken
// @Router /user/authorize [post]
func authorizeUser(us userUsecase.UseCaseUser, cs clientUsecase.UseCaseClient) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input UserInfoTokenRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		client, err := cs.VerifyClient(c, input.ClientID, input.ClientSecret)

		if err != nil {
			panic(http_exception.BadRequestException(err))
		}

		if client == nil {
			panic(http_exception.UnauthorizedExeception(errors.New("Invalid client")))
		}

		user, err := us.InfoAccessToken(input.AccessToken)

		if err != nil {
			panic(http_exception.UnauthorizedExeception(err))
		}

		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")

		c.JSON(http.StatusOK, user)

	}

}

// registerUser godoc
// @Summary Register user
// @Description
// @Tags users
// @Accept json
// @Param input body UserRegisterRequest false "Create User"
// @Success 200 {object} domain.TokenResponse
// @Router /user/register [post]
func registerUser(us userUsecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		var input UserRegisterRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			panic(err)
		}

		info, err := us.RegisterUser(c, input.Username, input.FirstName, input.LastName, input.Password, input.Email)

		if err != nil {
			panic(http_exception.BadRequestException(err))
		}

		c.JSON(http.StatusOK, info)

	}

}

func infoUser(us userUsecase.UseCaseUser) gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := utils.ExtractTokenFromRequest(c.Request)

		err := domain.IsValidToken(tokenString)

		if err != nil {
			panic(http_exception.UnauthorizedExeception(err))
		}

		user, err := us.InfoAccessToken(tokenString)

		if err != nil {
			panic(http_exception.UnauthorizedExeception(err))
		}

		c.JSON(http.StatusOK, user)

	}

}
