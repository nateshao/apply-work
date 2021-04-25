package main
import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 2, 4},
	}

	for _, c := range cases {
		result := Add(c.first, c.second)
		if result != c.excepted {
			t.Fatalf("add function failed, first: %d, second:%d, execpted:%d, result:%d", c.first, c.second, c.excepted, result)
		}
	}
}