package account

import (
	"github.com/SethukumarJ/CashierX-API/pkg/account/routes"
	"github.com/SethukumarJ/CashierX-API/pkg/auth"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

   

	routes := r.Group("/account")
	routes.Use(a.AuthRequired)
	routes.POST("/create", svc.CreateAccount)
	routes.GET("/get/:id", svc.FindAccount )
	
}


func (svc *ServiceClient) CreateAccount(ctx *gin.Context) {
	routes.CreateAccount(ctx, svc.Client)
}
func (svc *ServiceClient) FindAccount(ctx *gin.Context) {
	routes.FindAccount(ctx, svc.Client)
}
