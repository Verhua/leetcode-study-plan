package day01

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type nums struct {
		nums []int
	}
	tests := []struct {
		name string
		nums nums
		want bool
	}{
		{
			name: "true",
			nums: nums{nums: []int{1, 2, 4, 3, 1}},
			want: true,
		},
		{
			name: "false",
			nums: nums{nums: []int{1, 2, 4, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				res := containsDuplicate(tt.nums.nums)
				if res != tt.want {
					t.Errorf("containsDuplicate() got1 = %v, want %v", res, tt.want)
				}
			},
		)
	}

}
