package service

import (
	pb "auth_service/genproto/user"
	"auth_service/storage/postgres"
	"context"
	"database/sql"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Repo *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		Repo: postgres.NewUserRepo(db),
	}
}
func (u *UserService) GetProfile(ctx context.Context, rep *pb.Id) (*pb.ProfileResponse, error) {
	profil, err := u.Repo.GetProfile(rep)
	if err != nil {
		return &pb.ProfileResponse{}, err
	}
	return profil, nil
}

func (u *UserService) UpdateProfile(ctx context.Context, rep *pb.UpdateProfileRequest) (*pb.Void, error) {
	_, err := u.Repo.UpdateProfile(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (u *UserService) UpdateUserType(ctx context.Context, rep *pb.UpdateUserTypeRequest) (*pb.Void, error) {
	_, err := u.Repo.UpdateUserType(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (u *UserService) GetUsers(ctx context.Context, rep *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := u.Repo.GetUsers(rep)
	if err != nil {
		return &pb.GetUsersResponse{}, err
	}
	return users, nil
}

func (u *UserService) DeleteUser(ctx context.Context, rep *pb.DeleteUserRequest) (*pb.Void, error) {
	_, err := u.Repo.DeleteUser(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}
