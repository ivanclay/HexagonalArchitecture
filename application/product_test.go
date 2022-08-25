package application_test

import (
	"github.com/codeedu/go-exagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enebled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}