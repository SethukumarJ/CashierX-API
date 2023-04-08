package account

import (
	"fmt"

	"github.com/SethukumarJ/CashierX-API/pkg/account/pb"
	"github.com/SethukumarJ/CashierX-API/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.AccountServiceClient
}

func InitServiceClient(c *config.Config) pb.AccountServiceClient {
    // using WithInsecure() because no SSL running
    cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewAccountServiceClient(cc)
}