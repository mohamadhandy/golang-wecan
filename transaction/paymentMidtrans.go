package transaction

import (
	"kitabisa/user"
	"os"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type midtransService struct {
}

type MidtransService interface {
	GetPaymentURL(Transaction, user.User) (string, error)
}

func NewMidTransService() *midtransService {
	return &midtransService{}
}

func (ms *midtransService) GetPaymentURL(tr Transaction, us user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: us.Email,
			FName: us.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(tr.ID),
			GrossAmt: int64(tr.Amount),
		},
	}
	snapToken, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapToken.RedirectURL, nil
}
