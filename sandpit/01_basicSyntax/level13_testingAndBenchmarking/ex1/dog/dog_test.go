package dog

import "testing"

func TestHumanYearsInDogYears(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1 human year eq 7 dog years", args: args{1}, want: 7},
		{name: "2 human year eq 14 dog years", args: args{2}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanYearsInDogYears(tt.args.n); got != tt.want {
				t.Errorf("HumanYearsInDogYears() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanYearsInDogYearsTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1 human year eq 7 dog years", args: args{1}, want: 7},
		{name: "2 human year eq 14 dog years", args: args{2}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanYearsInDogYearsTwo(tt.args.n); got != tt.want {
				t.Errorf("HumanYearsInDogYearsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHumanYearsInDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HumanYearsInDogYears(3)
	}
}

func BenchmarkHumanYearsInDogYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HumanYearsInDogYearsTwo(3)
	}
}
