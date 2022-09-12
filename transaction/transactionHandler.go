package transaction

import (
	"kitabisa/helper"
	"kitabisa/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService TransactionService
}

func NewTransactionHandler(service TransactionService) *transactionHandler {
	return &transactionHandler{service}
}

func (t *transactionHandler) GetUserTransactions(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	transactions, err := t.transactionService.GetByUserTransactions(currentUser.ID)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error "+err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.ResponseAPI(transactions, "success", http.StatusOK, "Success get All transactions "+currentUser.Name)
		ctx.JSON(http.StatusOK, response)
	}
}

func (t *transactionHandler) GetCampaignTransactions(ctx *gin.Context) {
	var input GetCampaignTransactionsInput
	idCampaign := ctx.Param("campaignid")
	campaignId, _ := strconv.Atoi(idCampaign)
	input.ID = campaignId
	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	transactions, err := t.transactionService.GetByCampaignID(input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseAPI(transactions, "success", http.StatusOK, "success get transactions")
	ctx.JSON(http.StatusOK, response)
}

func (t *transactionHandler) CreateTransaction(ctx *gin.Context) {
	var input CreateTransactionInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error create transactions"+err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := ctx.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := t.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error create transactions"+err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.ResponseAPI(newTransaction, "error", http.StatusCreated, "Success Create Transactions")
		ctx.JSON(http.StatusCreated, response)
	}
}
