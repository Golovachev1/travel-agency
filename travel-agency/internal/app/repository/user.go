package repository

import (
  "fmt"
  "github.com/jmoiron/sqlx"
  "github.com/Golovachev1/travel-agency/internal/app/model"
)

type UserRepository struct {
  db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository { 
  return &UserRepository{db: db} 
}

func (r *UserRepository) GetUsers() []model.User {
  users := []model.User{}
  err := r.db.Select(&users, "SELECT name, phonenumber, email, password FROM \"user\"")
  if err != nil {
    fmt.Println("get users err", err)
    return users
  }

  return users
}

const getUserQuery = `
SELECT "user".id, "user".name, "user".phonenumber, "user".email, "user".password FROM "user" WHERE "user".id=$1
`

func (r *UserRepository) GetUser(id int) *model.User {
  dest := struct {
    ID int
    Name string
    PhoneNumber string
    Email string
    Password string
  }{}
  err := r.db.Get(&dest, getUserQuery, id)
  if err != nil {
    fmt.Println("get user err", err)
    return nil
  }

  return &model.User{
    ID: dest.ID,
    Name: dest.Name,
    PhoneNumber: dest.PhoneNumber,
    Email: dest.Email,
    Password: dest.Password,
  }
}

func (r *UserRepository) CreateUser(name string, phoneNumber string, email string, password string) error {
  user := []model.User{
    {Name: name, PhoneNumber: phoneNumber, Email: email, Password: password},
    {Name: "Павел Конев", PhoneNumber: "89654745557", Email: "pavel777@mail.ru", Password: "********"},
  }

  _, err := r.db.NamedExec("INSERT INTO \"user\" (name, phonenumber, email, password) VALUES (:name, :phonenumber, :email, :password)", user)
  if err != nil {
    return err 
  }

  return nil
}

func (r *UserRepository) DeleteUser(id int) error {
  _, err := r.db.Exec("DELETE FROM \"user\" WHERE id=$1", id)
  if err != nil {
      fmt.Println("delete user err", err)
      return err
  }

  return nil
}

func (r *UserRepository) UpdateUser(id int, name string, phoneNumber string, email string, password string) error {
  _, err := r.db.NamedExec(`UPDATE "user" SET name=:name, phonenumber=:phonenumber, email=:email, password=:password WHERE id=:id`, 
      map[string]interface{}{
        "id":          id,
        "name":        name,
        "phonenumber": phoneNumber,
        "email":       email,
        "password":    password,
      })
  if err != nil {
    return err
  }
  return nil
}