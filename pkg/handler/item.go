package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vladmeh/go-todo-app"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItem(ctx *gin.Context) {

}

func (h *Handler) getItemById(ctx *gin.Context) {

}

func (h *Handler) updateItem(ctx *gin.Context) {

}

func (h *Handler) deleteItem(ctx *gin.Context) {

}
