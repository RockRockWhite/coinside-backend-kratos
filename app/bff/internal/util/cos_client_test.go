package util

import (
	"context"
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewCOSClinet(t *testing.T) {
	c := NewCOSClinet(config.NewConfig())

	name := "unit_test"
	f := strings.NewReader("test")

	res, err := c.Object.Put(context.Background(), name, f, nil)
	fmt.Println(err)
	fmt.Println(res)

	assert.Nil(t, err)
}
