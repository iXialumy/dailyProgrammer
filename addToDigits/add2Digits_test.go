package addToDigits

import "testing"

func Test_add2digits(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{111}, 222},
		{"", args{222}, 333},
		{"", args{333}, 444},
		{"", args{444}, 555},
		{"", args{555}, 666},
		{"", args{998}, 10109},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add2digits(tt.args.i); got != tt.want {
				t.Errorf("add2digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
