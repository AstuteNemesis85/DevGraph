package auth

import (
	"net/http"
	"time"

	"devgraph/internal/cache"
	"devgraph/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}

		newUser := user.User{
			ID:           uuid.New(),
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: hashedPassword,
			CreatedAt:    time.Now(),
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user registered successfully",
			"user_id": newUser.ID,
		})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser user.User
		if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if !CheckPasswordHash(req.Password, existingUser.PasswordHash) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Generate access token
		accessToken, err := GenerateAccessToken(existingUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
			return
		}

		// Generate refresh token
		refreshToken, err := GenerateRefreshToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
			return
		}

		hashedRefresh := HashRefreshToken(refreshToken)

		session := Session{
			ID:               uuid.New(),
			UserID:           existingUser.ID,
			RefreshTokenHash: hashedRefresh,
			ExpiresAt:        time.Now().Add(7 * 24 * time.Hour),
			CreatedAt:        time.Now(),
		}

		// Store session in DB
		if err := db.Create(&session).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store session"})
			return
		}

		// Cache session in Redis
		redisClient := cache.NewRedisClient()
		err = redisClient.Set(
			cache.Ctx,
			"session:"+hashedRefresh,
			session.UserID.String(),
			time.Until(session.ExpiresAt),
		).Err()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cache session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func Refresh(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RefreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashed := HashRefreshToken(req.RefreshToken)

		redisClient := cache.NewRedisClient()

		_, err := redisClient.Get(cache.Ctx, "session:"+hashed).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}

		var session Session
		if err := db.Where("refresh_token_hash = ?", hashed).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}

		if time.Now().After(session.ExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
			return
		}

		// Rotate refresh token
		newRefresh, err := GenerateRefreshToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to rotate refresh token"})
			return
		}

		session.RefreshTokenHash = HashRefreshToken(newRefresh)
		session.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)

		if err := db.Save(&session).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update session"})
			return
		}

		// New access token
		accessToken, err := GenerateAccessToken(session.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": newRefresh,
		})
	}
}

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RefreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashed := HashRefreshToken(req.RefreshToken)

		// Delete from Redis
		redisClient := cache.NewRedisClient()
		redisClient.Del(cache.Ctx, "session:"+hashed)

		// Delete from DB
		db.Where("refresh_token_hash = ?", hashed).Delete(&Session{})

		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}
}
