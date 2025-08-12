package application_test

import (
	"testing"

	"github.com/kalinskilk/arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"
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
