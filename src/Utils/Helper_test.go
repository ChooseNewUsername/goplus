package Utils

import (
	"strconv"
	"testing"
)

func TestJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"t1",args{[]string{"a","b"}},"ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.strs...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkJoin(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		Join(strconv.Itoa(i))

	}
}
func BenchmarkJoin1(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		Join1(strconv.Itoa(i))

	}
}
func BenchmarkJoin2(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		Join1(strconv.Itoa(i))

	}
}