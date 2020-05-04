package test

import (
	"github.com/vasconcelosvcd/go-cielo"
	"os"
	"testing"
)

func TestDoUndoSale(t *testing.T) {
	c, err := cielo.NewClient(os.Getenv("CIELO_MERCHANT_ID"), os.Getenv("CIELO_MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err != nil {
		println(err.Error())
	}
	sale:= cielo.Sale{
		MerchantOrderID: "2014111903",
		Customer: &cielo.Customer{
			Name:         "Comprador crédito autenticação",
			Email:        "",
			BirthDate:    "",
			Identity:     "12345678912",
			IdentityType: "cpf",
		},
		Payment: &cielo.Payment{
			Installments:     1,
			CreditCard: &cielo.CreditCard{
				CardNumber:     "1234123412341231",
				Holder:         "Teste Holder",
				ExpirationDate: "12/2030",
				SecurityCode:   "123",
				SaveCard:       false,
				Brand:          "Visa",
			},
			Tid:                 "",
			ProofOfSale:         "",
			AuthorizationCode:   "",
			SoftDescriptor:      "",
			ReturnURL:           "https://www.iupay.com.br/instantpayment/",
			Provider:            "Cielo",
			PaymentID:           "",
			Type:                "CreditCard",
			Amount:              15700,
			ReceiveDate:         "",
			CapturedAmount:      0,
			CapturedDate:        "",
			Currency:            "",
			Country:             "",
			ReturnCode:          "",
			ReturnMessage:       "",
			Status:              0,
			Links:               nil,
			ExtraDataCollection: nil,
			ExpirationDate:      "",
			URL:                 "",
			Number:              "",
			BarCodeNumber:       "",
			DigitableLine:       "",
			Address:             "",
			ExternalAuthentication: &cielo.ExternalAuthentication{
				Cavv:        "123456789",
				Xid:         "987654321",
				Eci:         "5",
				Version:     "",
				ReferenceId: "",
			},
		},
	}
	var saleResult  *cielo.Sale
	saleResult,err=c.DoSale(&sale)
	if err != nil{
		println(err.Error())
		return
	}
	println(saleResult.Payment.AuthorizationCode)
	UndoSaleResponse,err :=c.UndoSale(saleResult)

	if err != nil{
		println(err.Error())
		return
	}
	println(UndoSaleResponse.AuthorizationCode)
}