package main

import (
	"context"
	"os"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/azzidev/zensegur-provider-tenant/pkg/tenant"
	"github.com/azzidev/zensegur-provider-tenant/pkg/firestore"
)

func main() {
	// Initialize Firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Initialize repository
	repo := firestore.NewRepository(client)

	// Initialize config
	config := &tenant.Config{
		FirestoreProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		DefaultTenant:     "system",
	}

	// Initialize middleware
	middleware := tenant.NewMiddleware(config, repo)

	// Setup routes
	r := gin.Default()
	
	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://portal.zensegur.com.br")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/tenant", func(c *gin.Context) {
			ctx := tenant.FromContext(c)
			c.JSON(200, ctx)
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}