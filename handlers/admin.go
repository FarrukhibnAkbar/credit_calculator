package handlers

import (
	"delivery/constants"
	"delivery/entities"
	"delivery/pkg/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CalculateCredit(c *gin.Context) {
	var body entities.CalculateModel
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{
			"message": "error in body parse",
			"errors":  handleBodyParseError(err),
		})
		return
	}
	if err := h.validate.Struct(body); err != nil {
		h.handleResponse(c, http.UnprocessableEntity, err.Error())
		return
	}
	if !slices.Contains([]string{"differential", "annuitet"}, body.CreditType) {
		h.handleResponse(c, http.BadRequest, "credit type must be differential or annuitet")
		return
	}
	if body.AnnualRate>30 {
		h.handleResponse(c, http.BadRequest, "yillik foiz 30 foizdan ko'p bo'lishi mumkin emas")
		return
	}
	if body.Principal>100000000 {
		h.handleResponse(c, http.BadRequest, "100 000 000 so'mdan ortiq kredit berilmaydi")
		return
	}
	body.ID = uuid.New()

	res, err := h.adminController.CalculateCredit(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}

	h.handleResponse(c, http.OK, res)
}
