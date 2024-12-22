package repository

import (
    "fmt"
    "time"
    "github.com/jmoiron/sqlx"
    "github.com/Golovachev1/travel-agency/internal/app/model"
)

type TourRepository struct {
    db *sqlx.DB
}

func NewTourRepository(db *sqlx.DB) *TourRepository {
    return &TourRepository{db: db}
}

func (r *TourRepository) GetTours() ([]model.Tour, error) {
  tours := []struct {
      ID               int       `db:"id"`
      Start_Date       time.Time `db:"start_date"`
      End_Date         time.Time `db:"end_date"`
      Amount_Of_People string    `db:"amount_of_people"`
      UserID           int       `db:"user_id"`
      UserName         string    `db:"user_name"`
  }{}

  err := r.db.Select(&tours, `
      SELECT "tour".id, "tour".start_date, "tour".end_date, "tour".amount_of_people, "tour".user_id, "user".name AS user_name 
      FROM "tour" 
      JOIN "user" ON "tour".user_id = "user".id
  `)
  
  if err != nil {
      fmt.Println("get tours err", err)
      return nil, err
  }

  result := make([]model.Tour, len(tours))
  for i, t := range tours {
      result[i] = model.Tour{
          ID:               t.ID,
          Start_Date:       t.Start_Date,
          End_Date:         t.End_Date,
          Amount_Of_People: t.Amount_Of_People,
          User:             model.User{ID: t.UserID, Name: t.UserName},
      }
  }

  return result, nil
}


func (r *TourRepository) GetTour(id int) (*model.Tour, error) {
  dest := struct {
    ID               int       `db:"id"`
    Start_Date       time.Time `db:"start_date"`
    End_Date         time.Time `db:"end_date"`
    Amount_Of_People string    `db:"amount_of_people"`
    UserID           int       `db:"user_id"`
    UserName         string    `db:"user_name"`
  }{}
  
  err := r.db.Get(&dest, `
      SELECT "tour".id, "tour".start_date, "tour".end_date, "tour".amount_of_people, "tour".user_id, "user".name AS user_name 
      FROM "tour" 
      JOIN "user" ON "tour".user_id = "user".id 
      WHERE "tour".id = $1
  `, id)
  
  if err != nil {
      fmt.Println("get tour err", err)
      return nil, err
  }

  return &model.Tour{
      ID:               dest.ID,
      Start_Date:       dest.Start_Date,
      End_Date:         dest.End_Date,
      Amount_Of_People: dest.Amount_Of_People,
      User:             model.User{ID: dest.UserID, Name: dest.UserName},
  }, nil
}

func (r *TourRepository) CreateTour(startDate string, endDate string, amountOfPeople int, userID int) error {
  _, err := r.db.Exec(`
    INSERT INTO "tour" (start_date, end_date, amount_of_people, user_id)
    VALUES ($1, $2, $3, $4)
  `, startDate, endDate, amountOfPeople, userID)

  return err
}

func (r *TourRepository) DeleteTour(id int) error {
    _, err := r.db.Exec("DELETE FROM \"tour\" WHERE id=$1", id)
    
    if err != nil {
        fmt.Println("delete tour err", err)
        return err
    }

    return nil
}

func (r *TourRepository) UpdateTour(id int, startDate string, endDate string, amountOfPeople int, userID int) error {
  _, err := r.db.Exec(`
    UPDATE "tour"
    SET start_date = $1, end_date = $2, amount_of_people = $3
    WHERE id = $4
  `, startDate, endDate, amountOfPeople, userID)

  return err
}