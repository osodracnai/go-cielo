package test

import (
	"github.com/vasconcelosvcd/go-cielo"
	"os"
	"testing"
)

func TestTokenize(t *testing.T){
	c,err := cielo.NewClient(os.Getenv("CIELO_MERCHANT_ID"),os.Getenv("CIELO_MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err!= nil{
		println(err.Error())
	}
	cc:= cielo.CreditCard{
		CardNumber:     "5247712516640978",
		CustomerName:   "Tester Name",
		Holder:         "Teste Holder",
		ExpirationDate: "11/2026",
		SaveCard:       true,
		Brand:          "Master",
	}
	_,err =c.CreateTokenizeCard(&cc)


	if err!=nil{
		print(err.Error())
	}
	if  len(cc.CardToken)<=0 {
		t.Error("Não foi gerado o token do cartao \n "+ err.Error())
	}
}

func TestGetTokenized(t *testing.T){
	c,err := cielo.NewClient(os.Getenv("MERCHANT_ID"),os.Getenv("MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err!= nil{
		println(err.Error())
	}
	cc:= cielo.CreditCard{
		CardNumber:     "5247712516640978",
		CustomerName:   "Tester Name",
		Holder:         "Teste Holder",
		ExpirationDate: "11/2026",
		SaveCard:       true,
		Brand:          "Master",
	}
	_,err =c.CreateTokenizeCard(&cc)


	if err!=nil{
		t.Error(err.Error())
	}
	if  len(cc.CardToken)<=0 {
		t.Error("Não foi gerado o token do cartao \n "+ err.Error())
	}
	getCard,err:= c.GetTokenizeCard(cc.CardToken)
	if err!=nil{
		t.Error(err.Error())
	}
	if  len(getCard.CardNumber)<=0 {
		t.Error("Não foi gerado o token do cartao \n "+ err.Error())
	}
}