package jwt

import (
	"github.com/AlexMinsk2017/PetAutchTest/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	type args struct {
		user     models.User
		app      models.App
		duration time.Duration
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("passw0rd"), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test case 1",
			args: args{
				user:     models.User{1, "asadsda@aasad", passwordHash},
				app:      models.App{1, "user", "123"},
				duration: 100,
			},
			want:    "ожидаемый токен для теста 1",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToken(tt.args.user, tt.args.app, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
