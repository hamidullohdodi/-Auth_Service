package postgres

import (
	pb "auth_service/genproto/user"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_GetProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	mockProfile := &pb.ProfileResponse{
		Id:        "1",
		Username:  "testuser",
		Email:     "test@example.com",
		FullName:  "Test User",
		UserType:  "regular",
		CreatedAt: "2024-07-14T12:00:00Z",
		UpdatedAt: "2024-07-14T12:30:00Z",
	}

	mock.ExpectQuery("select id, username, email, full_name, user_type, created_at, updated_at from users where id = ?").
		WithArgs(mockProfile.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "full_name", "user_type", "created_at", "updated_at"}).
			AddRow(mockProfile.Id, mockProfile.Username, mockProfile.Email, mockProfile.FullName, mockProfile.UserType, mockProfile.CreatedAt, mockProfile.UpdatedAt))

	resultProfile, err := userRepo.GetProfile(&pb.Id{Id: mockProfile.Id})

	assert.NoError(t, err)
	assert.Equal(t, mockProfile.Id, resultProfile.Id)
	assert.Equal(t, mockProfile.Username, resultProfile.Username)
	assert.Equal(t, mockProfile.Email, resultProfile.Email)
	assert.Equal(t, mockProfile.FullName, resultProfile.FullName)
	assert.Equal(t, mockProfile.UserType, resultProfile.UserType)
	assert.Equal(t, mockProfile.CreatedAt, resultProfile.CreatedAt)
	assert.Equal(t, mockProfile.UpdatedAt, resultProfile.UpdatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	mockUpdateRequest := &pb.UpdateProfileRequest{
		Id:        "1",
		Username:  "updateduser",
		Email:     "updated@example.com",
		FullName:  "Updated User",
		UserType:  "admin",
		Bio:       "Updated bio",
		UpdatedAt: "2024-07-14T13:00:00Z",
	}

	mock.ExpectExec("UPDATE users SET").
		WithArgs(mockUpdateRequest.Username, mockUpdateRequest.Email, mockUpdateRequest.FullName, mockUpdateRequest.Bio, mockUpdateRequest.UserType, mockUpdateRequest.UpdatedAt, mockUpdateRequest.Id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate one row affected

	_, err = userRepo.UpdateProfile(mockUpdateRequest)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUserType(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	mockUpdateRequest := &pb.UpdateUserTypeRequest{
		Id:        "1",
		Username:  "updateduser",
		Email:     "updated@example.com",
		FullName:  "Updated User",
		Bio:       "Updated bio",
		UserType:  "admin",
		UpdatedAt: "2024-07-14T13:00:00Z",
	}

	mock.ExpectExec("UPDATE user SET").
		WithArgs(mockUpdateRequest.Username, mockUpdateRequest.Email, mockUpdateRequest.FullName, mockUpdateRequest.Bio, mockUpdateRequest.UserType, mockUpdateRequest.UpdatedAt, mockUpdateRequest.Id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate one row affected

	_, err = userRepo.UpdateUserType(mockUpdateRequest)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	mockGetRequest := &pb.GetUsersRequest{
		Page:  1,
		Limit: 10,
	}

	rows := sqlmock.NewRows([]string{"id", "username", "full_name", "user_type"}).
		AddRow("1", "user1", "User One", "regular").
		AddRow("2", "user2", "User Two", "admin")

	mock.ExpectQuery("SELECT id, username, full_name, user_type FROM users ORDER BY id LIMIT ? OFFSET ?").
		WithArgs(mockGetRequest.Limit, (mockGetRequest.Page-1)*mockGetRequest.Limit).
		WillReturnRows(rows)

	result, err := userRepo.GetUsers(mockGetRequest)

	assert.NoError(t, err)

	expectedUsers := []*pb.User{
		{Id: "1", Username: "user1", FullName: "User One", UserType: "regular"},
		{Id: "2", Username: "user2", FullName: "User Two", UserType: "admin"},
	}
	assert.Equal(t, expectedUsers, result.Users)
	assert.Equal(t, int32(2), result.Total)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	mockDeleteRequest := &pb.DeleteUserRequest{
		UserId: "1",
	}

	mock.ExpectExec("update users set deleted_at = EXTRACT(EPOCH FROM NOW()) where id = ?").
		WithArgs(mockDeleteRequest.UserId).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate one row affected

	_, err = userRepo.DeleteUser(mockDeleteRequest)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
