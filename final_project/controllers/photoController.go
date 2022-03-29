package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MuhGhifari/GolangBootcamp/final-project/config"
	"github.com/MuhGhifari/GolangBootcamp/final-project/helpers"
	"github.com/MuhGhifari/GolangBootcamp/final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type photoCreatedResult struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type photoUpdatedResult struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    int        `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoShow struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserId    int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *PhotoUser
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	var photoShows []PhotoShow
	var photoUser PhotoUser
	db := config.GetDB()
	db.Find(&photos)

	for _, photo := range photos {
		var photoShow PhotoShow

		db.Model(&models.User{}).Where("id = ?", photo.UserId).Find(&photoUser)

		photoShow.Id = int(photo.ID)
		photoShow.Title = photo.Title
		photoShow.Caption = photo.Caption
		photoShow.PhotoUrl = photo.PhotoUrl
		photoShow.UserId = photo.UserId
		photoShow.CreatedAt = photo.CreatedAt
		photoShow.UpdatedAt = photo.UpdatedAt
		photoShow.User = &photoUser

		photoShows = append(photoShows, photoShow)
	}

	c.JSON(http.StatusOK, photoShows)
}

func CreatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	var Photo models.Photo
	userID := int(userData["id"].(float64))

	Photo.UserId = userID

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result := photoCreatedResult{
		ID:        int(Photo.ID),
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		PhotoUrl:  Photo.PhotoUrl,
		UserId:    Photo.UserId,
		CreatedAt: Photo.CreatedAt,
	}

	c.JSON(http.StatusCreated, result)
}

func UpdatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
		UserId:   Photo.UserId,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	db.Where("id = ?", photoId).First(&Photo)
	result := photoUpdatedResult{
		ID:        int(Photo.ID),
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		PhotoUrl:  Photo.PhotoUrl,
		UserId:    Photo.UserId,
		UpdatedAt: Photo.UpdatedAt,
	}

	c.JSON(http.StatusOK, result)
}

func DeletePhoto(c *gin.Context) {
	var Photo models.Photo
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("photoId")).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db.Delete(&Photo)
	c.JSON(http.StatusOK, gin.H{"message": "Your photo has been successfully deleted"})
}
