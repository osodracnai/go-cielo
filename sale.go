package cielo

import (
	"errors"
	"fmt"
)

//Sale returns the payment/sale authorization
//Endpoint POST /1/sales
func (c *Client) DoSale(sale *Sale) (*Sale, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Environment.APIUrl, "/1/sales/"), sale)

	salePayed := &Sale{}

	if err != nil {
		return salePayed, err
	}

	err = c.Send(req, salePayed)
	return salePayed, nil
}

//UndoSale returns the payment/sale authorization
//Endpoint POST /1/sales
func (c *Client) UndoSale(sale *Sale) (*UndoSaleResponse, error) {
	if sale.Payment.PaymentID ==""{
		return nil,errors.New("missing PaymentId parameter")
	}
	if sale.Payment.Amount == 0{
		return nil,errors.New("missing Amount parameter")
	}
	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s%s%s", c.Environment.APIUrl, "/1/sales/",sale.Payment.PaymentID,"/void"), sale)

	salePayed := &UndoSaleResponse{}

	if err != nil {
		return salePayed, err
	}

	err = c.Send(req, salePayed)
	return salePayed, nil
}
