package campaign

import (
	"kitabisa/auth"
	"kitabisa/helper"
	"kitabisa/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService CampaignService
	authService     auth.Service
}

func NewCampaignHandler(campaignService CampaignService, authService auth.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService, authService: authService}
}

func (c *campaignHandler) FindAllCampaigns(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.ID != 0 {
		campaigns, err := c.campaignService.FindAllCampaign()
		if err != nil {
			response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error get campaigns")
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseAPI(campaigns, "error", http.StatusOK, "Success get all campaign")
		ctx.JSON(http.StatusOK, response)
		return
	} else {
		response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "You dont have permissions")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
}

func (c *campaignHandler) FindCampaignById(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.ID != 0 {
		campaignIDString := ctx.Param("campaignid")
		campaignId, _ := strconv.Atoi(campaignIDString)
		campaign, err := c.campaignService.FindByIdCampaign(campaignId)
		if err != nil {
			response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error get detail campaign")
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := helper.ResponseAPI(campaign, "error", http.StatusOK, "Success get campaign")
			ctx.JSON(http.StatusOK, response)
			return
		}
	} else {
		response := helper.ResponseAPI(nil, "error", http.StatusUnauthorized, "You dont have permissions")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
}

func (c *campaignHandler) CreateCampaign(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.ID != 0 {
		var input CreateCampaignInput
		err := ctx.ShouldBindJSON(&input)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.ResponseAPI(errorMessage, "error", http.StatusUnprocessableEntity, "Error create campaign1")
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		newCampaign, err := c.campaignService.CreateCampaign(input)
		if err != nil {
			response := helper.ResponseAPI(nil, "error", http.StatusBadRequest, "Error create campaign2")
			ctx.JSON(http.StatusBadRequest, response)
			return
		} else {
			response := helper.ResponseAPI(newCampaign, "success", http.StatusCreated, "Success Create Campaign")
			ctx.JSON(http.StatusCreated, response)
			return
		}
	}
}
