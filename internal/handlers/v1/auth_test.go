package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/internal/services/mocks"

	"github.com/stretchr/testify/mock"

	"github.com/go-playground/validator/v10"
)

func generateRequest(method string, url string, body any) *http.Request {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

func Test_handlerV1_SignupHandler(t *testing.T) {
	validate := validator.New()

	validRequest := entities.UserSignupRequest{
		Username: "soham",
		Email:    "example@gmail.com",
		Password: "password",
	}

	invalidRequest := validRequest
	invalidRequest.Email = ""

	vResp := entities.UserResponse{
		ID:       1,
		Username: "soham",
		Email:    "example@gmail.com",
	}

	mockService := &mocks.Service{}
	mockService.On("Signup", mock.Anything, mock.AnythingOfType("*entities.UserSignupRequest")).Return(&vResp, "123", nil)

	mockServiceErr := &mocks.Service{}
	mockServiceErr.On("Signup", mock.Anything, mock.AnythingOfType("*entities.UserSignupRequest")).Return(nil, "", errors.New("some service error"))

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *handlerV1
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid request",
			h: &handlerV1{
				Service:  mockService,
				Validate: validate,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest(http.MethodPost, "auth/signup", validRequest).WithContext(context.Background()),
			},
			wantErr: false,
		},

		{
			name: "json decode error",
			h: &handlerV1{
				Service:  mockService,
				Validate: validate,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest(http.MethodPost, "auth/signup", "not json format body"),
			},
			wantErr: true,
		},
		{
			name: "validation failed",
			h: &handlerV1{
				Service:  mockService,
				Validate: validate,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest(http.MethodPost, "auth/signup", invalidRequest).WithContext(context.Background()),
			},
			wantErr: true,
		},
		{
			name: "service error",
			h: &handlerV1{
				Service:  mockServiceErr,
				Validate: validate,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest(http.MethodPost, "auth/signup", validRequest).WithContext(context.Background()),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.SignupHandler(tt.args.w, tt.args.r)
		})
	}
}
