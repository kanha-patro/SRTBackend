package service

import (
	"errors"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
)

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
	org, err := s.orgRepo.GetByID(orgID)
	if err != nil {
		return err
	}
	org.Activate()
	return s.orgRepo.Update(org)
}

func (s *orgService) SuspendOrg(orgID string) error {
	org, err := s.orgRepo.GetByID(orgID)
	if err != nil {
		return err
	}
	org.Deactivate()
	return s.orgRepo.Update(org)
}

func (s *orgService) GetOrg(orgID string) (*domain.Org, error) {
	return s.orgRepo.GetByID(orgID)
}

// GetActiveTrips returns active trips for the organization. (Not yet implemented)
func (s *orgService) GetActiveTrips() ([]*domain.Trip, error) {
	return nil, errors.New("GetActiveTrips not implemented")
}

// ForceStopTrip forcefully ends a trip. (Not yet implemented)
func (s *orgService) ForceStopTrip(tripID string) error {
	return errors.New("ForceStopTrip not implemented")
}

// RevokeOTPSession revokes a driver OTP session. (Not yet implemented)
func (s *orgService) RevokeOTPSession(sessionID string) error {
	return errors.New("RevokeOTPSession not implemented")
}

// GetActiveOrgs returns orgs that are currently active.
func (s *orgService) GetActiveOrgs() ([]domain.Org, error) {
	orgs, err := s.orgRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var active []domain.Org
	for _, o := range orgs {
		if o.IsActive {
			active = append(active, o)
		}
	}
	return active, nil
}

// UpdateOrg updates an organization's details.
func (s *orgService) UpdateOrg(orgID string, org *domain.Org) error {
	if org == nil {
		return errors.New("org cannot be nil")
	}
	existing, err := s.orgRepo.GetByID(orgID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("organization not found")
	}
	existing.Update(org.Name, org.Code)
	return s.orgRepo.Update(existing)
}
