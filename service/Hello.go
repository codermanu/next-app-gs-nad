package service

import (
	"github.com/CloudImpl-Inc/next-coder-sdk/polycode"
	"portal/register/model"
)

func Hello(ctx polycode.ServiceContext, req model.HelloRequest) (model.HelloResponse, error) {
	return model.HelloResponse{Message: "Hello " + req.Name}, nil
}
