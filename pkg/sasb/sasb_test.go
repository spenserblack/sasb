package sasb

import "testing"

func TestSasb(t *testing.T) {
	tests := []struct {
		input       string
		binary      []string
		octal       []string
		decimal     []string
		hexadecimal []string
	}{
		{
			input:       "@",
			binary:      []string{"01000000"},
			octal:       []string{"100"},
			decimal:     []string{"64"},
			hexadecimal: []string{"40"},
		},
		{
			input:       "ガ",
			binary:      []string{"11100011", "10000010", "10101100"},
			octal:       []string{"343", "202", "254"},
			decimal:     []string{"227", "130", "172"},
			hexadecimal: []string{"e3", "82", "ac"},
		},
	}

	for _, tt := range tests {
		binary, octal, decimal, hexadecimal := StringAsBytes(tt.input)

		checks := []struct {
			name   string
			actual []string
			want   []string
		}{
			{"binary", binary, tt.binary},
			{"octal", octal, tt.octal},
			{"decimal", decimal, tt.decimal},
			{"hexadecimal", hexadecimal, tt.hexadecimal},
		}

		for _, ttt := range checks {
			if len(ttt.actual) != len(ttt.want) {
				t.Fatalf(`%s: %v and %v have different lengths`, tt.input, ttt.actual, ttt.want)
			}
			for i, v := range ttt.actual {
				wantv := ttt.want[i]
				if v != wantv {
					t.Fatalf(`%s of %q = %q, want %q`, ttt.name, tt.input, ttt.actual, ttt.want)
				}
			}
		}
	}
}
