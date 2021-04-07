package log

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"testing"
)

var flagTests = []struct {
	env   string
	level zapcore.Level
}{
	{"debug", zapcore.DebugLevel}, {"DEBUG", zapcore.DebugLevel},
	{"info", zapcore.InfoLevel}, {"INFO", zapcore.InfoLevel},
	{"warn", zapcore.WarnLevel}, {"WARN", zapcore.WarnLevel},
	{"error", zapcore.ErrorLevel}, {"ERROR", zapcore.ErrorLevel},
	{"panic", zapcore.PanicLevel}, {"PANIC", zapcore.PanicLevel},
	{"fatal", zapcore.FatalLevel}, {"FATAL", zapcore.FatalLevel},
}

func TestSetupLogger(t *testing.T) {
	for _, tt := range flagTests {
		t.Run(tt.env, func(t *testing.T) {
			_ = os.Setenv("LOG_LEVEL", tt.level.String())

			// when
			SetupLogger()

			// then
			assert.True(t, Logger != nil)
			Logger.Desugar().Core().Enabled(tt.level)

			// reset once to run again
			once = new(sync.Once)
		})
	}
}
