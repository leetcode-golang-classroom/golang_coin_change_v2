package sol

import "testing"

func BenchmarkTestDFS(b *testing.B) {
	amount := 5
	coins := []int{1, 2, 5}
	for idx := 0; idx < b.N; idx++ {
		changeDFS(amount, coins)
	}
}
func Test_changeDFS(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "amount = 5, coins = [1,2,5]",
			args: args{amount: 5, coins: []int{1, 2, 5}},
			want: 4,
		},
		{
			name: "amount = 3, coins = [2]",
			args: args{amount: 3, coins: []int{2}},
			want: 0,
		},
		{
			name: "amount = 10, coins = [10]",
			args: args{amount: 10, coins: []int{10}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeDFS(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("changeDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
