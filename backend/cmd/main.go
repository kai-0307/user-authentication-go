package main

import (
	"log"
	"time"

	"github.com/cos-plat/backend/config"
	"github.com/cos-plat/backend/internal/controller"
	"github.com/cos-plat/backend/internal/model"
	"github.com/cos-plat/backend/internal/pkg/auth"
	"github.com/cos-plat/backend/internal/pkg/database"
	"github.com/cos-plat/backend/internal/pkg/middleware"
	"github.com/cos-plat/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// データベース接続
	db, err := database.NewPostgresDB(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(db)

	// JWTサービスの初期化
	expiresIn, err := time.ParseDuration(cfg.JWT.ExpiresIn)
	if err != nil {
		log.Fatalf("Invalid JWT expiration duration: %s", err)
	}
	jwtService := auth.NewJWTService(cfg.JWT.Secret, expiresIn)

	// サービスの初期化
	// authService := service.NewAuthService(userRepo, jwtService)

	// コントローラーの初期化
	authController := controller.NewAuthService(userRepo, jwtService)
	userController := controller.NewUserController(userRepo)

	// ルーターの設定
	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(middleware.CORS())

	// ルーティング
	api := r.Group("/api")
	{
		// ヘルスチェック
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// 認証不要のエンドポイント
		auth := api.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				var req model.RegisterRequest
				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
				user, err := authController.Register(req)
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, user)
			})
			auth.POST("/login", func(c *gin.Context) {
				var req model.LoginRequest
				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
				token, err := authController.Login(req)
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}
				c.JSON(200, gin.H{"token": token})
			})
		}

		// 認証が必要なエンドポイント
		authorized := api.Group("/")
		authorized.Use(middleware.AuthMiddleware(jwtService))
		{
			user := authorized.Group("/user")
			{
				user.GET("/profile", userController.GetProfile)
				user.PUT("/profile", userController.UpdateProfile)
			}
		}
	}

	// サーバー起動
	log.Printf("Server starting on %s", cfg.Server.Address)
	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
