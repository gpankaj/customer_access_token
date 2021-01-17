package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/storage_access_tokken/src/domain/access_token"
	"github.com/gpankaj/storage_access_tokken/src/services/access_token_service"
	"github.com/gpankaj/go-utils/rest_errors_package"

	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.Service
}
//Handler needs access token service to work.
func NewHandler(service access_token_service.Service) AccessTokenHandler  {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context) {
	access_token_id := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := handler.service.GetById(access_token_id)

	if err!= nil {

		ctx.JSON(err.Code,err)
	}

	ctx.JSON(http.StatusOK,accessToken)
}


func (handler *accessTokenHandler) Create(ctx *gin.Context) {
	var request access_token.AccessTokenRequest

	if err:= ctx.ShouldBindJSON(&request); err!= nil {
		restError := rest_errors_package.NewBadRequestError("Invalid JSON body passed")
		ctx.JSON(restError.Code, restError)
		return
	}
	at, err:= handler.service.Create(request)

	if err!= nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}
