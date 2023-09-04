package controllers

import (
	"context"
	"fmt"

	"github.com/kishorens18/ecommerce/interfaces"
	"github.com/kishorens18/ecommerce/models"

	"github.com/kishorens18/ecommerce/proto"
	pro "github.com/kishorens18/ecommerce/proto"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	var r *ecommerce.CustomerDetails
	var address models.Address
	if r != nil {
		address = models.Address{
			Country: r.Country,
			Street1: r.Street1,
			Street2: r.Street2,
			City:    r.City,
			State:   r.State,
			Zip:     r.Zip,
		}
	}
	addresses := []models.Address{address}
	var req1 *ecommerce.ShippingAddress
	var shippingAddress models.ShippingAddress
	if req1 != nil {
		shippingAddress = models.ShippingAddress{
			Street1: req1.Street1,
			Street2: req1.Street2,
			City:    req1.City,
			State:   req1.State,
		}
	}
	shippingAddresses := []models.ShippingAddress{shippingAddress}
	fmt.Println(req.Firstname)
	dbCustomer := models.Customer{CustomerId: req.CustomerId, Firstname: req.Firstname, Lastname: req.Lastname, HashesAndSaltedPassword: req.HashesAndSaltedPassword, EmailVerified: req.EmailVerified, Address:addresses,ShippingAddress: shippingAddresses}
	result, err := CustomerService.CreateCustomer(&dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Customer_ID: result.Customer_id,
		}
		return responseCustomer, nil
	}
}
