package personal_forms

import (
	"server/controller/forms/personal_forms/itr_forms"

	"github.com/gin-gonic/gin"
)

// RegisterITRRoutes registers the routes related to ITR forms
func RegisterPersonalITRRoutes(router *gin.Engine) {
	router.POST("/personal/itr", itr_forms.CreateManualITR)
	router.GET("/personal/itr/:userId", itr_forms.GetManualITR)
	router.PATCH("/personal/itr/:userId", itr_forms.UpdateManualITR)
}
