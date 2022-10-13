package web

import (
	"github.com/gin-gonic/gin"
	"simple_bank/web/controllers"
)

type ControllerStore struct {
	AccountController controllers.AccountController
}

// declares all the endpoints

func NewServer(store ControllerStore) *gin.Engine {
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"0.0.0.127"}); err != nil {
		return nil
	}
	router.POST("/account", store.AccountController.CreateAccount)
	router.GET("/account/:id", store.AccountController.GetAccount)
	router.GET("/account", store.AccountController.ListAccounts)
	return router
}
