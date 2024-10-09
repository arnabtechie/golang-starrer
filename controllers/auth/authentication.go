package auth

import (
	"errors"

	"github.com/google/uuid"
)

func Register(data User) (RegResp, error) {
	if data.Password != data.ConfirmPassword {
		return RegResp{}, errors.New("passwords do not match")
	}

	return RegResp{
		ID:    uuid.New().String(),
		Email: data.Email,
		Token: "jasncjsn",
	}, nil
}
