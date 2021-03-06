package cielo

import (
	"bytes"
	"fmt"
 "gopkg.in/go-playground/validator.v9"
)

//GetBinValidation returns card BIN information
//Endpoint GET /1/cardBin/{BIN}
func (c *Client) GetBinValidation(cardBin string) (*BINResponse, error) {
	buf := bytes.NewBuffer([]byte(""))
	validator.New().Var(cardBin,"len=6")
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.Environment.APIQueryURL, "/1/cardBin/", cardBin), buf)

	bResponse := &BINResponse{}

	if err != nil {
		return bResponse, err
	}

	err = c.Send(req, bResponse)
	return bResponse, err
}
