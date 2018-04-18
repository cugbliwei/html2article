package html2article

import (
	"testing"
	"time"
)

func TestCompress(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test0",
			args: args{
				str: "test ",
			},
			want: "test ",
		},
		{
			name: "test1",
			args: args{
				str: " test ",
			},
			want: " test ",
		},
		{
			name: "test2",
			args: args{
				str: "test 2  ",
			},
			want: "test 2 ",
		},
		{
			name: "test3",
			args: args{
				str: "test 3  \n    ",
			},
			want: "test 3 ",
		},
		{
			name: "test4",
			args: args{
				str: "test4",
			},
			want: "test4",
		},
		{
			name: "test5",
			args: args{
				str: "test5  test5  \n test   ",
			},
			want: "test5 test5 test ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress(tt.args.str); got != tt.want {
				t.Errorf("Compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{"1", "abc", "ab", 1},
		{"2", "abc", "abd", 1},
		{"3", "ab", "abcef", 3},
		{"3", "abcdfdsaffsdfsfsfdsfsfsfsfsfs3saafsfsafsfsafsfsfsafsfdsf", "abcdfdsaffsdfsfsfsfsfsf23232safsfsaf2fsafsfsfsafsfdsf11", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.a, tt.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	tests := []struct {
		name string
		a    string
		want int64
	}{
		{"1", "fdaf5小时前 ggagg", time.Now().Add(-5*time.Hour).Unix() / 3600 * 3600},
		{"2", "hgha3天前fdsa", (time.Now().Add(-3 * time.Hour * 24).Unix()) / int64(24*3600) * int64(24*3600)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTime(tt.a); got != tt.want {
				t.Errorf("Time = %v, want %v", got, tt.want)
			}
		})
	}
}
