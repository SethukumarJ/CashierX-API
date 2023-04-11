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

// @Summary Get Balance
// @Description Get Balance
// @Tags Account-Service
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Account ID"
// @Success 200 {object} pb.GetBalanceResponse{}
// @Failure 400 {object} pb.GetBalanceResponse{}
// @Failure 500 {object} pb.GetBalanceResponse{}
// @Router /account/balance/{id} [get]
// GetBalance is a function to get the balance of the account
func GetBalance(ctx *gin.Context, c pb.AccountServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.GetBalanceResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}
	res, err := c.GetBalance(context.Background(), &pb.GetBalanceRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), pb.GetBalanceResponse{Error: fmt.Sprint(res.Error)})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}