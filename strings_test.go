package strings

import "testing"

func TestX(t *testing.T) {
	t.Run("Simple case", func(t *testing.T) {
		e := "My test"
		a := ToUpperFirst("my test")
		if e != a {
			t.Errorf("Got: %s, want: %s", a, e)
		}
	})
}
