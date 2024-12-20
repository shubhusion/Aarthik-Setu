package itr_forms

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	utils "server/config/firebase"
	"strings"
    "os"
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"
)
var client *firestore.Client
func init() {
	client = utils.InitFirestore() // Initialize Firestore client
}

func UploadITRDocuments(c *gin.Context) {
	
    
    userId := c.Param("userId")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data : " + err.Error()})
		return
	}

	files := form.File["itr_documents"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one file is required"})
		return
	}

	if len(files) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 3 files can be uploaded"})
		return
	}

	storageClient := utils.InitStorage()
	if storageClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize storage client"})
		return
	}

	bucketName := os.Getenv("BUCKET_NAME")
	var uploadedFiles []string

	for _, file := range files {
        file.Filename = userId + "_" + file.Filename

		// Check if the file is a PDF
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".pdf") {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File %s is not a PDF", file.Filename)})
			return
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to open file %s", file.Filename)})
			return
		}
		defer src.Close()

		objectName := fmt.Sprintf("business_itr_documents/%s", file.Filename)
		if err := uploadPDF(storageClient, bucketName, objectName, src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload file %s: %v", file.Filename, err)})
			return
		}

		// Construct the file URL
		fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
		uploadedFiles = append(uploadedFiles, fileURL)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d file(s) uploaded successfully", len(uploadedFiles)),
		"files":   uploadedFiles,
	})
	doc := client.Collection("itr_documents").Doc(userId)
	_, err = doc.Set(c, map[string]interface{}{
		"userId":   userId,
		"files": uploadedFiles,
		"userType": "business",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file info to Firestore"})
		return
	}
}

func uploadPDF(client *storage.Client, bucketName, objectName string, src io.Reader) error {
	ctx := context.Background()
	bucket := client.Bucket(bucketName)
	object := bucket.Object(objectName)

	writer := object.NewWriter(ctx)
	defer writer.Close()


	if _, err := io.Copy(writer, src); err != nil {
		log.Printf("Error copying data to Cloud Storage for object %s: %v\n", objectName, err)
		return err
	}
	return nil
}

