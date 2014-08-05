package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/opentarock/frontend-user-management/server/logutil"
	"github.com/opentarock/frontend-user-management/server/service"
	"github.com/opentarock/frontend-user-management/server/service/proto_user"
)

type Api struct {
	mux         *gin.Engine
	userService service.UserService
}

func New(userService service.UserService) *Api {
	api := &Api{
		mux:         gin.Default(),
		userService: userService,
	}
	api.mux.POST("/user", api.postUser)
	return api
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mux.ServeHTTP(w, r)
}

func (api *Api) postUser(c *gin.Context) {
	var user proto_user.User
	c.Bind(&user)
	redirectURI := c.Request.URL.Query().Get("redirect_uri")
	response, err := api.userService.RegisterUser(&user, redirectURI)
	if err != nil {
		logutil.Printf(c.Request, "Error registering user: %s", err)
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}
	if !response.GetValid() {
		c.JSON(http.StatusBadRequest, response)
	} else {
		c.Writer.Header().Set("Location", *response.RedirectUri)
		c.JSON(http.StatusCreated, response)
	}
}
