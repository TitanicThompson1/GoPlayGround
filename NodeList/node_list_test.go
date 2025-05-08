package NodeList

import "testing"

func TestNodeListFind(t *testing.T) {
	ll := NodeList[string]{nil, "Hello"};
	ll.Append("World")
	ll.Append("Goodbye")
	ll.Append("World")

	cases := []struct{
		in string
		want int
	}{
		{"World", 1},
		{"world", -1},
		{"Hello", 0},
		{"Helloo", -1},
	}

	for _, c := range cases {
		got := ll.Find(c.in)
		if  got != c.want {
			t.Errorf("NodeListFind(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}