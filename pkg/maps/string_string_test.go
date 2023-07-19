package maps

import (
	"testing"
)

func Test_strings_HasKey(t *testing.T) {
	t.Parallel()

	type args struct {
		key string
	}

	tests := []struct {
		name string
		m    strings
		args args
		want bool
	}{
		{
			name: "nil map returns false",
			m:    nil,
			args: args{
				key: "one",
			},
			want: false,
		},
		{
			name: "empty map returns false",
			m:    strings{},
			args: args{
				key: "one",
			},
			want: false,
		},
		{
			name: "map without key returns false",
			m:    strings{"one": "two", "three": "four"},
			args: args{
				key: "two",
			},
			want: false,
		},
		{
			name: "map with key returns true",
			m:    strings{"one": "two", "three": "four"},
			args: args{
				key: "three",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.m.HasKey(tt.args.key); got != tt.want {
				t.Errorf("strings.HasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strings_HasValue(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	tests := []struct {
		name string
		m    strings
		args args
		want bool
	}{
		{
			name: "nil map returns false",
			m:    nil,
			args: args{
				value: "one",
			},
			want: false,
		},
		{
			name: "empty map returns false",
			m:    strings{},
			args: args{
				value: "one",
			},
			want: false,
		},
		{
			name: "map without value returns false",
			m:    strings{"one": "two", "three": "four"},
			args: args{
				value: "three",
			},
			want: false,
		},
		{
			name: "map with value returns true",
			m:    strings{"one": "two", "three": "four"},
			args: args{
				value: "four",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.m.HasValue(tt.args.value); got != tt.want {
				t.Errorf("strings.HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
