package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "testing hello",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "testing question",
			args: args{
				s: "who am i?",
			},
			want: "?i ma ohw",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			/*
				if got := ReverseString(tc.args.s) ; got != tc.want {
					t.Errorf("ReverseString() = %v, want %v", got, tc.want)
				}
			*/
			got := ReverseString(tc.args.s)
			if got != tc.want {
				t.Errorf("ReverseString() = %v, want %v", got, tc.want)
			}

		})
	}

	/*
		if ReverseString("") != "1" {
			t.Fail()
		}

		if ReverseString("hello") != "lleh" {
			t.Fail()
		}
	*/
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString("masum")
	}
}
