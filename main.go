package main

import (
	"github.com/cloudimpl/next-coder-sdk/polycode"
	_ "portal/register/.polycode"
	"portal/register/controllers"
	"portal/register/lib"
)

func main() {
	v := lib.NewValidator()
	polycode.SetValidator(v)

	s := controllers.GetServer()
	polycode.StartApp(s)
}
