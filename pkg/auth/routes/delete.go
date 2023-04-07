package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SethukumarJ/CashierX-API/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.AuthServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	fmt.Println("id", id)
	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
