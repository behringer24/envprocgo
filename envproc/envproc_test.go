package envproc

import "testing"

func TestEnvproc(t *testing.T) {
	want := "Hello, world."
	if got := Envproc(); got != want {
		t.Errorf("Envproc() = %q, want %q", got, want)
	}
}
