package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	DatabaseURL         string `mapstructure:"DATABASE_URL"`
	RedisURL            string `mapstructure:"REDIS_URL"`
	NatsURL             string `mapstructure:"NATS_URL"`
	OTPExpiryDuration    int    `mapstructure:"OTP_EXPIRY_DURATION"`
	TripIdleTimeout      int    `mapstructure:"TRIP_IDLE_TIMEOUT"`
	HeartbeatTimeout     int    `mapstructure:"HEARTBEAT_TIMEOUT"`
	LocationRateLimit    int    `mapstructure:"LOCATION_RATE_LIMIT"`
	AutoEndThreshold     int    `mapstructure:"AUTO_END_THRESHOLD"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("app") 
	viper.SetConfigType("yaml") 
	viper.AddConfigPath(".")     
	viper.AddConfigPath("config") 

	viper.AutomaticEnv() 

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}