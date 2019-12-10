package utils

import (
	"os"

	"github.com/nats-io/nats.go"
)

// NATS
var (
	NATSURL = os.Getenv("SERVER_URL")
)

// GetURL -
func GetURL() string {
	if NATSURL == "" {
		return nats.DefaultURL
	}
	return NATSURL
}
