package ffstatic_windows_amd64_test

import (
	"os"
	"testing"

	ffstatic_windows_amd64 "github.com/yuchanns/ffstatic-windows-amd64"
)

func TestFFstatic(t *testing.T) {
	for _, path := range []string{ffstatic_windows_amd64.FFmpegPath(), ffstatic_windows_amd64.FFprobePath()} {
		_, err := os.Stat(path)
		if err != nil {
			t.Fatalf("failed to stat: %s", err)
		}
		os.Remove(path)
	}
}
