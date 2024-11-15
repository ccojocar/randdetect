package randdetect

import "testing"

func TestIsRandom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "English words are not random",
			input:  "thisisnotrandom",
			output: false,
		},
		{
			name:   "English words are not random with mixed caps",
			input:  "This is not random Number",
			output: false,
		},
		{
			name:   "Short words are not random",
			input:  "zhf",
			output: false,
		},
		{
			name:   "Random alpha numerical world",
			input:  "ydiWASqq",
			output: true,
		},
		{
			name:   "Random alpha numerical world",
			input:  "aowkaoskaos",
			output: true,
		},
		{
			name:   "Repeated chars are random",
			input:  "rrrrrrrrrrrr",
			output: true,
		},
		{
			name:   "Dominated repeated bigrams are random",
			input:  "bltsnotbltsbl",
			output: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, expected := IsRandom(test.input), test.output; got != expected {
				t.Fatalf("IsRandom(%q) returned %t; expected: %t", test.input, got, expected)
			}
		})
	}
}
