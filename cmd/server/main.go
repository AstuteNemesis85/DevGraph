package main

import (
	"log"

	"devgraph/internal/analysis"
	"devgraph/internal/auth"
	"devgraph/internal/code"
	"devgraph/internal/config"
	"devgraph/internal/graph"
	"devgraph/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Auto-migrate schemas
	err = db.AutoMigrate(
		&user.User{},
		&auth.Session{},
		&code.CodeSubmission{},
		&analysis.CodeAnalysis{},
		&analysis.AlgorithmPattern{},
		&analysis.SubmissionPattern{},
		&graph.UserSimilarityEdge{}, // 👈 PHASE 7 TABLE
	)

	if err != nil {
		log.Fatal("Database migration failed:", err)
	}
	analysis.StartWorkerPool(db, 4)



	// Gin router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth routes (public)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", auth.Register(db))
		authGroup.POST("/login", auth.Login(db))
		authGroup.POST("/refresh", auth.Refresh(db))
		authGroup.POST("/logout", auth.Logout(db))


	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"user_id": userID,
				"message": "you are authenticated",
			})
		})
		protected.POST("/submit", code.SubmitCode(db))
		protected.GET("/submissions", analysis.GetUserSubmissions(db))
		protected.GET("/analysis/:id", analysis.GetAnalysis(db))
		protected.GET("/recommendations", graph.GetRecommendations(db))
		protected.POST("/build-graph", graph.BuildGraph(db))


	}

	log.Println("Server running on :8080")
	r.Run(":8080") // 🚀 ALWAYS LAST
}
