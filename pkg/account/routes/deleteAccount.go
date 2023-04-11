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

// @Summary Delete account
// @ID Delete account
// @Tags Account-Service
// @Produce json
// @Security BearerAuth
// @param id path string true "Id of the account to be deleted"
// @Success 200 {object} pb.DeleteAccountResponse{}
// @Failure 422 {object} pb.DeleteAccountResponse{}
// @Failure 502 {object} pb.DeleteAccountResponse{}
// @Router /account/delete/{id} [delete]
// DeleteUser deletes a user with the given ID.
func DeleteAccount(ctx *gin.Context, c pb.AccountServiceClient) {
	paramsID := ctx.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, pb.DeleteAccountResponse{Error: fmt.Sprint(errors.New("id not found"))})
		return
	}

	res, err := c.DeleteAccount(context.Background(), &pb.DeleteAccountRequest{Id: int64(id)})
	if err != nil {
		ctx.AbortWithStatusJSON(int(res.Status), res.Error)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
