package directory

import "testing"

func TestExists(t *testing.T) {
	t.Parallel()

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ensure a path that is known to exist actually exists",
			args: args{
				path: "./",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ensure a path that is known not to exist actually exists",
			args: args{
				path: "./asdfasdf",
			},
			want:    false,
			wantErr: false,
		},
		// NOTE: this will only work on popular linux-based systems.  if there comes a point in time where
		//       it needs to be updated for other OS's, it can be addressed at that time.
		{
			name: "ensure a path that is a file returns an error",
			args: args{
				path: "/etc/resolv.conf",
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
