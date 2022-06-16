package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vladmeh/go-todo-app"
	"net/http"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) singIn(c *gin.Context) {

}
