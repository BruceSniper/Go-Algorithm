/**
 * @ Author: darkknight
 * @ GitHub: https://github.com/BruceSniper
 * @ Description:
 * @ Date: 2023/7/10 23:29
 * @ File: main_test.go.go
 **/

package main

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
				target: -2,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
