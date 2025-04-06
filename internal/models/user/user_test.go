package user

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/So-ham/Content-Moderation/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock() (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening gorm stub database connection", err)
	}
	return gormDB, db, mock
}

func Test_user_Create(t *testing.T) {
	gdb, db, mock := NewMock()
	defer db.Close()

	valid := entities.User{}

	validUser := entities.User{

		ID:        123,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},

		Username: "firstname",
		Email:    "email@test.com",
		Password: "123",
	}
	invalidUser := validUser
	invalidUser.Email = "notfound@test.com"

	// rows as result
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "username", "email", "password"}).
		AddRow(validUser.ID, validUser.CreatedAt, validUser.UpdatedAt, validUser.Username, validUser.Email, validUser.Password)

	// statements
	validstmt := gdb.Session(&gorm.Session{DryRun: true}).Model(&valid).Create(&validUser).Statement.SQL.String()
	invalidstmt := gdb.Session(&gorm.Session{DryRun: true}).Model(&valid).Create(&invalidUser).Statement.SQL.String()

	// expectmocks
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(validstmt)).WithArgs().WillReturnRows(rows)
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(invalidstmt)).WithArgs().WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()
	type args struct {
		ctx  context.Context
		user *entities.User
	}
	tests := []struct {
		name    string
		m       *user
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "valid case",
			m:       &user{DB: gdb},
			args:    args{ctx: context.Background(), user: &validUser},
			wantErr: false,
		},
		{
			name:    "insert error",
			m:       &user{DB: gdb},
			args:    args{ctx: context.Background(), user: &invalidUser},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Create(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
