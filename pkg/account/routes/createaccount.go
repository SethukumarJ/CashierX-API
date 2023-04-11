package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/SethukumarJ/CashierX-API/cmd/docs"
	"github.com/SethukumarJ/CashierX-API/pkg/account/pb"
	a "github.com/SethukumarJ/CashierX-API/pkg/auth/pb"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func getAuthClient() (a.AuthServiceClient, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	authConn, err := grpc.Dial(conf.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authClient := a.NewAuthServiceClient(authConn)
	return authClient, nil
}

type CreateAccountRequestBody struct {
	Type    string `json:"type"`
	Balance int64  `json:"balance"`
}

// @Summary Create account
// @ID Create Account
// @Tags Account-Service
// @Produce json
// @Security BearerAuth
// @param CreateAccount body CreateAccountRequestBody{} true "Create Account"
// @Success 200 {object} pb.CreateAccountResponse{}
// @Failure 422 {object} pb.CreateAccountResponse{}
// @Failure 502 {object} pb.CreateAccountResponse{}
// @Router /account/create [post]
func CreateAccount(ctx *gin.Context, c pb.AccountServiceClient) {

	fmt.Println("Create account product called in api")
	body := CreateAccountRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, err := strconv.Atoi(ctx.Writer.Header().Get("userId"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("userId header not provided2"))
		return
	}
	authClient, err := getAuthClient()
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.New("could not connect with auth client"))
        return
    }

    authRes, authErr := authClient.FindUser(context.Background(), &a.FindUserRequest{
        Id: int64(userId),
    })

    if authErr != nil {
		fmt.Println(authErr, "authErr")
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("failed to retrieve user information"))
        return
    }

	userID := int64(userId)
	fmt.Println(authErr, "after authErr")
	res, err := c.CreateAccount(context.Background(), &pb.CreateAccountRequest{
    	AccountHolder: authRes.Data.UserName,
    	Type:          body.Type,
    	Balance:       body.Balance,
    	UserId:        userID,
})

    if err != nil {
		fmt.Println(authErr, "fuccked")
        ctx.AbortWithStatusJSON(http.StatusOK, res.Error)
        return
    }

    ctx.JSON(int(res.Status), &res)
}