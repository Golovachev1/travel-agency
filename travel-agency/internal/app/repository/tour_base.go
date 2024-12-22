package repository

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/Golovachev1/travel-agency/internal/app/model"
)

type Tour_BaseRepository struct {
    db *sqlx.DB
}

func NewTour_BaseRepository(db *sqlx.DB) *Tour_BaseRepository {
    return &Tour_BaseRepository{db: db}
}

func (r *Tour_BaseRepository) GetTourBases() ([]model.Tour_Base, error) {
	tours := []model.Tour_Base{}
	err := r.db.Select(&tours, "SELECT id, city_of_flight, arrival_country, duration, date_of_tour, tour_cost FROM \"tour_base\"")
	if err != nil {
	 fmt.Println("get tours err", err)
	 return nil, err
	}
	return tours, nil
   }


   func (r *Tour_BaseRepository) GetTourBase(id int) (*model.Tour_Base, error) {
	tour := model.Tour_Base{}
	err := r.db.Get(&tour, "SELECT id, city_of_flight, arrival_country, duration, date_of_tour, tour_cost FROM \"tour_base\" WHERE id=$1", id)
	if err != nil {
	 fmt.Println("get tour err", err)
	 return nil, err
	}
	return &tour, nil
   }

   func (r *Tour_BaseRepository) CreateTourBase(cityOfFlight, arrivalCountry string, duration int, dateOfTour string, tourCost float64, tourID int) error {
    _, err := r.db.Exec("INSERT INTO tour_base (city_of_flight, arrival_country, duration, date_of_tour, tour_cost, tour_id) VALUES ($1, $2, $3, $4, $5, $6)", cityOfFlight, arrivalCountry, duration, dateOfTour, tourCost, tourID)
    return err
}


func (r *Tour_BaseRepository) DeleteTourBase(id int) error {
    _, err := r.db.Exec("DELETE FROM \"tour_base\" WHERE id=$1", id)
    
    if err != nil {
        fmt.Println("delete tour_base err", err)
        return err
    }

    return nil
}

func (r *Tour_BaseRepository) UpdateTourBase(id int, cityOfFlight, arrivalCountry string, duration int, dateOfTour string, tourCost float64, tourID int) error {
    _, err := r.db.Exec(
        "UPDATE tour_base SET city_of_flight = $1, arrival_country = $2, duration = $3, date_of_tour = $4, tour_cost = $5, tour_id = $6 WHERE id = $7",
        cityOfFlight, arrivalCountry, duration, dateOfTour, tourCost, tourID, id,
    )
    return err
}
