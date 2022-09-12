package transaction

import (
	"kitabisa/helper"
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
