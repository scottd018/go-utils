package list

import (
	"reflect"
	"testing"
)

func Test_integers_Has(t *testing.T) {
	t.Parallel()

	type args struct {
		integer int
	}

	tests := []struct {
		name string
		list integers
		args args
		want bool
	}{
		{
			name: "empty integer list returns false",
			list: integers{},
			args: args{
				integer: 0,
			},
			want: false,
		},
		{
			name: "integer list returns true if it has an integer",
			list: integers{1, 2},
			args: args{
				integer: 1,
			},
			want: true,
		},
		{
			name: "integer list returns false if it does not have an integer",
			list: integers{1, 2},
			args: args{
				integer: 3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Has(tt.args.integer); got != tt.want {
				t.Errorf("integers.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integers_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list integers
		want []int
	}{
		{
			name: "empty integer list returns empty",
			list: integers{},
			want: []int{},
		},
		{
			name: "non-empty integer list with overlapping returns unique strings",
			list: integers{1, 1, 2, 2},
			want: []int{1, 2},
		},
		{
			name: "non-empty integer list without overlapping returns unique integers",
			list: integers{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.list.Unique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("integers.Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
