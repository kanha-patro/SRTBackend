package service

import (
	"errors"
	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
)

type OrgService interface {
	RegisterOrg(org *domain.Org) error
	ApproveOrg(orgID string) error
	SuspendOrg(orgID string) error
	GetOrg(orgID string) (*domain.Org, error)
}

type orgService struct {
	orgRepo repository.OrgRepository
}

func NewOrgService(orgRepo repository.OrgRepository) OrgService {
	return &orgService{
		orgRepo: orgRepo,
	}
}

func (s *orgService) RegisterOrg(org *domain.Org) error {
	if org == nil {
		return errors.New("organization cannot be nil")
	}
	return s.orgRepo.Create(org)
}

func (s *orgService) ApproveOrg(orgID string) error {
	org, err := s.orgRepo.FindByID(orgID)
	if err != nil {
		return err
	}
	org.Status = domain.OrgStatusApproved
	return s.orgRepo.Update(org)
}

func (s *orgService) SuspendOrg(orgID string) error {
	org, err := s.orgRepo.FindByID(orgID)
	if err != nil {
		return err
	}
	org.Status = domain.OrgStatusSuspended
	return s.orgRepo.Update(org)
}

func (s *orgService) GetOrg(orgID string) (*domain.Org, error) {
	return s.orgRepo.FindByID(orgID)
}