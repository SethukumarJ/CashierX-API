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
	routes.PUT("/update/:id", svc.UpdateAccount)
	routes.GET("/getall/:id", svc.GetAccounts)
	routes.DELETE("/delete/:id", svc.DeleteAccount)
	routes.GET("/balance/:id", svc.GetBalance)
	
}


func (svc *ServiceClient) CreateAccount(ctx *gin.Context) {
	routes.CreateAccount(ctx, svc.Client)
}
func (svc *ServiceClient) FindAccount(ctx *gin.Context) {
	routes.FindAccount(ctx, svc.Client)
}
func (svc *ServiceClient) UpdateAccount(ctx *gin.Context) {
	routes.UpdateAccount(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteAccount(ctx *gin.Context) {
	routes.DeleteAccount(ctx, svc.Client)
}
func (svc *ServiceClient) GetBalance(ctx *gin.Context) {
	routes.GetBalance(ctx, svc.Client)
}
func (svc *ServiceClient) GetAccounts(ctx *gin.Context) {
	routes.GetUserAccounts(ctx, svc.Client)
}