package servant

import (
	"context"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if dl, ok := ctx.Deadline(); ok {
		timeout := time.Until(dl)
		t.Logf("timeout: %d", timeout/time.Millisecond)
	}
}
