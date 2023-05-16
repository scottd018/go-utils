package list

import (
	"reflect"
	"testing"
)

func Test_strings_Has(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
	}

	tests := []struct {
		name string
		list strings
		args args
		want bool
	}{
		{
			name: "empty string list returns false",
			list: strings{},
			args: args{
				str: "something",
			},
			want: false,
		},
		{
			name: "string list returns true if it has a string",
			list: strings{"one", "two"},
			args: args{
				str: "one",
			},
			want: true,
		},
		{
			name: "string list returns false if it does not have a string",
			list: strings{"one", "two"},
			args: args{
				str: "three",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Has(tt.args.str); got != tt.want {
				t.Errorf("strings.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strings_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list strings
		want []string
	}{
		{
			name: "empty string list returns empty",
			list: strings{},
			want: []string{},
		},
		{
			name: "non-empty string list with overlapping returns unique strings",
			list: strings{"one", "one", "two", "two"},
			want: []string{"one", "two"},
		},
		{
			name: "non-empty string list without overlapping returns unique strings",
			list: strings{"one", "two", "three", "four"},
			want: []string{"one", "two", "three", "four"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Unique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strings.Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
