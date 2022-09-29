package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	models "simple_bank/db"
	"simple_bank/internal/repository"
	"simple_bank/internal/service"
)

type AccountController interface {
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
	ListAccounts(ctx *gin.Context)
}

type accountController struct {
	accountService service.AccountStore
}

func NewAccountController() AccountController {
	return &accountController{
		accountService: service.NewAccountService(repository.DB),
	}
}

// =================================================== //

func (c *accountController) CreateAccount(ctx *gin.Context) {
	var createAccountRequest models.Account
	if err := ctx.ShouldBindJSON(&createAccountRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId, err := c.accountService.CreateAccount(createAccountRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, accountId)
}

func (c *accountController) GetAccount(ctx *gin.Context) {
	var req models.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := c.accountService.GetAccount(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (c *accountController) ListAccounts(ctx *gin.Context) {
	accounts, err := c.accountService.ListAccounts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
