package repository

import (
	"context"
	"database/sql"

	"github.com/shoelfikar/finpay-realtime-transaction/model"
)

type userRepository struct {
	DB *sql.DB
}

type UserRepository interface {
   CreateUser(user *model.User) *model.User
   GetUserByID(userId int) *model.User
   GetUserByEmail(email string) *model.User
}

func NewUserRepository(db *sql.DB) UserRepository {
   return &userRepository{
      DB: db,
   }
}

func (u *userRepository) CreateUser(user *model.User) *model.User {
   tx, err := u.DB.Begin()
   ctx := context.Background()
   if err != nil {
      panic("Error database transaction"+ err.Error())
   }

   defer tx.Commit()
   query := `
      INSERT INTO users (usernmae, email, password, status, created_by)
      VALUES ($1, $2, $3, $4, $5)
      RETURNING id;
   `
   result, err := tx.ExecContext(ctx, query, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedBy)

   if err != nil {
      panic("Error insert user"+ err.Error())
   }

   id, err := result.LastInsertId()
   if err != nil {
      panic("Error get user id"+ err.Error())
   }

   user.Id = int(id)

   return user
}

func (u *userRepository) GetUserByID(userId int) *model.User {
   ctx := context.Background()
   query := `
      SELECT * FROM users WHERE id = $1
   `

   var user model.User
   result := u.DB.QueryRowContext(ctx, query, &userId).Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy)

   if result == sql.ErrNoRows {
      return nil
   }

   return &user
}

func (u *userRepository) GetUserByEmail(email string) *model.User {
   ctx := context.Background()
   query := `
      SELECT * FROM users WHERE email = $1
   `

   var user model.User
   result := u.DB.QueryRowContext(ctx, query, &email).Scan(&user.Username, &user.Email, &user.Password, &user.Status, &user.Role, &user.Id, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy)

   if result == sql.ErrNoRows {
      return nil
   }

   return &user
}