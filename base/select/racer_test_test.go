package racer

import(
	"net/http/httptest"
	"testing"
	"time"
)

func Test_makeDelayedServer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		delay time.Duration
		want  *httptest.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeDelayedServer(tt.delay)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("makeDelayedServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
