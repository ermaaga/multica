package selfhosttelemetry

import "testing"

func TestConfigFromOptIn(t *testing.T) {
	cases := map[string]bool{
		"":       false,
		"0":      false,
		"false":  false,
		"yes":    false,
		"1":      true,
		" TRUE ": true,
		"true":   true,
	}
	for raw, want := range cases {
		if got := ConfigFromOptIn(raw).Enabled; got != want {
			t.Errorf("ConfigFromOptIn(%q).Enabled = %v, want %v", raw, got, want)
		}
	}
}
