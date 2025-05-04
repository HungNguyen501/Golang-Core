package manager

import "golang-core/api/src/infrastructure/repository"

type Managers struct {
	UserManager UserManager
}

func NewManagers(repo repository.Repositories) *Managers {
	userManager := NewUserManager(repo.UserRepository)
	return &Managers{
		UserManager: userManager,
	}
}
