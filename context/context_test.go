package context

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := NewContext(context.Background(), map[string]interface{}{
		"test": "test value",
	})

	t.Log(FromContext(ctx))
}
