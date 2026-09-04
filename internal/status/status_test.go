package status_test

import (
	"strings"
	"testing"

	"github.com/jumpstart-demo/pr-demo-repo/internal/status"
)

func TestReport_Healthy(t *testing.T) {
	got := status.Report("api", true)
	if !strings.Contains(got, "api") || !strings.Contains(got, "OK") {
		t.Errorf("Report(healthy) = %q, want it to contain 'api' and 'OK'", got)
	}
}

func TestReport_Unhealthy(t *testing.T) {
	got := status.Report("api", false)
	if !strings.Contains(got, "DOWN") {
		t.Errorf("Report(unhealthy) = %q, want it to contain 'DOWN'", got)
	}
}
