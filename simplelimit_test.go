package resourcelimit

import (
	"testing"
)

func Test_limit(t *testing.T) {
	limit := NewLimit(5)
	for i := 0; i < 5; i++ {
		limit.UseResource()
	}
	for i := 0; i < 5; i++ {
		limit.FreeResource()
	}
	return
}
