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

// @Summary Get accounts of  user
// @ID Get accounts
// @Tags Account-Service
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} pb.GetAccountsResponse{}
// @Failure 422 {object} pb.GetAccountsResponse{}
// @Failure 502 {object} pb.GetAccountsResponse{}
// @Router /account/getall/{id} [get]
func GetUserAccounts(ctx *gin.Context, c pb.AccountServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.FindAccountResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}
	res, err := c.GetAccounts(context.Background(), &pb.GetAccountsRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}