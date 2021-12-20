package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hb.study/mysql/code/gorm_gen_demo/repository"
)

type userController struct {
	repo repository.UserRepo
}

func NewUserController(repo repository.UserRepo) userController {
	return userController{repo}
}

func (uc userController) CreateUser(c *gin.Context) {
	uc.repo.CreateUser(c.Request.Context())
	c.JSON(http.StatusOK, nil)
}

func (uc userController) GetUser(c *gin.Context) {
	user, err := uc.repo.GetUser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc userController) GetUserByCondition(c *gin.Context) {
	user, err := uc.repo.GetUserByCondition(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc userController) GetUserBySelectFiled(c *gin.Context) {
	user, err := uc.repo.GetUserBySelectFiled(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc userController) GetUserByPage(c *gin.Context) {
	user, err := uc.repo.GetUserByPage(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc userController) GetSingleFiledByPluck(c *gin.Context) {
	user, err := uc.repo.GetSingleFiledByPluck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc userController) UpdateSingleFiled(c *gin.Context) {
	err := uc.repo.UpdateSingleFiled(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc userController) UpdateSelectFiled(c *gin.Context) {
	err := uc.repo.UpdateSelectFiled(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc userController) UpdateOmitFiled(c *gin.Context) {
	err := uc.repo.UpdateOmitFiled(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc userController) DeleteAndUpdate(c *gin.Context) {
	err := uc.repo.DeleteAndUpdate(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc userController) DeleteAndCreate(c *gin.Context) {
	err := uc.repo.DeleteAndCreate(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
