package tests

import (
	"context"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestingT interface {
	require.TestingT
	Helper()
	Cleanup(func())
}

func getContext(tb TestingT) (ctx context.Context) {
	ctx = context.Background()
	var cancel func()

	t, isTest := tb.(interface {
		Deadline() (deadline time.Time, ok bool)
	})
	if isTest {
		d, hasDeadline := t.Deadline()
		if hasDeadline {
			ctx, cancel = context.WithDeadline(ctx, d)
			tb.Cleanup(cancel)
			return
		}
	}

	ctx, cancel = context.WithCancel(ctx)
	tb.Cleanup(cancel)
	return
}

// DefaultWaitTimeout is the default wait timeout. If you have a *testing.T, use WaitTimeout instead.
const DefaultWaitTimeout = 30 * time.Second

// WaitTimeout returns a timeout based on the test's Deadline, if available.
// Especially important to use in parallel tests, as their individual execution
// can get paused for arbitrary amounts of time.
func WaitTimeout(t *testing.T) time.Duration {
	if d, ok := t.Deadline(); ok {
		// 10% buffer for cleanup and scheduling delay
		return time.Until(d) * 9 / 10
	}
	return DefaultWaitTimeout
}

// TestInterval is just a sensible poll interval that gives fast tests without
// risk of spamming
const TestInterval = 100 * time.Millisecond

// AssertEventually waits for f to return true
func AssertEventually(t *testing.T, f func() bool) {
	assert.Eventually(t, f, WaitTimeout(t), TestInterval/2)
}

// RequireSignal waits for the channel to close (or receive anything) and
// fatals the test if the default wait timeout is exceeded
func RequireSignal(t *testing.T, ch <-chan struct{}, failMsg string) {
	select {
	case <-ch:
	case <-time.After(WaitTimeout(t)):
		t.Fatal(failMsg)
	}
}

// SkipShort skips tb during -short runs, and notes why.
func SkipShort(tb testing.TB, why string) {
	if testing.Short() {
		tb.Skipf("skipping: %s", why)
	}
}

const envVarRunFlakey = "CL_RUN_FLAKEY"

var runFlakey = sync.OnceValues(func() (bool, error) {
	s := os.Getenv(envVarRunFlakey)
	if s == "" {
		return false, nil
	}
	return strconv.ParseBool(s)
})

func SkipFlakey(t *testing.T, ticketURL string) {
	if ok, err := runFlakey(); !ok {
		if err != nil {
			t.Logf("Failed to parse %s: %v", envVarRunFlakey, err)
		}
		t.Skip("Flakey", ticketURL)
	}
}
