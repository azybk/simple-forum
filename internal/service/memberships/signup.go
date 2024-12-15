package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/azybk/simple-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username atau email sudah ada")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	userModel := memberships.UserModel{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, userModel)
}
