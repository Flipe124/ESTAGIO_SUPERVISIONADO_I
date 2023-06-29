package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User - Defina a estrutura do modelo de usuário
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  uint   `gorm:"not null"`
}

// UserRepository - Defina uma interface para a camada de dados
type UserRepository interface {
	GetAll() ([]*User, error)
	GetByID(id uint) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}

// GormUserRepository - Implemente a interface UserRepository usando GORM
type GormUserRepository struct {
	db *gorm.DB
}

// GetAll -
func (repo *GormUserRepository) GetAll() ([]*User, error) {
	var users []*User
	err := repo.db.Find(&users).Error
	return users, err
}

// Create -
func (repo *GormUserRepository) Create(user *User) error {
	return repo.db.Create(user).Error
}

// GetByID -
func (repo *GormUserRepository) GetByID(id uint) (*User, error) {
	var user *User
	err := repo.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update -
func (repo *GormUserRepository) Update(user *User) error {
	return repo.db.Save(user).Error
}

// Delete -
func (repo *GormUserRepository) Delete(user *User) error {
	return repo.db.Delete(user).Error
}

// UserController - Defina o controlador para manipular as solicitações HTTP
type UserController struct {
	repo UserRepository
}

// GetAllUsers -
func (ctrl *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := ctrl.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// CreateUser -
func (ctrl *UserController) CreateUser(ctx *gin.Context) {
	var user *User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.repo.Create(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// GetUserByID -
func (ctrl *UserController) GetUserByID(ctx *gin.Context) {
	userID := convertToUint(ctx.Param("id"))
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := ctrl.repo.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateUser -
func (ctrl *UserController) UpdateUser(ctx *gin.Context) {
	userID := convertToUint(ctx.Param("id"))
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := ctrl.repo.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.repo.Update(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser -
func (ctrl *UserController) DeleteUser(ctx *gin.Context) {
	userID := convertToUint(ctx.Param("id"))
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := ctrl.repo.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err := ctrl.repo.Delete(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func convertToUint(str string) uint {
	parsed, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(parsed)
}

func main() {
	// Inicialize o roteador Gin.
	router := gin.Default()

	// Inicialize a conexão com o banco de dados SQLite usando GORM.
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Execute as migrações do banco de dados.
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Crie uma instância do repositório GormUserRepository.
	userRepo := &GormUserRepository{db: db}

	// Crie uma instância do controlador UserController.
	userController := &UserController{repo: userRepo}

	// Defina as rotas da API.
	router.GET("/user", userController.GetAllUsers)
	router.POST("/user", userController.CreateUser)
	router.GET("/user/:id", userController.GetUserByID)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	// Inicie o servidor HTTP.
	if err := router.Run(":5729"); err != nil {
		log.Fatal(err)
	}
}
