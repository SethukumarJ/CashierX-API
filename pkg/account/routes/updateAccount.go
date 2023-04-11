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

type UpdateAccountRequestBody struct {
	Type    string `json:"type"`
	Balance int64  `json:"balance"`
}
// @Summary Update account
// @ID Update Account
// @Tags Account-Service
// @Produce json
// @Security BearerAuth
// @Param id path int true "Account ID"
// @param UpdateAccount body UpdateAccountRequestBody{} true "Create Account"
// @Success 200 {object} pb.UpdateAccountResponse{}
// @Failure 422 {object} pb.UpdateAccountResponse{}
// @Failure 502 {object} pb.UpdateAccountResponse{}
// @Router /account/update/{id} [put]
func UpdateAccount(ctx *gin.Context, c pb.AccountServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.FindAccountResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}
	body := UpdateAccountRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}


	res, err := c.UpdateAccount(context.Background(), &pb.UpdateAccountRequest{
		Id:      int64(id),
		Type:    body.Type,
		Balance: body.Balance,
	})

    if err != nil {
		fmt.Println(err, "fuccked")
        ctx.AbortWithStatusJSON(http.StatusExpectationFailed, res.Error)
        return
    }

    ctx.JSON(int(res.Status), &res)
}
