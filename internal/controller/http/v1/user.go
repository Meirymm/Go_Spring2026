package v1
import (
	"net/http"
	"practice7/internal/entity"
	"practice7/internal/usecase"
	"practice7/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid")
type userRoutes struct {
	t usecase.UserInterface}
func newUserRoutes(handler *gin.RouterGroup, t usecase.UserInterface) {
	r := &userRoutes{t}
	handler.POST("/", r.RegisterUser)
	handler.POST("/login", r.LoginUser)
	protected := handler.Group("/")
	protected.Use(utils.JWTAuthMiddleware(), utils.RateLimiterMiddleware())
	{
		protected.GET("/me", r.GetMe)
		protected.PATCH("/promote/:id", utils.RoleMiddleware("admin"), r.PromoteUser)
	}}
func (r *userRoutes) RegisterUser(c *gin.Context) {
	var dto entity.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return}
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return}
	role := dto.Role
	if role == "" {
		role = "user"	}
	user := &entity.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPassword,
		Role:     role,}
	createdUser, sessionID, err := r.t.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return}
	c.JSON(http.StatusCreated, gin.H{
		"message":    "User registered successfully",
		"session_id": sessionID,
		"user":       createdUser,
	})
}
func (r *userRoutes) LoginUser(c *gin.Context) {
	var input entity.LoginUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := r.t.LoginUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (r *userRoutes) GetMe(c *gin.Context) {
	userIDStr, _ := c.Get("userID")
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return}
	user, err := r.t.GetMe(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return}
	c.JSON(http.StatusOK, user)}

func (r *userRoutes) PromoteUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return}
	user, err := r.t.PromoteUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return}
	c.JSON(http.StatusOK, user)
}