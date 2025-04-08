package usecase

import (
    "backClarineHetic/internal/domain"
)

type ProfileUseCase interface {
    SearchProfile(searchTerm string) ([]*domain.User, error)
}

type profileUseCase struct {
    userRepo domain.UserRepository
}

// NewProfileUseCase crée une nouvelle instance de ProfileUseCase.
func NewProfileUseCase(userRepo domain.UserRepository) ProfileUseCase {
    return &profileUseCase{
        userRepo: userRepo,
    }
}

func (p *profileUseCase) SearchProfile(searchTerm string) ([]*domain.User, error) {
    return p.userRepo.SearchProfiles(searchTerm)
}
