package repository

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/Golovachev1/travel-agency/internal/app/model"
)

type ReviewRepository struct {
    db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
    return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetReviews() ([]model.Review, error) {
	reviews := []model.Review{}
	err := r.db.Select(&reviews, "SELECT id, score, review_text, publish_date FROM \"review\"")
	if err != nil {
	 fmt.Println("get reviews err", err)
	 return nil, err
	}
	return reviews, nil
   }


   func (r *ReviewRepository) GetReview(id int) (*model.Review, error) {
	review := model.Review{}
	err := r.db.Get(&review, "SELECT id, score, review_text, publish_date FROM \"review\" WHERE id=$1", id)
	if err != nil {
	 fmt.Println("get review err", err)
	 return nil, err
	}
	return &review, nil
   }

   func (r *ReviewRepository) CreateReview(Score int, Review_Text string, Publish_date string, userID int, tourID int) error {
    _, err := r.db.Exec("INSERT INTO Review (score, review_text, publish_date, user_id, tour_id) VALUES ($1, $2, $3, $4, $5)", Score, Review_Text, Publish_date, userID, tourID)
    return err
}


func (r *ReviewRepository) DeleteReview(id int) error {
    _, err := r.db.Exec("DELETE FROM \"review\" WHERE id=$1", id)
    
    if err != nil {
        fmt.Println("delete review err", err)
        return err
    }

    return nil
}

func (r *ReviewRepository) UpdateReview(id int, score int, reviewText string, publishDate string, userID int, tourID int) error {
    _, err := r.db.Exec(
        "UPDATE Review SET score = $1, review_text = $2, publish_date = $3, user_id = $4, tour_id = $5 WHERE id = $6",
        score, reviewText, publishDate, userID, tourID, id,
    )
    return err
}
