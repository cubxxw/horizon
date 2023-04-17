package terminal

import (
	"fmt"
	"net/http"

	"github.com/horizoncd/horizon/pkg/server/route"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register routes.
func (api *API) RegisterRoute(engine *gin.Engine) {
	coreGroup := engine.Group("/apis/core/v2")
	coreRoutes := route.Routes{
		{
			Method:      http.MethodGet,
			Pattern:     fmt.Sprintf("/clusters/:%v/shell", _clusterIDParam),
			HandlerFunc: api.CreateShell,
		},
	}
	route.RegisterRoutes(coreGroup, coreRoutes)
}
