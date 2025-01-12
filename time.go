package j8a

import (
	"os"

	"github.com/rs/zerolog/log"
)

var TZ = "UTC"

func initTime() string {
	tz := os.Getenv("TZ")
	if len(tz) == 0 {
		tz = "UTC"
	}
	log.Info().Str("timeZone", tz).Msg("timeZone determined")
	TZ = tz
	return tz
}
