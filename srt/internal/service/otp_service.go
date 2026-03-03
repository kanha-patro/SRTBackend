package service

import (
	"errors"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/otp"
	"github.com/akpatri/srt/internal/repository"
)

type OTPService interface {
	GenerateOTP(orgCode, routeCode, driverCode string) (string, error)
	ValidateOTP(orgCode, routeCode, driverCode, otpCode string) (bool, error)
}

type otpService struct {
	otpRepo repository.OTPRepository
	otpGen   otp.Generator
	otpVal   otp.Validator
}

func NewOTPService(otpRepo repository.OTPRepository, otpGen otp.Generator, otpVal otp.Validator) OTPService {
	return &otpService{
		otpRepo: otpRepo,
		otpGen:  otpGen,
		otpVal:  otpVal,
	}
}

func (s *otpService) GenerateOTP(orgCode, routeCode, driverCode string) (string, error) {
	otpCode := s.otpGen.Generate()
	expiry := time.Now().Add(5 * time.Minute) // OTP valid for 5 minutes

	err := s.otpRepo.StoreOTP(orgCode, routeCode, driverCode, otpCode, expiry)
	if err != nil {
		return "", err
	}

	return otpCode, nil
}

func (s *otpService) ValidateOTP(orgCode, routeCode, driverCode, otpCode string) (bool, error) {
	otpData, err := s.otpRepo.GetOTP(orgCode, routeCode, driverCode)
	if err != nil {
		return false, err
	}

	if otpData == nil {
		return false, errors.New("OTP not found or expired")
	}

	if time.Now().After(otpData.Expiry) {
		return false, errors.New("OTP has expired")
	}

	if otpData.Code != otpCode {
		return false, errors.New("Invalid OTP")
	}

	return true, nil
}