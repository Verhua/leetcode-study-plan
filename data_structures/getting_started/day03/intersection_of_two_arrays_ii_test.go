package day03

import (
	"reflect"
	"testing"
)

// TestIntersect 测试交集
func TestIntersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "test02",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
		{
			name: "test03",
			args: args{
				nums1: []int{4},
				nums2: []int{5},
			},
			want: []int{},
		},
		{
			name: "test04",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2},
			},
			want: []int{2},
		},
		{
			name: "test05",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: []int{},
		},
		{
			name: "test06",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
