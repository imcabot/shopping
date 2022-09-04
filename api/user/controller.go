package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shoping/config"
	"shoping/domain/user"
	"shoping/utils/api_helper"
	jwtHelper "shoping/utils/jwt"
	"strconv"
	"time"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.ConfigUration
}

func NewUserController(service *user.Service, appConfig *config.ConfigUration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

func (c *Controller) CreateUser(g *gin.Context) {
	var req CreatUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return

	}
	newUser := user.NewUser(req.Username, req.password, req.password2)
	err := c.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
	}
	g.JSON(
		http.StatusCreated, CreatUserResponse{
			Username: req.Username,
		})
}

func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
	}
	currentUser, err := c.userService.GetUser(req.Username, req.password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	decodedClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(24 * time.Hour).Unix(),
				"isAdmin":  currentUser.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err := c.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}
	g.JSON(http.StatusOK, LoginResponse{
		Username: currentUser.Username, UserID: currentUser.ID, Token: currentUser.Token,
	})
}

func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)
	g.JSON(http.StatusOK, decodedClaims)
}
