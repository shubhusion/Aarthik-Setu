package profile_applications

import (
	"context"
	"net/http"
	utils "server/config/firebase"
	helper "server/helper"
	model "server/models/application"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func init() {
	client = utils.InitFirestore()
}

func CreatePersonalApplication(c *gin.Context) {
	profileId := c.Param("profileId")
	loanType := c.PostForm("loanType")
	id, _ := utils.GenerateRandomString(16)
	ctx := context.Background()

	businessApplication := model.BusinessLoanApplication{
		ProfileId: profileId,
		Id:        id,
		LoanType:  loanType,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	query := client.Collection("personal_applications").
		Where("id", "==", id).
		Where("profileId", "==", profileId).
		Documents(ctx)

	existingDocs, err := query.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing application", "details": err.Error()})
		return
	}
	if len(existingDocs) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Application already exists"})
		return
	}

	dataMap, err := helper.StructToMap(businessApplication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application", "details": err.Error()})
		return
	}

	_, err = client.Collection("personal_applications").Doc(id).Set(ctx, dataMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Application created successfully", "application": dataMap})
}

func GetPersonalApplications(c *gin.Context) {
	profileId := c.Param("profileId")
	ctx := context.Background()

	query := client.Collection("personal_applications").Where("profileId", "==", profileId).Documents(ctx)
	applications, err := query.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve applications", "details": err.Error()})
		return
	}

	if len(applications) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No applications found"})
		return
	}

	var applicationData []map[string]interface{}
	for _, application := range applications {
		applicationData = append(applicationData, application.Data())
	}

	c.JSON(http.StatusOK, gin.H{"applications": applicationData})
}

func UpdatePersonalApplication(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	var updateData map[string]interface{}
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	doc := client.Collection("personal_applications").Doc(id)
	docData, err := doc.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check application", "details": err.Error()})
		return
	}

	if !docData.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	_, err = doc.Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application", "details": err.Error()})
		return
	}

	// Fetch the updated application data
	updatedDocData, err := doc.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated application", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application updated successfully", "application": updatedDocData.Data()})
}

func DeletePersonalApplication(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	doc := client.Collection("personal_applications").Doc(id)
	docData, err := doc.Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check application", "details": err.Error()})
		return
	}

	if !docData.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	_, err = doc.Delete(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
