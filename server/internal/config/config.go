package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

func Load() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Log        Log
	HTTPServer HTTPServer
	MySQL      MySQL
}

type MySQL struct {
	User     string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	DB       string `env:"MYSQL_DATABASE,required"`
	DBHost   string `env:"MYSQL_DB_HOST,required"`
	Port     int    `env:"MYSQL_PORT"`
}

func (cm *MySQL) GetDSN() string {
	fullHost := cm.DBHost
	if cm.Port > 0 {
		fullHost = fmt.Sprintf("%s:%d", cm.DBHost, cm.Port)
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cm.User, cm.Password, fullHost, cm.DB)
}

type Log struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type HTTPServer struct {
	Addr         string        `env:"HTTP_SERVER_ADDR"`
	Port         int           `env:"PORT"`
	WriteTimeout time.Duration `env:"HTTP_SERVER_WRITETIMEOUT" envDefault:"10s"`
	ReadTimeout  time.Duration `env:"HTTP_SERVER_READTIMEOUT" envDefault:"10s"`
}

func (s *HTTPServer) GetAddr() string {
	if s.Addr == "" {
		log.Fatalln("Empty HTTP_SERVER_ADDR")
	}

	if s.Port > 0 {
		return fmt.Sprintf(":%d", s.Port)
	}

	return s.Addr
}

func (cfg *Config) Fields() log.Fields {
	return log.Fields{
		"server_addr": cfg.HTTPServer.GetAddr(),
		"mysql_host":  cfg.MySQL.DBHost,
		"mysql_port":  cfg.MySQL.Port,
		"log_level":   cfg.Log.LogLevel,
	}
}
