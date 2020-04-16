package cielo

import (
	"fmt"
)

//MakeZeroAuth returns if card is able to do transactions
//Endpoint POST /1/zeroauth
func (c *Client) MakeZeroAuth(card *ZeroAuthCard) (*ZeroAuthResponse, error) {
	//buf := bytes.NewBuffer([]byte(""))
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Environment.APIUrl, "/1/zeroauth"), card)
	response := &ZeroAuthResponse{}
	if err != nil {
		return response, err
	}
	err = c.Send(req, response)
	return response, err
}

