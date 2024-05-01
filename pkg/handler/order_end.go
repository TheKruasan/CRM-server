package handler

import (
	"net/http"

	models "github.com/TheKruasan/CRM-server/internal/repository/model"
	"github.com/gin-gonic/gin"
)

// @Summary GetOrder
// @Tags order
// @Description get order from db
// @ID get-order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order "data"
// @Failure 400,404 {object} HandleError
// @Failure 500 {object} HandleError
// @Failure default {object} HandleError
// @Router /order/get/:id [get]
func (h *Handler) getOrderById(c *gin.Context) {

}

// @Summary GetAllOrders
// @Tags orders
// @Description get all orders from db
// @ID get-orders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Order
// @Failure 400,404 {object} HandleError
// @Failure 500 {object} HandleError
// @Failure default {object} HandleError
// @Router /order/get/all [get]
func (h *Handler) getAllOrderById(c *gin.Context) {

}

// @Summary createOrder
// @Tags create-order
// @Description create a new order
// @ID create-order
// @Accept  json
// @Produce  json
// @Param input body models.Order true "order"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} HandleError
// @Failure 500 {object} HandleError
// @Failure default {object} HandleError
// @Router /order/add [post]
func (h *Handler) postOrder(c *gin.Context) {
	var input models.Order

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.Order.CreateOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
