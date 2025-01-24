package repository

import "backend/internal/entity"

func (r *userImpl) CreateUser(user entity.User) (entity.User, error) {
	r.users[user.ID] = user
	return user, nil
}

func (r *userImpl) FindUserByEmail(email string) (*entity.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
