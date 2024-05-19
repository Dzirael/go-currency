package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Dzirael/go-curenncy/internal/email"
	"github.com/Dzirael/go-curenncy/internal/pkg/models/users"
	"github.com/Dzirael/go-curenncy/internal/pkg/persistence"
	http_err "github.com/Dzirael/go-curenncy/pkg/http-err"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Email string `json:"email" binding:"required,email"`
}

// @Summary Create a new subscription
// @Description Create a new subscription with the provided email address
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param input body UserInput true "User input data"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /subscribe [post]
func CreateSubscribe(c *gin.Context) {
	s := persistence.GetUserRepository()
	var userInput UserInput
	if err := c.BindJSON(&userInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}
	user := users.User{
		Email:        userInput.Email,
		IsSubscribed: true,
	}
	if err := s.Add(&user); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			http_err.NewError(c, http.StatusConflict, err)
		} else {
			http_err.NewError(c, http.StatusInternalServerError, err)
		}
		log.Println(err)
	} else {
		c.Status(http.StatusOK)
	}
}

// @Summary Відправити e-mail з поточним курсом на всі підписані електронні пошти
// @Description Запит має отримувати актуальний курс USD до UAH за допомогою third-party сервісу та відправляти його на всі електронні адреси, які були підписані раніше.
// @ID sendEmails
// @Tags subscription
// @Produce json
// @Success 200 {string} string "E-mailʼи відправлено"
// @Router /sendEmails [post]
func SendEmails(c *gin.Context) {
	email.SendToAll()
	c.String(http.StatusOK, "E-mailʼи відправлено")
}
