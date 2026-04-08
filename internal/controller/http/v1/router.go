package v1
import (
	"practice7/internal/usecase"
	"github.com/gin-gonic/gin")
func NewRouter(handler *gin.Engine, u usecase.UserInterface) {
	api := handler.Group("/users")
	newUserRoutes(api, u)}