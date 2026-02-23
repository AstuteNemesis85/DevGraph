package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProfileResponse is the shape returned by GET /api/me and GET /api/profile.
type ProfileResponse struct {
	ID        uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
	Bio       string    `json:"bio"`
	CreatedAt string    `json:"created_at"`
}

// UpdateProfileRequest holds the fields a user is allowed to change.
type UpdateProfileRequest struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

// GetProfile handles GET /api/me and GET /api/profile.
// Returns the authenticated user's full profile from the database.
func GetProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uuid.UUID)

		var u User
		if err := db.Where("id = ?", userID).First(&u).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, ProfileResponse{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			AvatarURL: u.AvatarURL,
			Bio:       u.Bio,
			CreatedAt: u.CreatedAt.Format("2006-01-02"),
		})
	}
}

// UpdateProfile handles PUT /api/profile.
// Allows the user to update username, bio, and avatar_url.
// Empty string fields are ignored (not cleared) unless explicitly provided.
func UpdateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uuid.UUID)

		var req UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var u User
		if err := db.Where("id = ?", userID).First(&u).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		// Apply only non-empty fields from the request
		if req.Username != "" {
			u.Username = req.Username
		}
		// Bio and AvatarURL can be updated to empty (to clear them), so always apply
		u.Bio = req.Bio
		u.AvatarURL = req.AvatarURL

		if err := db.Save(&u).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
			return
		}

		c.JSON(http.StatusOK, ProfileResponse{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			AvatarURL: u.AvatarURL,
			Bio:       u.Bio,
			CreatedAt: u.CreatedAt.Format("2006-01-02"),
		})
	}
}
