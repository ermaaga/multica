package selfhosttelemetry

import (
	"log/slog"
	"strings"
)

// Config is read once during API server startup. Endpoint selection is
// deliberately absent: self-host telemetry can only be sent to Multica's
// first-party collector compiled into the client.
type Config struct {
	Enabled bool
}

// ConfigFromOptIn enables self-host telemetry only when the operator explicitly
// sets MULTICA_SELF_HOST_TELEMETRY to 1 or true. Telemetry is off by default:
// nothing leaves a self-hosted deployment unless the operator asks for it.
func ConfigFromOptIn(raw string) Config {
	raw = strings.TrimSpace(raw)
	enabled := raw == "1" || strings.EqualFold(raw, "true")
	return Config{Enabled: enabled}
}

// ConfigFromDoNotTrack implements the industry-standard opt-out switch used by
// the self-hosted server. Only 1 and true (case-insensitive), after trimming,
// disable telemetry; every other value preserves the default-on behavior.
func ConfigFromDoNotTrack(raw string) Config {
	raw = strings.TrimSpace(raw)
	disabled := raw == "1" || strings.EqualFold(raw, "true")
	return Config{Enabled: !disabled}
}

// LogStartupStatus makes the default-off behavior observable without logging
// the raw environment value or any telemetry data.
func LogStartupStatus(logger *slog.Logger, config Config) {
	if config.Enabled {
		logger.Info("self-host telemetry enabled (MULTICA_SELF_HOST_TELEMETRY)")
		return
	}
	logger.Info("self-host telemetry disabled (opt-in via MULTICA_SELF_HOST_TELEMETRY=1)")
}
