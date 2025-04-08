package services

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/internal/models"
	userMock "github.com/So-ham/Content-Moderation/internal/models/user/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_service_Signup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *entities.UserSignupRequest
	}

	os.Setenv("JWT_SECRET_KEY", "mockedSecretKey")

	// Setup mock for successful signup
	successMock := userMock.User{}
	successMock.On("FindByEmail", "test@example.com").Return(nil, nil) // User not found
	successMock.On("Create", mock.Anything, mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*entities.User)
		user.ID = 1 // Set ID as if it was created in DB
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
	})

	// Setup mock for user already exists
	existingUserMock := userMock.User{}
	existingUser := &entities.User{ID: 2, Email: "invalid-email", Username: "existinguser"}
	existingUserMock.On("FindByEmail", "invalid-email").Return(existingUser, nil)

	// Setup mock for empty username validation
	emptyUsernameMock := userMock.User{}
	emptyUsernameMock.On("FindByEmail", "test@example.com").Return(nil, nil)
	emptyUsernameMock.On("Create", mock.Anything, mock.AnythingOfType("*entities.User")).Return(errors.New("username cannot be empty"))

	// Setup mock for weak password validation
	weakPasswordMock := userMock.User{}
	weakPasswordMock.On("FindByEmail", "test@example.com").Return(nil, nil)
	weakPasswordMock.On("Create", mock.Anything, mock.AnythingOfType("*entities.User")).Return(errors.New("password too weak"))

	tests := []struct {
		name    string
		s       *service
		args    args
		want    *entities.UserResponse
		want1   string
		wantErr bool
	}{
		{
			name: "successful signup",
			s: &service{
				model: models.Model{
					User: &successMock,
				},
			},
			args: args{
				req: &entities.UserSignupRequest{
					Username: "testuser",
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			want: &entities.UserResponse{
				ID:       1,
				Username: "testuser",
				Email:    "test@example.com",
			},
			want1:   "", // We can't predict the actual token value in tests
			wantErr: false,
		},
		{
			name: "user already exists",
			s: &service{
				model: models.Model{
					User: &existingUserMock,
				},
			},
			args: args{
				req: &entities.UserSignupRequest{
					Username: "testuser",
					Email:    "invalid-email",
					Password: "password123",
				},
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
		{
			name: "empty username",
			s: &service{
				model: models.Model{
					User: &emptyUsernameMock,
				},
			},
			args: args{
				req: &entities.UserSignupRequest{
					Username: "",
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
		{
			name: "weak password",
			s: &service{
				model: models.Model{
					User: &weakPasswordMock,
				},
			},
			args: args{
				req: &entities.UserSignupRequest{
					Username: "testuser",
					Email:    "test@example.com",
					Password: "123",
				},
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Signup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// For successful case, we can't predict the exact token value
			if tt.name == "successful signup" {
				if got1 == "" {
					t.Errorf("service.Signup() expected non-empty token but got empty string")
				}
				// Compare only username and email, as CreatedAt will be different
				if got.Username != tt.want.Username || got.Email != tt.want.Email || got.ID != tt.want.ID {
					t.Errorf("service.Signup() got = %v, want = %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("service.Signup() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("service.Signup() got1 = %v, want %v", got1, tt.want1)
				}
			}
		})
	}
}

func Test_service_Login(t *testing.T) {
	type args struct {
		req *entities.UserLoginRequest
	}

	os.Setenv("JWT_SECRET_KEY", "mockedSecretKey")

	// Setup mock for successful login
	successMock := userMock.User{}
	validUser := &entities.User{
		ID:        1,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "$2a$10$1234567890123456789012", // Mocked bcrypt hash
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	successMock.On("FindByEmail", "test@example.com").Return(validUser, nil)

	// Setup mock for user not found
	userNotFoundMock := userMock.User{}
	userNotFoundMock.On("FindByEmail", "nonexistent@example.com").Return(nil, nil)

	// Setup mock for database error
	dbErrorMock := userMock.User{}
	dbErrorMock.On("FindByEmail", "test@example.com").Return(nil, errors.New("database connection error"))

	// Setup mock for invalid password
	invalidPasswordMock := userMock.User{}
	invalidPasswordUser := &entities.User{
		ID:        3,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "$2a$10$differenthashforpassword", // Different hash that won't match
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	invalidPasswordMock.On("FindByEmail", "test@example.com").Return(invalidPasswordUser, nil)

	tests := []struct {
		name    string
		s       *service
		args    args
		want    *entities.UserResponse
		want1   string
		wantErr bool
	}{
		{
			name: "successful login",
			s: &service{
				model: models.Model{
					User: &successMock,
				},
			},
			args: args{
				req: &entities.UserLoginRequest{
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			want: &entities.UserResponse{
				ID:       1,
				Username: "testuser",
				Email:    "test@example.com",
			},
			want1:   "", // We can't predict the actual token value in tests
			wantErr: false,
		},
		{
			name: "user not found",
			s: &service{
				model: models.Model{
					User: &userNotFoundMock,
				},
			},
			args: args{
				req: &entities.UserLoginRequest{
					Email:    "nonexistent@example.com",
					Password: "password123",
				},
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
		{
			name: "database error",
			s: &service{
				model: models.Model{
					User: &dbErrorMock,
				},
			},
			args: args{
				req: &entities.UserLoginRequest{
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// For successful case, we can't predict the exact token value
			if tt.name == "successful login" {
				if got1 == "" {
					t.Errorf("service.Login() expected non-empty token but got empty string")
				}
				// Compare only username and email, as CreatedAt will be different
				if got.Username != tt.want.Username || got.Email != tt.want.Email || got.ID != tt.want.ID {
					t.Errorf("service.Login() got = %v, want = %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("service.Login() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("service.Login() got1 = %v, want %v", got1, tt.want1)
				}
			}
		})
	}
}
