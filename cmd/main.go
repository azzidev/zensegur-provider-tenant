package main

import (
	"context"
	"os"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"

	tenantFirestore "github.com/azzidev/zensegur-provider-tenant/pkg/firestore"
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
	repo := tenantFirestore.NewRepository(client)





	// Setup routes
	r := gin.Default()
	
	// CORS - Only allow auth service
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://zensegur.com.br")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public API for auth service only
	api := r.Group("/api")
	{
		// Get tenant by ID (for auth service)
		api.GET("/tenant/:id", func(c *gin.Context) {
			tenantID := c.Param("id")
			tenantInfo, err := repo.GetByID(c.Request.Context(), tenantID)
			if err != nil {
				c.JSON(404, gin.H{"error": "Tenant not found"})
				return
			}
			c.JSON(200, tenantInfo)
		})
		
		// Get tenant by alias (for auth service)
		api.GET("/tenant/alias/:alias", func(c *gin.Context) {
			alias := c.Param("alias")
			tenantInfo, err := repo.GetByAlias(c.Request.Context(), alias)
			if err != nil {
				c.JSON(404, gin.H{"error": "Tenant not found"})
				return
			}
			c.JSON(200, tenantInfo)
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}