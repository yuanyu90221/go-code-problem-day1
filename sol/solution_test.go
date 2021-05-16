package sol

import "testing"

func TestSearchInsertionPosition(t *testing.T) {
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
			name: "Example1",
			args: args{
				nums:   []int{1, 3, 5, 7, 9},
				target: 6,
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				nums:   []int{1, 3, 3, 3, 3},
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsertionPosition(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsertionPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
