package config

import (
	"time"
	// "vicore_hrd/exception"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		// ErrorHandler:  exception.ErrorHandler,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Hospital Managemen System API MODUL HRD v1.0.1",
	}
}
