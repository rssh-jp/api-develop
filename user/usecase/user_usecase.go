package usecase

import (
	"context"

	"github.com/rssh-jp/api-develop/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uu *userUsecase) Fetch(ctx context.Context) ([]domain.User, error) {
	return uu.userRepository.Fetch(ctx)
}
func (uu *userUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return uu.userRepository.GetByID(ctx, id)
}
func (uu *userUsecase) Update(ctx context.Context, user domain.User) error {
	return uu.userRepository.Update(ctx, user)
}
