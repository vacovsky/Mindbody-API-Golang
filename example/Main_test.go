package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "dadawwa"},
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
