package main

import (
	"fmt"
	"log"

	"github.com/SethukumarJ/CashierX-API/pkg/auth"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title CashierX API
// @version 1.0
// @description This is Money management project. You can visit the GitHub repository at https://github.com/SethukumarJ/CashierX-API-Gateway

// @contact.name API Support
// @contact.url sethukumarj.com
// @contact.email sethukumarj.76@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3005
// @BasePath /
// @query.collection.format multi
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println("authSvc", authSvc)
	// account.RegisterRoutes(r, &c, &authSvc)
	// transaction.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
