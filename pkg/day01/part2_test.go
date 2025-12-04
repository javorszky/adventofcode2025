package day01

import "testing"

func Test_coversZero(t *testing.T) {
	type args struct {
		current  int
		rotation int
	}
	tests := []struct {
		name       string
		args       args
		covers     int
		newCurrent int
	}{
		{
			name: "50 - 4",
			args: args{
				current:  50,
				rotation: -4,
			},
			covers:     0,
			newCurrent: 46,
		},
		{
			name: "50 - 50",
			args: args{
				current:  50,
				rotation: -50,
			},
			covers:     1,
			newCurrent: 0,
		},
		{
			name: "50 - 51",
			args: args{
				current:  50,
				rotation: -51,
			},
			covers:     1,
			newCurrent: -1,
		},
		{
			name: "50 + 50",
			args: args{
				current:  50,
				rotation: 50,
			},
			covers:     1,
			newCurrent: 100,
		},
		{
			name: "50 + 851",
			args: args{
				current:  50,
				rotation: 851,
			},
			covers:     9,
			newCurrent: 901,
		},
		{
			name: "-50 + 250",
			args: args{
				current:  -50,
				rotation: 250,
			},
			covers:     3,
			newCurrent: 200,
		},
		{
			name: "0 + 0",
			args: args{
				current:  0,
				rotation: 0,
			},
			covers:     0,
			newCurrent: 0,
		},
		{
			name: "-99 + 201",
			args: args{
				current:  -99,
				rotation: 201,
			},
			covers:     2,
			newCurrent: 102,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := coversZero(tt.args.current, tt.args.rotation)
			if got != tt.covers {
				t.Errorf("coversZero() got = %v, want %v", got, tt.covers)
			}
			if got1 != tt.newCurrent {
				t.Errorf("coversZero() got1 = %v, want %v", got1, tt.newCurrent)
			}
		})
	}
}

func Test_calculateCoversZero(t *testing.T) {
	type args struct {
		rotations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "back and forth",
			args: args{
				rotations: []int{-55, 57, -60, 8, -102, 130, -300},
				//                -5, 52,  -8, 0, -102,  28, -272
				//                 1,  1,   1, 1,    2,   2,    3
			},
			want: 11,
		},
		{
			name: "lands on",
			args: args{
				rotations: []int{-55, 105},
			},
			want: 3,
		},
		{
			name: "zig zag, but positive",
			args: args{
				rotations: []int{-45, 95, -45, 95, -45, 95, -45, 95, -45},
				//                10, 105,  60, 155, 110, 205, 160, 255, 210
				//                 0,   1,   1,   1,   0,   1,   1,   1,   0
			},
			want: 6,
		},
		{
			name: "zig zag, but never touches a 0",
			args: args{
				rotations: []int{-25, 50, -50, 50, -50, 50, -50, 50, -50, 50},
			},
			want: 0,
		},
		{
			name: "one big spin",
			args: args{
				rotations: []int{2500},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCoversZero(tt.args.rotations); got != tt.want {
				t.Errorf("calculateCoversZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
