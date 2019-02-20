package checkDigitsUPC

import "testing"

func Test_checkDigitsUPC(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{4210000526}, 4},
		{"", args{3600029145}, 2},
		{"", args{12345678910}, 4},
		{"", args{1234567}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDigitsUPC(tt.args.input); got != tt.want {
				t.Errorf("checkDigitsUPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
