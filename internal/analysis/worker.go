package analysis

import (
	"log"
	//"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func StartWorkerPool(db *gorm.DB, workers int) {
	for i := 0; i < workers; i++ {
		go worker(db, i)
	}
}

func worker(db *gorm.DB, id int) {
	log.Printf("analysis worker %d started\n", id)

	for submissionID := range JobQueue {
		log.Printf("worker %d processing submission %s\n", id, submissionID)
		analyzeSubmission(db, submissionID)
	}
}

func analyzeSubmission(db *gorm.DB, submissionID interface{}) {
	var submission struct {
		ID         uuid.UUID
		SourceCode string
	}

	if err := db.Table("code_submissions").
		Select("id, source_code").
		Where("id = ?", submissionID).
		Scan(&submission).Error; err != nil {
		log.Println("failed to load submission:", err)
		return
	}

	code := submission.SourceCode

	patterns := detectPatterns(code)
	timeC, spaceC := inferComplexity(code)

	analysis := CodeAnalysis{
		ID:              uuid.New(),
		SubmissionID:    submission.ID,
		TimeComplexity:  timeC,
		SpaceComplexity: spaceC,
		CreatedAt:       time.Now(),
	}

	if err := db.Create(&analysis).Error; err != nil {
		log.Println("failed to store analysis:", err)
		return
	}

	for _, p := range patterns {
		var pattern AlgorithmPattern
	err := db.Where("name = ?", p).First(&pattern).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			pattern = AlgorithmPattern{
				ID:   uuid.New(),
				Name: p,
			}
			db.Create(&pattern)
		} else {
			log.Println("failed to load pattern:", err)
			return
		}
	}


			db.FirstOrCreate(
		&SubmissionPattern{},
		SubmissionPattern{
			SubmissionID: submission.ID,
			PatternID:    pattern.ID,
		},
	)

	}

	log.Printf("analysis stored for submission %s\n", submission.ID)
}
