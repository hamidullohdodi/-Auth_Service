package postgres

import (
	pb "auth_service/genproto/auth"
	"database/sql"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) Register(auth *pb.RegisterRequest) (*pb.Void, error) {
	query := `INSERT INTO users (username, email, full_name, user_type) VALUES ($1, $2, $3, $4)`
	_, err := a.db.Exec(query, auth.Username, auth.Email, auth.FullName, auth.UserType)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}
