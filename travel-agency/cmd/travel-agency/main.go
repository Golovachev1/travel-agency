package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "gitlab.com/Golovachev1/travel-agency/docs"
	"github.com/Golovachev1/travel-agency/internal/app/handler"
	"github.com/Golovachev1/travel-agency/internal/app/repository"
	"github.com/Golovachev1/travel-agency/internal/app/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func GetUsers (userRepository *repository.UserRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userRepository.GetUsers()
		if len(users) == 0 {
			w.Write([]byte("Users are empty"))
			return
		}
		for _, user := range users {
			w.Write([]byte(fmt.Sprintf("\nName: %s, PhoneNumber: %s, Email: %s, Password: %s", user.Name, user.PhoneNumber, user.Email, user.Password)))
		}
	}
}

func CreateUser (userRepository *repository.UserRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := userRepository.CreateUser("Test", "89633204323", "gubanov@mail.ru", "***")
		if err != nil {
			w.Write([]byte("User can`t be created"))
			return
		}
		w.Write([]byte("Create User"))
	}
}

func GetUser (userRepository *repository.UserRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	   vars := mux.Vars(r)
	   userID, _ := strconv.Atoi(vars["id"])

	   user := userRepository.GetUser(userID)
	   if user == nil {
		   w.Write([]byte("User not found"))
		   return
	   }
	   w.Write([]byte(fmt.Sprintf("Name: %s, PhoneNumber: %s, Email: %s, Password: %s", user.Name, user.PhoneNumber, user.Email, user.Password)))
	   }
	
	
}

func DeleteUser (userRepository *repository.UserRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, _ := strconv.Atoi(vars["id"])

	    err := userRepository.DeleteUser(userID)
		if err != nil {
			w.Write([]byte("User can`t be deleted"))
			return
		}
		w.Write([]byte("Delete User"))
	}
}

func UpdateUser (userRepository *repository.UserRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	    err := userRepository.UpdateUser(2, "Максим Василевский", "89633204323", "gubanov@mail.ru", "***")
		if err != nil {
			w.Write([]byte("User can`t be updated"))
			return
		}
		w.Write([]byte("Update User"))
    }
}

func GetTours(tourRepository *repository.TourRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	 tours, err := tourRepository.GetTours()
	 if err != nil || len(tours) == 0 {
	  w.Write([]byte("Tours are empty or an error occurred"))
	  return
	 }
	 for _, tour_base := range tours {
	  w.Write([]byte(fmt.Sprintf("\nStart_Date: %s, End_Date: %s, Amount_Of_People: %s, User: %s", tour_base.Start_Date.Format(time.DateOnly), tour_base.End_Date.Format(time.DateOnly), tour_base.Amount_Of_People, tour_base.User.Name)))
	 }
	}
   }

   func CreateTour(tourRepository *repository.TourRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		  return
		}
		startDate := "2023-09-09 05:00:00"
		endDate := "2023-09-19 15:00:00"
		amountOfPeople := 5
		userID := 2
		err := tourRepository.CreateTour(startDate, endDate, amountOfPeople, userID)
		if err != nil {
		  http.Error(w, "Error creating tour", http.StatusInternalServerError)
		  return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Tour created successfully"))
	  }
   }

   func GetTour(tourRepository *repository.TourRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  vars := mux.Vars(r)
	  idStr := vars["id"]
	  id, err := strconv.Atoi(idStr)
	  if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	  }
  
	  tour, err := tourRepository.GetTour(id)
	  if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tour not found"))
		return
	  }
  
	  w.Write([]byte(fmt.Sprintf("ID: %d, Start_Date: %s, End_Date: %s, Amount_Of_People: %s, User ID: %d, User Name: %s",
		tour.ID, tour.Start_Date.Format(time.DateOnly), tour.End_Date.Format(time.DateOnly),
		tour.Amount_Of_People, tour.User.ID, tour.User.Name)))
	}
  }

func DeleteTour (tourRepository *repository.TourRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tourID, _ := strconv.Atoi(vars["id"])

	    err :=tourRepository.DeleteTour(tourID)
		if err != nil {
			w.Write([]byte("Tour can`t be deleted"))
			return
		}
		w.Write([]byte("Delete Tour"))
	}
}

func UpdateTour(tourRepository *repository.TourRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPatch {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        vars := mux.Vars(r)
        idStr := vars["id"]

        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid tour ID", http.StatusBadRequest)
            return
        }
        startDate := "2023-09-09 05:00:00" 
        endDate := "2023-09-19 15:00:00" 
        amountOfPeople := 7
        user := 1

        err = tourRepository.UpdateTour(id, startDate, endDate, amountOfPeople, user)
        if err != nil {
            http.Error(w, "Error updating tour", http.StatusInternalServerError)
            return
        }
        
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Tour updated successfully"))
    }
}

func GetTourBases(tourBaseRepository *repository.Tour_BaseRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	 tours, err := tourBaseRepository.GetTourBases()
	 if err != nil {
	  http.Error(w, "Unable to fetch tour bases", http.StatusInternalServerError)
	  return
	 }
	 if len(tours) == 0 {
	  w.Write([]byte("Tour bases are empty"))
	  return
	 }
	 for _, tour := range tours {
	  w.Write([]byte(fmt.Sprintf("\nCity of Flight: %s, Arrival Country: %s, Duration: %d, Date of Tour: %s, Tour Cost: %.2f", 
	   tour.City_Of_Flight, tour.Arrival_Country, tour.Duration, tour.Date_Of_Tour.Format(time.DateOnly), tour.Tour_Cost)))
	 }
	}
   }

func CreateTourBase (tourBaseRepository *repository.Tour_BaseRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        CityOfFlight := "Москва"
        ArrivalCountry := "Греция"
        Duration := 10
        DateOfTour := "2023-09-19 15:00:00"
        TourCost := 130000.00
        tourID := 1
        err := tourBaseRepository.CreateTourBase(CityOfFlight, ArrivalCountry, Duration, DateOfTour, TourCost, tourID)
        if err != nil {
            http.Error(w, "Error creating tour base", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("Tour base created successfully"))
    }
}

func GetTourBase (tourBaseRepository *repository.Tour_BaseRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tourID, _ := strconv.Atoi(vars["id"])
	  
		tour_base, err := tourBaseRepository.GetTourBase(tourID)
		if err != nil {
		 w.Write([]byte("Tour base not found"))
		 return
		}
		w.Write([]byte(fmt.Sprintf("City of Flight: %s, Arrival Country: %s, Duration: %d, Date of Tour: %s, Tour Cost: %.2f",
		 tour_base.City_Of_Flight, tour_base.Arrival_Country, tour_base.Duration, tour_base.Date_Of_Tour.Format(time.DateOnly), tour_base.Tour_Cost)))
	   }
}

func DeleteTourBase (tourBaseRepository *repository.Tour_BaseRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tourID, _ := strconv.Atoi(vars["id"])
	  
		err := tourBaseRepository.DeleteTourBase(tourID)
		if err != nil {
		 http.Error(w, "Tour base can't be deleted", http.StatusInternalServerError)
		 return
		}
		w.Write([]byte("Tour base deleted successfully"))
	   }
}

func UpdateTourBase(tourBaseRepository *repository.Tour_BaseRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPatch {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        vars := mux.Vars(r)
        idStr := vars["id"]

        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid tour ID", http.StatusBadRequest)
            return
        }
        cityOfFlight := "Санкт-Петербург"
        arrivalCountry := "Германия"
        duration := 10
        dateOfTour := "2023-09-09 05:00:00"
        tourCost := 130000.00
        tourID := 1
        err = tourBaseRepository.UpdateTourBase(id, cityOfFlight, arrivalCountry, duration, dateOfTour, tourCost, tourID)
        if err != nil {
            http.Error(w, "Error updating tour base", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Tour base updated successfully"))
    }
}

func GetReservations (reservationRepository *repository.ReservationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reservations, err := reservationRepository.GetReservations()
		if err != nil {
		 http.Error(w, "Unable to fetch reservations", http.StatusInternalServerError)
		 return
		}
		if len(reservations) == 0 {
		 w.Write([]byte("Reservations are empty"))
		 return
		}
		for _, reservation := range reservations {
		 w.Write([]byte(fmt.Sprintf("\nReservation_Date: %s, Payment: %.2f", 
		 reservation.Reservation_Date.Format(time.DateOnly), reservation.Payment)))
		}
	   }
}

func CreateReservation (reservationRepository *repository.ReservationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        Reservation_Date := "2023-09-01 15:00:00"
        Payment := 130000.00
        tourID := 1
        err := reservationRepository.CreateReservation(Reservation_Date, Payment, tourID)
        if err != nil {
            http.Error(w, "Error creating reservation", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("Reservation created successfully"))
    }
}

func GetReservation (reservationRepository *repository.ReservationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tourID, _ := strconv.Atoi(vars["id"])
	  
		reservation, err := reservationRepository.GetReservation(tourID)
		if err != nil {
		 w.Write([]byte("Reservation not found"))
		 return
		}
		w.Write([]byte(fmt.Sprintf("Reservation_Date: %s, Payment: %.2f",
		 reservation.Reservation_Date.Format(time.DateOnly), reservation.Payment)))
	   }
}

func DeleteReservation (reservationRepository *repository.ReservationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tourID, _ := strconv.Atoi(vars["id"])

	    err :=reservationRepository.DeleteReservation(tourID)
		if err != nil {
			w.Write([]byte("Reservation can`t be deleted"))
			return
		}
		w.Write([]byte("Delete Reservation"))
	}
}

func UpdateReservation (reservationRepository *repository.ReservationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPatch {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        vars := mux.Vars(r)
        idStr := vars["id"]

        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid tour ID", http.StatusBadRequest)
            return
        }
        Reservation_Date := "2023-09-02 15:00:00"
        Payment := 130000.00
        tourID := 1
        err = reservationRepository.UpdateReservation(id, Reservation_Date, Payment, tourID)
        if err != nil {
            http.Error(w, "Error updating reservation", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Reservation updated successfully"))
    }
}

func GetReviews (reviewRepository *repository.ReviewRepository) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		reviews, err := reviewRepository.GetReviews()
		if err != nil {
		 http.Error(w, "Unable to fetch reviews", http.StatusInternalServerError)
		 return
		}
		if len(reviews) == 0 {
		 w.Write([]byte("Reviews are empty"))
		 return
		}
		for _, review := range reviews {
		 w.Write([]byte(fmt.Sprintf("\nScore: %d, Review_text: %s, Publish_date: %s", 
		 review.Score, review.Review_text, review.Publish_date.Format(time.DateOnly))))
		}
	   }
}

func CreateReview (reviewRepository *repository.ReviewRepository) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
		Score := 5
		Review_text := "Отличное агентство! Хороший персонал! Всё быстро, качественно"
        Publish_date := "2023-09-20 12:00:00"
		userID := 1
        tourID := 1
        err := reviewRepository.CreateReview(Score, Review_text, Publish_date, userID, tourID)
        if err != nil {
            http.Error(w, "Error creating review", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("Review created successfully"))
    }
}

func GetReview (reviewRepository *repository.ReviewRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, _ := strconv.Atoi(vars["id"])
	  
		review, err := reviewRepository.GetReview(userID)
		if err != nil {
		 w.Write([]byte("Review not found"))
		 return
		}
		w.Write([]byte(fmt.Sprintf("Score: %d, Review_text: %s, Publish_date: %s, User: %s",
		review.Score, review.Review_text, review.Publish_date.Format(time.DateOnly), review.User.Name)))
	   }
}

func DeleteReview (reviewRepository *repository.ReviewRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, _ := strconv.Atoi(vars["id"])

	    err := reviewRepository.DeleteReview(userID)
		if err != nil {
			w.Write([]byte("Review can`t be deleted"))
			return
		}
		w.Write([]byte("Delete Review"))
	}
}

func UpdateReview(reviewRepository *repository.ReviewRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPatch {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        vars := mux.Vars(r)
        idStr := vars["id"]
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid review ID", http.StatusBadRequest)
            return
        }
        Score := 3
        Review_text := "Хорошее агентство, но скорость обслуживания не очень высокая"
        Publish_date := "2023-09-25 09:00:00"
        userID := 1
        tourID := 1
        err = reviewRepository.UpdateReview(id, Score, Review_text, Publish_date, userID, tourID)
        if err != nil {
            http.Error(w, "Error updating review", http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Review updated successfully"))
    }
}


func main () {
    db, err := sqlx.Connect("postgres", "user=dev dbname=travel_agency sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	tourRepository := repository.NewTourRepository(db)
	tour_baseRepository := repository.NewTour_BaseRepository(db)
	reservationRepository := repository.NewReservationRepository(db)
	reviewRepository := repository.NewReviewRepository(db)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", DeleteUser(userRepository)).Methods(http.MethodDelete)
	r.HandleFunc("/users/{id}", UpdateUser(userRepository)).Methods(http.MethodPatch)

	r.HandleFunc("/tours", GetTours(tourRepository)).Methods(http.MethodGet)
	r.HandleFunc("/tours", CreateTour(tourRepository)).Methods(http.MethodPost)
	r.HandleFunc("/tours/{id}", GetTour(tourRepository)).Methods(http.MethodGet)
	r.HandleFunc("/tours/{id}", DeleteTour(tourRepository)).Methods(http.MethodDelete)
	r.HandleFunc("/tours/{id}", UpdateTour(tourRepository)).Methods(http.MethodPatch)

	r.HandleFunc("/tourBases", GetTourBases(tour_baseRepository)).Methods(http.MethodGet)
	r.HandleFunc("/tourBases", CreateTourBase(tour_baseRepository)).Methods(http.MethodPost)
	r.HandleFunc("/tourBases/{id}", GetTourBase(tour_baseRepository)).Methods(http.MethodGet)
	r.HandleFunc("/tourBases/{id}", DeleteTourBase(tour_baseRepository)).Methods(http.MethodDelete)
	r.HandleFunc("/tourBases/{id}", UpdateTourBase(tour_baseRepository)).Methods(http.MethodPatch)

	r.HandleFunc("/reservations", GetReservations(reservationRepository)).Methods(http.MethodGet)
	r.HandleFunc("/reservations", CreateReservation(reservationRepository)).Methods(http.MethodPost)
	r.HandleFunc("/reservations/{id}", GetReservation(reservationRepository)).Methods(http.MethodGet)
	r.HandleFunc("/reservations/{id}", DeleteReservation(reservationRepository)).Methods(http.MethodDelete)
	r.HandleFunc("/reservations/{id}", UpdateReservation(reservationRepository)).Methods(http.MethodPatch)

	r.HandleFunc("/reviews", GetReviews(reviewRepository)).Methods(http.MethodGet)
	r.HandleFunc("/reviews", CreateReview(reviewRepository)).Methods(http.MethodPost)
	r.HandleFunc("/reviews/{id}", GetReview(reviewRepository)).Methods(http.MethodGet)
	r.HandleFunc("/reviews/{id}", DeleteReview(reviewRepository)).Methods(http.MethodDelete)
	r.HandleFunc("/reviews/{id}", UpdateReview(reviewRepository)).Methods(http.MethodPatch)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}