package routes

import (
	"context"
	"fmt"
	"graph-gateway/graph/helpers"
	authhelpers "graph-gateway/helpers"
	"graph-gateway/protos/comms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	commsRoutes := router.Group("/comms")
	{
		// Add more comms-related routes here
		commsRoutes.POST("/MultiChannelCallbackRequest", authhelpers.ExternalCallbackMiddleWare(), multiChannelCallbackRequest)
		commsRoutes.POST("/ProcessOutgoingCallback", authhelpers.ExternalCallbackMiddleWare(), processOutgoingCallbackRequest)
		commsRoutes.POST("/ProcessIncomingCallback", authhelpers.ExternalCallbackMiddleWare(), processIncomingCallbackRequest)
	}
}

func multiChannelCallbackRequest(c *gin.Context) {
	var incomingSMSbody *comms.MultiChannelCallbackRequest
	err := c.ShouldBindJSON(&incomingSMSbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		conn, err := helpers.CommsConnection()
		if err != nil {
			fmt.Println("Error connecting to the client")
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer conn.Close()
		f := comms.NewCommunicatorClient(conn)
		_, err = f.ProcessMultiChannelCallback(context.Background(), incomingSMSbody)
		if err != nil {
			fmt.Println("Error getting response and calling function")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		//returning an empty object
		c.JSON(http.StatusCreated, gin.H{})
	}

}

func processOutgoingCallbackRequest(c *gin.Context) {
	var incomingSMSbody *comms.OutgoingCallbackRequest
	err := c.ShouldBindJSON(&incomingSMSbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		conn, err := helpers.CommsConnection()
		if err != nil {
			fmt.Println("Error connecting to the client")
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer conn.Close()
		f := comms.NewCommunicatorClient(conn)
		_, err = f.ProcessOutgoingCallback(context.Background(), incomingSMSbody)
		if err != nil {
			fmt.Println("Error getting response and calling function")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		//returning an empty object
		c.JSON(http.StatusCreated, gin.H{})
	}

}

func processIncomingCallbackRequest(c *gin.Context) {
	var incomingSMSbody *comms.IncomingSMSCallbackRequest
	err := c.ShouldBindJSON(&incomingSMSbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		conn, err := helpers.CommsConnection()
		if err != nil {
			fmt.Println("Error connecting to the client")
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer conn.Close()
		f := comms.NewCommunicatorClient(conn)
		_, err = f.ProcessIncomingSMSCallback(context.Background(), incomingSMSbody)
		if err != nil {
			fmt.Println("Error getting response and calling function")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		//returning an empty object
		c.JSON(http.StatusCreated, gin.H{})
	}

}
