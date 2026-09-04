package status

import "fmt"

// Report returns a short human-readable service status message.
func Report(service string, healthy bool) string {
	if healthy {
		return fmt.Sprintf("%s: OK", service)
	}
	return fmt.Sprintf("%s: DOWN", service)
}
