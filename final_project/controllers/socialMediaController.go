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

type socialMediaCreatedResult struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type socialMediaUpdatedResult struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         int        `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type SocialMediaUser struct {
	ID              int `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `default:"avatar.png" json:"profile_image_url"`
}

type SocialMediaShow struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         int        `json:"user_id"`
	UpdatedAt      *time.Time `json:"created_at"`
	CreatedAt      *time.Time `json:"updated_at"`
	User           *SocialMediaUser
}

type SocialMediaArray struct {
	SocialMedia []SocialMediaShow `json:"social_medias"`
}

func GetSocialMedia(c *gin.Context) {
	var socialMedias []models.SocialMedia
	var socialMediaShows []SocialMediaShow
	var socialMediaUser SocialMediaUser
	db := config.GetDB()
	db.Find(&socialMedias)

	for _, socialMedia := range socialMedias {
		var socialMediaShow SocialMediaShow
		var user models.User
		
		db.Model(&models.User{}).Where("id = ?", socialMedia.UserId).Find(&user)

		socialMediaUser.ID = int(user.ID)
		socialMediaUser.Username = user.Username
		socialMediaUser.ProfileImageUrl = "default-avatar.png"


		socialMediaShow.ID = int(socialMedia.ID)
		socialMediaShow.Name = socialMedia.Name
		socialMediaShow.SocialMediaUrl = socialMedia.SocialMediaUrl
		socialMediaShow.UserId = socialMedia.UserId
		socialMediaShow.CreatedAt = socialMedia.CreatedAt
		socialMediaShow.UpdatedAt = socialMedia.UpdatedAt
		socialMediaShow.User = &socialMediaUser

		socialMediaShows = append(socialMediaShows, socialMediaShow)
	}

	result := SocialMediaArray{socialMediaShows}

	c.JSON(http.StatusOK, result)
}

func CreateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	var SocialMedia models.SocialMedia
	userID := int(userData["id"].(float64))

	SocialMedia.UserId = userID

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	result := socialMediaCreatedResult{
		ID:             int(SocialMedia.ID),
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
		UserId:         SocialMedia.UserId,
		CreatedAt:      SocialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Where("id = ?", socialMediaId).First(&SocialMedia)
	result := socialMediaUpdatedResult{
		ID:             int(SocialMedia.ID),
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
		UserId:         SocialMedia.UserId,
		UpdatedAt:      SocialMedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, result)
}

func DeleteSocialMedia(c *gin.Context) {
	var SocialMedia models.SocialMedia
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("socialMediaId")).First(&SocialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db.Delete(&SocialMedia)
	c.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted"})
}
