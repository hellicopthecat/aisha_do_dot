package service

import "context"

func (s UserService) FindUserByEmailService(ctx context.Context, email string) (*string, error) {
	userMail, err := s.userService.FindUserByEmailRepo(ctx, email)
	if err != nil {
		return nil, err
	}
	return userMail, nil
}
