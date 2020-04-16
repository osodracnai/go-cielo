package test

import (
	"github.com/vasconcelosvcd/go-cielo"
	"os"
	"testing"
)

func TestZeroAuth(t *testing.T){
	c,err := cielo.NewClient(os.Getenv("CIELO_MERCHANT_ID"),os.Getenv("CIELO_MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err!= nil{
		println(err.Error())
	}
	response := &cielo.ZeroAuthResponse{}
	card := cielo.ZeroAuthCard{
		CardType:       cielo.Credit,
		CardNumber:     "111111111111111",
		Holder:         "TEST HOLDER",
		ExpirationDate: "11/2026",
		SecurityCode:   "123",
		Brand:          cielo.MasterCard,
	}
	response,err =c.MakeZeroAuth(&card)


	if err!=nil{
		t.Error("Zero auth failed " + err.Error())
	}
	if response == nil{
		print(err.Error())
	}

}
