package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ta_microservice_auth/app/models"
	"ta_microservice_auth/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthRepo struct {
	Db *gorm.DB
}

func NewAuth() *AuthRepo {
	db := db.InitDb()
	db.AutoMigrate(&models.Anggota{}, &models.AnggotaDetail{})
	return &AuthRepo{Db: db}
}

var SECRET = []byte(os.Getenv("SECRET.KEY"))

func GenerateToken(username string, tokenType string, exp time.Time) string {
	// user := models.Auth{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"token_type": tokenType,
		"exp":        exp.Unix(),
	})

	tokenStr, _ := token.SignedString(SECRET)
	return tokenStr
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return SECRET, nil
	})

	claims, oke := token.Claims.(jwt.MapClaims)
	if oke && token.Valid {
		fmt.Println("token_type and exp :", claims["token_type"], claims["exp"])
	} else {
		log.Println("this not valid")
		log.Println(err)
		return nil, err
	}
	return claims, nil
}

func (auth *AuthRepo) Login(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	req := models.Anggota{}
	if err := c.BindJSON(&req); err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}

	user, err := models.FindUserByNRA(auth.Db, req.Username)
	if err != nil {
		errorMsg := "Invalid credentials."
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}

	if req.Password != user.Password {
		errorMsg := "Invalid credentials."
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		c.Abort()
		return
	}

	accessToken := GenerateToken(user.Username, "access_token", time.Now().Add(15*time.Minute))
	refreshToken := GenerateToken(user.Username, "refresh_token", time.Now().AddDate(0, 0, 5))

	resObj := map[string]interface{}{
		"user": map[string]interface{}{
			"Username": user.Username,
			"Role_id":  user.Role_id,
		},
		"token": map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	}

	res.Data = resObj
	c.JSON(200, res)
}

func (auth *AuthRepo) RefreshToken(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	req := models.ReqAuth{}
	_ = c.BindJSON(&req)
	if req.RefreshToken == "" {
		errMsg := "refresh token is required"
		res.Success = false
		res.Error = &errMsg
		c.JSON(400, res)
		c.Abort()
	}

	claims, err := DecodeToken(req.RefreshToken)
	if err != nil {
		errMsg := err.Error()
		res.Success = false
		res.Error = &errMsg
		c.JSON(401, res)
		c.Abort()
	}

	username, found := claims["username"]
	if !found {
		res.Success = false
		msgErr := "Refresh Token is invalid"
		res.Error = &msgErr
		c.JSON(401, res)
		c.Abort()
		return
	}

	if claims["token_type"] != "refresh_token" {
		res.Success = false
		msgErr := "Refresh Token is invalid"
		res.Error = &msgErr
		c.JSON(401, res)
		c.Abort()
		return
	}

	res.Data = map[string]string{
		"access_token": GenerateToken(username.(string), "access_token", time.Now().Add(15*time.Minute)),
	}

	c.JSON(200, res)

}

func (auth *AuthRepo) CreateAnggota(c *gin.Context) {
	res := models.JsonResponse{Success: true}
	req := models.Anggota{}

	if err := c.BindJSON(&req); err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		return
	}

	_, err := models.FindUserByNRA(auth.Db, req.Username)
	if err == nil {
		errorMsg := "User already exists"
		res.Success = false
		res.Error = &errorMsg
		c.JSON(400, res)
		return
	}

	user := &models.Anggota{
		Username: req.Username,
		Password: req.Password,
		Role_id:  req.Role_id,
		Detail: &models.AnggotaDetail{
			NRA:          req.Detail.NRA,
			Name:         req.Detail.Name,
			Email:        req.Detail.Email,
			Phone_Number: req.Detail.Phone_Number,
			Address:      req.Detail.Address,
			Divisi:       req.Detail.Divisi,
		},
	}

	err = auth.Db.Create(&user).Error
	if err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(500, res)
		return
	}

	// Fetch user's details and associated role
	err = auth.Db.Model(&user).Preload("Role").Preload("Detail").First(&user).Error
	if err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(500, res)
		return
	}

	res.Data = map[string]interface{}{
		"id":       user.Id,
		"role":     user.Role.Name,
		"nra":      user.Detail.NRA,
		"name":     user.Detail.Name,
		"email":    user.Detail.Email,
		"phone":    user.Detail.Phone_Number,
		"address":  user.Detail.Address,
		"division": user.Detail.Divisi,
	}

	c.JSON(200, res)
}

func (repo *AuthRepo) GetAllAnggota(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	var anggota []models.AnggotaDetail

	if err := repo.Db.Find(&anggota).Error; err != nil {
		errorMsg := err.Error()
		res.Success = false
		res.Error = &errorMsg
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	// if err := repo.Db.Preload("Detail").Find(&anggota).Error; err != nil {
	// 	errorMsg := err.Error()
	// 	res.Success = false
	// 	res.Error = &errorMsg
	// 	c.JSON(http.StatusInternalServerError, res)
	// 	return
	// }

	res.Data = anggota
	c.JSON(http.StatusOK, res)
}

func (repo *AuthRepo) TotalAnggota(c *gin.Context) {
	res := models.JsonResponse{Success: true}

	total, err := models.TotalAnggota()
	if err != nil {
		res.Success = false
		MsgErr := err.Error()
		res.Error = &MsgErr
		c.JSON(400, res)
		return
	}

	res.Data = total

	c.JSON(200, res)
}
