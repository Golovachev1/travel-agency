package repository

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/Golovachev1/travel-agency/internal/app/model"
)

type ReservationRepository struct {
    db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) *ReservationRepository {
    return &ReservationRepository{db: db}
}

func (r *ReservationRepository) GetReservations() ([]model.Reservation, error) {
	reservations := []model.Reservation{}
	err := r.db.Select(&reservations, "SELECT id, reservation_date, payment FROM \"reservation\"")
	if err != nil {
	 fmt.Println("get reservations err", err)
	 return nil, err
	}
	return reservations, nil
   }


   func (r *ReservationRepository) GetReservation(id int) (*model.Reservation, error) {
	reservation := model.Reservation{}
	err := r.db.Get(&reservation, "SELECT id, reservation_date, payment FROM \"reservation\" WHERE id=$1", id)
	if err != nil {
	 fmt.Println("get reservation err", err)
	 return nil, err
	}
	return &reservation, nil
   }

   func (r *ReservationRepository) CreateReservation(Reservation_Date  string, Payment float64, tourID int) error {
    _, err := r.db.Exec("INSERT INTO reservation (reservation_date, payment, tour_id) VALUES ($1, $2, $3)", Reservation_Date, Payment, tourID)
    return err
}


func (r *ReservationRepository) DeleteReservation(id int) error {
    _, err := r.db.Exec("DELETE FROM \"reservation\" WHERE id=$1", id)
    
    if err != nil {
        fmt.Println("delete reservation err", err)
        return err
    }

    return nil
}

func (r *ReservationRepository) UpdateReservation(id int, Reservation_Date string, Payment float64, tourID int) error {
    _, err := r.db.Exec(
        "UPDATE Reservation SET reservation_date = $1, payment = $2, tour_id = $3 WHERE id = $4",
        Reservation_Date, Payment, tourID, id,
    )
    return err
}
