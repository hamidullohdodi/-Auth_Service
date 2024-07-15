package postgres

import (
	pb "auth_service/genproto/user"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strings"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetProfile(rep *pb.Id) (*pb.ProfileResponse, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.ProfileResponse{}, err
	}
	profile := &pb.ProfileResponse{}
	err = u.db.QueryRow("select id, username, email, full_name, user_type, created_at, updated_at from users where id = $1", rep.Id).Scan(profile.Id, &profile.Username, &profile.Email, &profile.FullName, &profile.UserType, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return &pb.ProfileResponse{}, err
	}
	return profile, nil
}

func (u *UserRepo) UpdateProfile(rep *pb.UpdateProfileRequest) (*pb.Void, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{}, err
	}
	query := ` UPDATE users SET `

	if rep.Username != "" {
		query += ` username = $2, `
	}
	if rep.Email != "" {
		query += ` email = $3, `
	}
	if rep.FullName != "" {
		query += ` full_name = $4, `
	}
	if rep.Bio != "" {
		query += ` bio = $5, `
	}
	if rep.UserType != "" {
		query += ` user_type = $6, `
	}
	if rep.UpdatedAt != "" {
		query += ` updated_at = $7, `
	}
	query = strings.TrimSuffix(query, ", ")
	query += ` WHERE id = $1 AND deleted_at = 0 `
	_, err = u.db.Exec(query,
		rep.Id,
		rep.Username,
		rep.Email,
		rep.FullName,
		rep.Bio,
		rep.UserType,
		rep.UpdatedAt,
	)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (u *UserRepo) UpdateUserType(rep *pb.UpdateUserTypeRequest) (*pb.Void, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{}, err
	}
	query := ` UPDATE user SET `

	if rep.Username != "" {
		query += ` username = $2, `
	}
	if rep.Email != "" {
		query += ` email = $3, `
	}
	if rep.FullName != "" {
		query += ` full_name = $4, `
	}
	if rep.Bio != "" {
		query += ` bio = $5, `
	}
	if rep.UserType != "" {
		query += ` user_type = $6, `
	}
	if rep.UpdatedAt != "" {
		query += ` updated_at = $7, `
	}
	query = strings.TrimSuffix(query, ", ")
	query += ` WHERE id = $1 AND deleted_at = 0 `
	_, err = u.db.Exec(query,
		rep.Id,
		rep.Username,
		rep.Email,
		rep.FullName,
		rep.Bio,
		rep.UserType,
		rep.UpdatedAt,
	)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (u *UserRepo) GetUsers(rep *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	offset := (rep.Page - 1) * rep.Limit
	rows, err := u.db.Query("SELECT id, username, full_name, user_type FROM users ORDER BY id LIMIT $1 OFFSET $2", rep.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*pb.User{}
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.Username, &user.FullName, &user.UserType)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	var total int32
	err = u.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return nil, err
	}

	response := &pb.GetUsersResponse{
		Users: users,
		Total: total,
		Page:  rep.Page,
		Limit: rep.Limit,
	}
	return response, nil
}

func (u *UserRepo) DeleteUser(rep *pb.DeleteUserRequest) (*pb.Void, error) {
	_, err := uuid.Parse(rep.UserId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{}, err
	}
	query :=
		`update users set deleted_at = EXTRACT(EPOCH FROM NOW()) where id = $1`
	_, err = u.db.Exec(query, rep.UserId)
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}
