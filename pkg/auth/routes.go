package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/SethukumarJ/CashierX-API/pkg/auth/routes"
    "github.com/SethukumarJ/CashierX-API/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }
    authMiddleware := InitAuthMiddleware(svc)

    // Create a new group for the auth routes that require authentication
    authRoutes := r.Group("/auth")
    authRoutes.Use(authMiddleware.RefreshTokenMiddleware)
    authRoutes.POST("/token-refresh",svc.TokenRefresh)

    routes := r.Group("/auth")
    routes.POST("/register", svc.Register)
    routes.POST("/login", svc.Login)
    
    userRoutes := r.Group("/user")
    userRoutes.DELETE("/delete/:id",svc.DeleteUser)
  
    return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
    routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
    routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
    routes.DeleteUser(ctx, svc.Client)
}
func (svc *ServiceClient) TokenRefresh(ctx *gin.Context) {
    routes.TokenRefresh(ctx, svc.Client)
}