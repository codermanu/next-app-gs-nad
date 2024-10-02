package main

import (
	"github.com/CloudImpl-Inc/next-coder-sdk/polycode"
	_ "portal/register/.polycode"
)

func main() {
	polycode.Start(8080)
}
