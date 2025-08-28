package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/SampleAPI/internal/dto"
	"github.com/raulsilva-tech/SampleAPI/internal/entity"
	"github.com/raulsilva-tech/SampleAPI/internal/repository"
	"github.com/raulsilva-tech/SampleAPI/internal/usecase"
)

type UserHandler struct {
	UserRepository      *repository.UserRepository
	EventRepository     *repository.EventRepository
	EventTypeRepository *repository.EventTypeRepository
	JWTSecret           string
	JWTExpiresIn        int
}

func NewUserHandler(ur *repository.UserRepository, er *repository.EventRepository, etr *repository.EventTypeRepository, jwtSecret string, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserRepository:      ur,
		EventRepository:     er,
		EventTypeRepository: etr,
		JWTSecret:           jwtSecret,
		JWTExpiresIn:        jwtExpiresIn,
	}
}

func (h *UserHandler) Login(c *gin.Context) {

	var input dto.LoginInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	output, err := uc.Login(c.Request.Context(), usecase.LoginUseCaseInput{Email: input.Email, Password: input.Password, JWTSecret: h.JWTSecret, JWTExpiresIn: h.JWTExpiresIn})
	if err != nil {
		if err == entity.ErrLoginFailed {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else if err == sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, output.Token)

}

func (h *UserHandler) Insert(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to insert records."})
		return
	}

	var input dto.UserInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ucInput := usecase.UserUseCaseInput{
		Name:          input.Name,
		Email:         input.Email,
		Password:      input.Password,
		CreatorUserId: userID.(string),
	}

	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	_, err = uc.RegisterUser(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]any{"msg": "created successfully"})
}

func (h *UserHandler) Delete(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to delete records."})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	err := uc.RemoveUser(c.Request.Context(), id, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg": "deleted successfully"})
}

func (h *UserHandler) Update(c *gin.Context) {

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to update records."})
		return
	}

	var input dto.UserInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ucInput := usecase.UserUseCaseInput{
		Id:            id,
		Name:          input.Name,
		Email:         input.Email,
		Password:      input.Password,
		CreatorUserId: userId.(string),
	}
	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	err = uc.UpdateUser(c.Request.Context(), ucInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]any{"msg": "updated successfully"})
}

func (h *UserHandler) GetOne(c *gin.Context) {

	_, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to get a record."})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	output, err := uc.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)

}

func (h *UserHandler) GetAll(c *gin.Context) {

	_, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No authenticated user. An user is required to get records"})
		return
	}

	uc := usecase.NewUserUseCase(h.UserRepository, h.EventRepository, h.EventTypeRepository)
	outputList, err := uc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	c.JSON(http.StatusOK, outputList)
}
