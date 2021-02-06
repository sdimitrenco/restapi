package handler

import "github.com/gin-gonic/gin"

//Handler struct
type Handler struct {

}

//InitRouts - initializing routs
func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

}