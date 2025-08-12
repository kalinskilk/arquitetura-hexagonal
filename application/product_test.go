package application_test

import (
	"testing"

	"github.com/kalinskilk/arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"

	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	p := application.Product{Price: 10}
	p.Name = "hELLO"
	p.Status = application.DISABLED
	p.Price = 10

	err := p.Enable()
	require.Nil(t, err)

	p.Price = 0
	err = p.Enable()
	require.Equal(t, "price must be greater than zero", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	p := application.Product{Price: 10}
	p.Name = "hELLO"
	p.Status = application.DISABLED
	p.Price = 0

	err := p.Disable()
	require.Nil(t, err)

	p.Price = 1
	err = p.Disable()
	require.Equal(t, "price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	p := application.Product{Price: 10}
	p.Id = uuid.NewV4().String()
	p.Name = "hELLO"
	p.Status = application.DISABLED
	p.Price = 10

	_, err := p.IsValid()
	require.Nil(t, err)

	p.Status = "invalid"
	_, err = p.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	p.Status = application.ENABLED
	_, err = p.IsValid()
	require.Nil(t, err)

	p.Price = -10
	_, err = p.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
