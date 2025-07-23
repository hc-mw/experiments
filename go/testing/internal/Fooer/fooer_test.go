package fooer

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output string
	}{
		{"ip = 9, op = foo", 9, "Foo"},
		{"ip = 3, op = foo", 3, "Foo"},
		{"ip = 1, op = 1", 1, "1"},
		{"ip = 0, op = Foo", 0, "Foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.output {
				t.Errorf("got %s, want %s", ans, tt.output)
			}
		})
	}
}

func TestFooer2(t *testing.T) {
	ip := 3
	res := Fooer(ip)
	t.Logf("input: %d", ip)
	if res != "Foo" {
		t.Errorf("result was incorrect, got %s, want %s", res, "Foo")
	}
	t.Fatalf("stop test")
	t.Error("wont be executed")
}

func TestFooerParallel(t *testing.T) {
	t.Run("test 3 in parallel", func(t *testing.T) {
		t.Parallel()
		res := Fooer(3)
		want := "Foo"
		if res != want {
			t.Errorf("res was incorrect, got %s, want %s", res, want)
		}
	})
	t.Run("test 7 in parallel", func(t *testing.T) {
		t.Parallel()
		res := Fooer(7)
		want := "7"
		if res != want {
			t.Errorf("res was incorrect, got %s, want %s", res, want)
		}
	})
}

func TestFooerSkiped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res := Fooer(3)
	if res != "Foo" {
		t.Errorf("result was incorrect, got %s, want %s", res, "Foo")
	}
}

func Test_With_Cleanup(t *testing.T) {

	t.Cleanup(func() {

	})
}
