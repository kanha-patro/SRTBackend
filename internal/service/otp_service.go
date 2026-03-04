package service

import (
	"errors"
	"time"

	"github.com/akpatri/srt/internal/otp"
	"github.com/akpatri/srt/internal/repository"
)

type otpService struct {
	otpRepo repository.OTPRepository
	otpGen  *otp.Generator
	otpVal  *otp.OTPValidator
}

func NewOTPService(otpRepo repository.OTPRepository, otpGen *otp.Generator, otpVal *otp.OTPValidator) *otpService {
	return &otpService{
		otpRepo: otpRepo,
		otpGen:  otpGen,
		otpVal:  otpVal,
	}
}

func (s *otpService) GenerateOTP(orgCode, routeCode, driverCode string) (string, error) {
	code, err := s.otpGen.Generate()
	if err != nil {
		return "", err
	}
	expiry := time.Now().Add(5 * time.Minute)
	return code, s.otpRepo.StoreOTP(orgCode, routeCode, driverCode, code, expiry)
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
