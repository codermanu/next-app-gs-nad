package controllers

import (
	"github.com/cloudimpl/next-coder-sdk/apicontext"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-gonic/gin"
	"portal/register/model"
)

func Greeting(c *gin.Context) {
	apiCtx, err := apicontext.FromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	greetingService := apiCtx.Service("greeting-service").Get()

	var input model.HelloRequest
	if err = c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var output model.HelloResponse
	res := greetingService.RequestReply(polycode.TaskOptions{}, "greeting", input)
	if err = res.Get(&output); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, output)
}
