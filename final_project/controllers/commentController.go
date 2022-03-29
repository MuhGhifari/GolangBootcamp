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

type CommentUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CommentPhoto struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type CommentShow struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	UpdatedAt *time.Time
	CreatedAt *time.Time
	User      *CommentUser
	Photo     *CommentPhoto
}

func GetComments(c *gin.Context) {
	var comments []models.Comment
	var commentShows []CommentShow
	var commentUser CommentUser
	var commentPhoto CommentPhoto
	db := config.GetDB()
	db.Find(&comments)

	for _, comment := range comments {
		var commentShow CommentShow

		db.Model(&models.User{}).Where("id = ?", comment.UserId).Find(&commentUser)
		db.Model(&models.Photo{}).Where("id = ?", comment.PhotoId).Find(&commentPhoto)

		commentShow.Id = int(comment.ID)
		commentShow.Message = comment.Message
		commentShow.UserId = comment.UserId
		commentShow.PhotoId = comment.PhotoId
		commentShow.CreatedAt = comment.CreatedAt
		commentShow.UpdatedAt = comment.UpdatedAt
		commentShow.User = &commentUser
		commentShow.Photo = &commentPhoto

		commentShows = append(commentShows, commentShow)
	}

	c.JSON(http.StatusOK, commentShows)
}

func CreateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	var Comment models.Comment
	userID := int(userData["id"].(float64))

	Comment.UserId = userID

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	fmt.Println(Comment)

	err := db.Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userID
	Comment.ID = uint(commentId)

	fmt.Println(Comment.PhotoId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	var Comment models.Comment
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("commentId")).First(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db.Delete(&Comment)
	c.JSON(http.StatusOK, gin.H{"message": "Your comment has been successfully deleted"})
}
