package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vasconcelosvcd/go-cielo"
	"os"
	"testing"
)

func TestGetBINValidation(t *testing.T){
	c,err := cielo.NewClient(os.Getenv("CIELO_MERCHANT_ID"),os.Getenv("CIELO_MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err!= nil{
		println(err.Error())
	}
	BIN:= "444458"
	bResponse,err :=c.GetBinValidation(BIN)


	assert.Equal(t,nil,err)
	assert.Equal(t,"VISA",bResponse.Provider)
}
