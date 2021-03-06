package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MuhGhifari/GolangBootcamp/final-project/config"
	"github.com/MuhGhifari/GolangBootcamp/final-project/helpers"
	"github.com/MuhGhifari/GolangBootcamp/final-project/models"
	"github.com/gin-gonic/gin"
)

type userUpdateResult struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       int        `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"username": User.Username,
		"id":       User.ID,
	})
}

func UserLogin(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Model(&User).Where("id = ?", userId).Updates(models.User{
		Email:    User.Email,
		Username: User.Username,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		db.Where("id = ?", userId).First(&User)
		result := userUpdateResult{
			ID:        int(User.ID),
			Email:     User.Email,
			Username:  User.Username,
			Age:       User.Age,
			UpdatedAt: User.UpdatedAt,
		}
		c.JSON(http.StatusOK, result)
	}
}

func DeleteUser(c *gin.Context) {
	var User models.User
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("userId")).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db.Delete(&User)
	c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}
