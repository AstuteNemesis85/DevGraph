package graph

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Trigger graph building manually
func BuildGraph(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := RebuildSimilarityGraph(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to build graph"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "graph built successfully"})
	}
}

// Background function to rebuild the similarity graph
func RebuildSimilarityGraph(db *gorm.DB) error {
	log.Println("Building similarity graph...")

	// Build user profiles
	profiles, err := BuildUserProfiles(db)
	if err != nil {
		log.Println("failed to build profiles:", err)
		return err
	}

	log.Printf("Built %d user profiles\n", len(profiles))

	if len(profiles) < 2 {
		log.Println("Need at least 2 users to build graph")
		return nil
	}

	// Build similarity graph with threshold 0.1 (10% similarity)
	edges := BuildSimilarityGraph(profiles, 0.1)
	log.Printf("Found %d similarity edges\n", len(edges))

	// Persist to database
	err = PersistGraph(db, edges)
	if err != nil {
		log.Println("failed to persist graph:", err)
		return err
	}

	log.Println("Similarity graph built successfully")
	return nil
}
