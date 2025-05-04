// internal/auth/supabase.go
package auth

import (
	"context"
	"os"

	svc "nookli/pkg/service/auth"

	authgo "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

// supabaseService implements svc.Service using auth-go
type supabaseService struct {
	client authgo.Client
}

// NewSupabaseService constructs the Supabase Auth service.
func NewSupabaseService() svc.Service {
	ref := os.Getenv("SUPABASE_PROJECT_REF")
	key := os.Getenv("SUPABASE_KEY")
	client := authgo.New(ref, key)
	return &supabaseService{client: client}
}

func (s *supabaseService) Signup(_ctx context.Context, email, password string) (*svc.User, error) {
	resp, err := s.client.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	u := resp.User
	return &svc.User{
		ID:    u.ID.String(),
		Email: u.Email,
	}, nil
}

func (s *supabaseService) Signin(_ctx context.Context, email, password string) (*svc.User, string, string, error) {
	resp, err := s.client.Token(types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return nil, "", "", err
	}
	u := resp.User
	return &svc.User{
		ID:    u.ID.String(),
		Email: u.Email,
	}, resp.AccessToken, resp.RefreshToken, nil
}

func (s *supabaseService) Refresh(_ctx context.Context, refreshToken string) (string, string, error) {
	resp, err := s.client.Token(types.TokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	})
	if err != nil {
		return "", "", err
	}
	return resp.AccessToken, resp.RefreshToken, nil
}

func (s *supabaseService) Logout(_ctx context.Context, _refreshToken string) error {
	// auth-go currently does not expose a Logout endpoint
	return nil
}

func (s *supabaseService) ResetPassword(_ctx context.Context, email string) error {
	// Recover sends the recovery email and returns only error
	return s.client.Recover(types.RecoverRequest{Email: email})
}

func (s *supabaseService) GetUser(_ctx context.Context, accessToken string) (*svc.User, error) {
	// Create a client bound to the user’s token:
	c := s.client.WithToken(accessToken)
	resp, err := c.GetUser()
	if err != nil {
		return nil, err
	}
	u := resp.User
	return &svc.User{
		ID:    u.ID.String(),
		Email: u.Email,
	}, nil
}
