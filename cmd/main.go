package main

import (
	"fmt"
	"log"

	"github.com/SethukumarJ/CashierX-API/pkg/auth"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println("authSvc", authSvc)
	// account.RegisterRoutes(r, &c, &authSvc)
	// transaction.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
