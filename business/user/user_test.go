package user_test

import (
	"errors"
	"testing"
	"time"

	"hexagonalArchitecture/business/user"
)

// MockRepository implements user.Repository for testing
type MockRepository struct {
	FindUserByIDFunc func(id int) (*user.FindUser, error)
}

func (m *MockRepository) FindUserByID(id int) (*user.FindUser, error) {
	return m.FindUserByIDFunc(id)
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		userName string
		email    string
		phone    string
		username string
		password string
		wantErr  bool
	}{
		{"Valid User", 1, "Tio", "tio@example.com", "123456", "tiodev", "pass123", false},
		{"Invalid Email", 1, "Tio", "invalid-email", "123456", "tiodev", "pass123", true},
		{"Invalid Username Short", 1, "Tio", "tio@example.com", "123456", "ti", "pass123", true},
		{"Invalid Username Special", 1, "Tio", "tio@example.com", "123456", "tio!", "pass123", true},
		{"Empty Email", 1, "Tio", "", "123456", "tiodev", "pass123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := user.NewUser(tt.id, tt.userName, tt.email, tt.phone, tt.username, "pass", time.Now(), time.Now())
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModifyUser(t *testing.T) {
	now := time.Now()
	original := &user.FindUser{
		ID: 1, Name: "Tio", Email: "tio@example.com", PhoneNumber: "123", Username: "tiodev",
		CreatedAt: now, UpdatedAt: now,
	}

	tests := []struct {
		name    string
		newName string
		newPhone string
		wantName string
		wantPhone string
	}{
		{"Update Both", "Thio", "456", "Thio", "456"},
		{"Update Name Only", "Thio", "", "Thio", "123"},
		{"Update Phone Only", "", "456", "Tio", "456"},
		{"Update None", "", "", "Tio", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := original.ModifyUser(tt.newName, tt.newPhone, time.Now())
			if got.Name != tt.wantName || got.PhoneNumber != tt.wantPhone {
				t.Errorf("ModifyUser() = %v, %v; want %v, %v", got.Name, got.PhoneNumber, tt.wantName, tt.wantPhone)
			}
		})
	}
}

func TestUserService_FindUserByID(t *testing.T) {
	tests := []struct {
		name    string
		mockFn  func(id int) (*user.FindUser, error)
		wantNil bool
		wantErr bool
	}{
		{
			name: "Success",
			mockFn: func(id int) (*user.FindUser, error) {
				return &user.FindUser{ID: id, Name: "Tio"}, nil
			},
			wantNil: false,
			wantErr: false,
		},
		{
			name: "User Not Found",
			mockFn: func(id int) (*user.FindUser, error) {
				return nil, user.ErrUserNotFound
			},
			wantNil: true,
			wantErr: true,
		},
		{
			name: "Internal DB Error",
			mockFn: func(id int) (*user.FindUser, error) {
				return nil, errors.New("critical db failure")
			},
			wantNil: true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepository{FindUserByIDFunc: tt.mockFn}
			svc := user.NewService(repo)
			got, err := svc.FindUserByID(1)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("FindUserByID() got = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}
