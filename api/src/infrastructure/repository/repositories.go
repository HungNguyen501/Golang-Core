package repository

import "golang-core/api/src/infrastructure/database"

type Repositories struct {
	UserRepository UserRepository
}

func NewRopositories(db *database.Db) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
