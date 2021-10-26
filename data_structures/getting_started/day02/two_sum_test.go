package day02

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{nums: []int{2, 7, 11, 15}, target: 18},
			want: []int{1, 2},
		},
		{
			name: "test02",
			args: args{nums: []int{3, 2, 4}, target: 5},
			want: []int{0, 1},
		},
		{
			name: "test03",
			args: args{nums: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
		{
			name: "test04",
			args: args{nums: []int{3, 2, 3}, target: 6},
			want: []int{0, 2},
		},
		{
			name: "test05",
			args: args{nums: []int{}, target: 2},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
