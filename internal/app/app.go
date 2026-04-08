package app
import 
(   "log"
	"practice7/internal/controller/http/v1"
	"practice7/internal/entity"
	"practice7/internal/usecase"
	"practice7/internal/usecase/repo"
	"practice7/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv")

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")}
	pg, err := postgres.New()
	if err != nil {
		log.Fatal("DB connection error:", err)}
	pg.Conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	pg.Conn.AutoMigrate(&entity.User{})
	userRepo := repo.NewUserRepo(pg)
	userUseCase := usecase.NewUserUseCase(userRepo)
	router := gin.Default()
	v1.NewRouter(router, userUseCase)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server error:", err)
	}
}