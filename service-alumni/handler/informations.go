package handler

import (
	"net/http"
	"service-alumni/helper"
	"service-alumni/informations"
	"strconv"

	"github.com/gin-gonic/gin"
)

type informationsHandler struct {
	informationsService informations.Service
}

func NewInformationsHandler(informationsService informations.Service) *informationsHandler {
	return &informationsHandler{informationsService}
}

func (h *informationsHandler) CreateInformations(c *gin.Context) {
	var input informations.CreateInformationsInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors :": errors}
		response := helper.APIResponse("Information not added", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	information, err := h.informationsService.CreateInformations(input)

	if err != nil {
		response := helper.APIResponse("Information not added", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Information has been added", http.StatusOK, "success", information)
	c.JSON(http.StatusOK, response)
}

func (h *informationsHandler) GetInformations(c *gin.Context) {
	IDUser, _ := strconv.Atoi(c.Query("id_user"))

	information, err := h.informationsService.GetInformations(IDUser)

	if err != nil {
		response := helper.APIResponse("Error get Informations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Render Information Success", http.StatusOK, "success", information)
	c.JSON(http.StatusOK, response)

}

func (h *informationsHandler) GetInformationsDetail(c *gin.Context) {
	var input informations.GetInformationsDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Error get Informations Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	informationDetail, err := h.informationsService.GetInformationsByID(input)

	if err != nil {
		response := helper.APIResponse("Error get Informations Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Render Information Detail Success", http.StatusOK, "success", informationDetail)
	c.JSON(http.StatusOK, response)

}

func (h *informationsHandler) UpdateInformations(c *gin.Context) {
	var ID informations.GetInformationsDetailInput

	err := c.ShouldBindUri(&ID)

	if err != nil {
		response := helper.APIResponse("Error update Informations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input informations.CreateInformationsInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors :": errors}
		response := helper.APIResponse("Error Update Information", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedInformation, err := h.informationsService.UpdateInformations(ID, input)
	if err != nil {
		response := helper.APIResponse("Error update Informations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Update Information Success", http.StatusOK, "success", updatedInformation)
	c.JSON(http.StatusOK, response)
}

func (h *informationsHandler) DeleteInformation(c *gin.Context) {
	var ID informations.GetInformationsDetailInput

	err := c.ShouldBindUri(&ID)

	if err != nil {
		response := helper.APIResponse("Error Delete Informations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.informationsService.DeleteInformation(ID)

	if err != nil {
		response := helper.APIResponse("Error Delete Information", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Information Deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
