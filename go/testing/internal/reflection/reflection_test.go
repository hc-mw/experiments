package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"struct",
			struct {
				Name string
				City string
				Age  int
			}{"Chris", "London", 33},
			[]string{"Chris", "London"},
		},
		{
			"nested struct",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"array",
			[2]Profile{
				{33, "London"},
				{34, "Lisbon"},
			},
			[]string{"London", "Lisbon"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := []string{}
			Walk(tt.Input, func(ip string) {
				got = append(got, ip)
			})

			if !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("got %v, want %v", got, tt.Expected)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(ip string) {
			got = append(got, ip)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{23, "K"}
			ch <- Profile{23, "H"}
			close(ch)
		}()

		got := []string{}

		Walk(ch, func(ip string) {
			got = append(got, ip)
		})

		want := []string{"K", "H"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didnt'", haystack, needle)
	}
}
