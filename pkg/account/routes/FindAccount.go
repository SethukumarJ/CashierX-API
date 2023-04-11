package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SethukumarJ/CashierX-API/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

// @Summary Find user by id
// @ID Find account by id
// @Tags Account-Service
// @Produce json
// @Security BearerAuth 
// @param id path string true "Find account by id"
// @Success 200 {object} pb.FindAccountResponse{}
// @Failure 422 {object} pb.FindAccountResponse{}
// @Failure 502 {object} pb.FindAccountResponse{}
// @Router /account/get/{id} [get]
func FindAccount(ctx *gin.Context, c pb.AccountServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.FindAccountResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}
	res, err := c.FindAccount(context.Background(), &pb.FindAccountRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), pb.FindAccountResponse{Error: fmt.Sprint(res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
