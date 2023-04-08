package account

import (
	"github.com/SethukumarJ/CashierX-API/pkg/auth"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	// a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/account")
	// routes.Use(a.AuthRequired)
	// routes.POST("/", svc.CreateProduct)
	// routes.GET("/:id", svc.FindOne)
}
