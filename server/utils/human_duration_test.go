package utils

import (
	"fmt"
	"path"
	"strings"
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "5h20m",
			args:    args{"5h20m"},
			want:    time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d5h20m",
			args:    args{"1d5h20m"},
			want:    24*time.Hour + time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d",
			args:    args{"1d"},
			want:    24 * time.Hour,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	f := "/dir/file.txt"
	fmt.Println(f)
	ext := path.Ext(f)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(f, ext))
	fmt.Println(path.Split(f))
	fmt.Println(path.Dir(f))
	fmt.Println(path.Base(f))
	fmt.Println(path.Clean(f))
	fmt.Println(path.IsAbs(f))
	fmt.Println(path.Join(f, "test"))
	fmt.Println(path.Match(f, "test"))

}
