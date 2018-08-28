package main

import "testing"

/**
    @author good mood
    2018/7/15 下午8:45
*/

func TestTriangle(t *testing.T)  {
	tests := []struct{a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTaiangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); " +
				"got %d; expected %d",
					tt.a, tt.b, actual, tt.c)
		}
	}
}