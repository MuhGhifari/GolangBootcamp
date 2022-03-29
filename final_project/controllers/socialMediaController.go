package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MuhGhifari/GolangBootcamp/final-project/config"
	"github.com/MuhGhifari/GolangBootcamp/final-project/helpers"
	"github.com/MuhGhifari/GolangBootcamp/final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SocialMediaShow struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	UpdatedAt *time.Time
	CreatedAt *time.Time
	User      *SocialMediaUser
}

func GetSocialMedia(c *gin.Context) {
	var socialMedias []models.Comment
	var socialMediaShows []SocialMediaShow
	var socialMediaUser SocialMediaUser
	db := config.GetDB()
	db.Find(&socialMedias)

	for _, socialMedia := range socialMedias {
		var socialMediaShow SocialMediaShow

		db.Model(&models.User{}).Where("id = ?", socialMedia.UserId).Find(&socialMediaUser)

		socialMediaShow.Id = int(socialMedia.ID)
		socialMediaShow.Message = socialMedia.Message
		socialMediaShow.UserId = socialMedia.UserId
		socialMediaShow.PhotoId = socialMedia.PhotoId
		socialMediaShow.CreatedAt = socialMedia.CreatedAt
		socialMediaShow.UpdatedAt = socialMedia.UpdatedAt
		socialMediaShow.User = &socialMediaUser

		socialMediaShows = append(socialMediaShows, socialMediaShow)
	}

	c.JSON(http.StatusOK, socialMediaShows)
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

	fmt.Println(SocialMedia)

	err := db.Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	SocialMediaId, _ := strconv.Atoi(c.Param("SocialMediaId"))
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userID
	SocialMedia.ID = uint(SocialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", SocialMediaId).Updates(models.SocialMedia{
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

	c.JSON(http.StatusOK, SocialMedia)
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
