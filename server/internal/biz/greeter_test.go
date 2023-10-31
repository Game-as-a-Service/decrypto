package biz_test

import (
	"server/internal/biz"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGreeterUsecase(t *testing.T) {
	// mockRepo := &mocks.GreeterRepo{}
	// mockLogger := &log.Helper{}


	greeterUsecase := biz.NewGreeterUsecase(nil, nil)

	assert.NotNil(t, greeterUsecase)
}
