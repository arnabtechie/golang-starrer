package auth

import "errors"

func Register(data User) (RegResp, error) {
	if data.Password != data.ConfirmPassword {
		return RegResp{}, errors.New("passwords do not match")
	}

	return RegResp{
		ID:    "jweenfjw",
		Email: data.Email,
		Token: "jasncjsn",
	}, nil
}
