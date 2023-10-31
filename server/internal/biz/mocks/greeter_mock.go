package mocks

import (
	"context"
	"server/internal/biz"
)

type GreeterRepo struct {}

func(gp *GreeterRepo) Save(context.Context, *biz.Greeter) (*biz.Greeter, error) {
	return nil, nil
} 

func(gp *GreeterRepo) Update(context.Context, *biz.Greeter) (*biz.Greeter, error) {
	return nil, nil
}

func(gp *GreeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return &biz.Greeter{}, nil
}

func(gp *GreeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func(gp *GreeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}

