package common

import "testing"

func TestReverseString(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ensure string is reversed properly",
			args: args{
				"abcdefghijklmnopqrstuvwxyz",
			},
			want: "zyxwvutsrqponmlkjihgfedcba",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ReverseString(tt.args.str); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
