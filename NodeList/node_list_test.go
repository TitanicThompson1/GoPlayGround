package nodelist

import "testing"

func TestNodeList(t *testing.T) {
	llstr := NodeList[string]{nil, "Hello"}
	llstr.Append("World")
	llstr.Append("Goodbye")
	llstr.Append("World")

	llempty := NodeList[string]{nil, ""}
	
	llnumber := NodeList[int]{nil, 0}
	llnumber.Append(1)
	llnumber.Append(7)
	llnumber.Append(2)
	llnumber.Append(0)

	t.Run("Find", func(t *testing.T) {
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
			got := llstr.Find(c.in)
			if  got != c.want {
				t.Errorf("NodeListFind(%s) == %d, want %d", c.in, got, c.want)
			}
		}
		
		casesInt := []struct{
			in int
			want int
		}{
			{0, 0},
			{7, 2},
			{2, 3},
		}

		for _, c := range casesInt {
			got := llnumber.Find(c.in)
			if  got != c.want {
				t.Errorf("NodeListFind(%d) == %d, want %d", c.in, got, c.want)
			}
		}


		got := llempty.Find("Hey")

		if  got != -1 {
			t.Errorf("NodeListFind(%s) == %d, want %d", "Hey", got, -1)
		}
	})

	t.Run("Append", func(t *testing.T) {
		llappend := NodeList[string]{nil, ""}

		llappend.Append("Hello")

		if llappend.Next == nil || llappend.Next.Val != "Hello"{
			t.Errorf("NodeListAppend: Trying to append Hello, but Next was %v", llappend.Next)
		}

		llappend.Append("Goodbye")

		if llappend.Next == nil || llappend.Next.Val != "Hello"{
			t.Errorf("NodeListAppend: Trying to append Goodbye, but first Next was %v", llappend.Next)
		}

		if llappend.Next.Next == nil || llappend.Next.Next.Val != "Goodbye"{
			t.Errorf("NodeListAppend: Trying to append Goodbye, but second Next was %v", llappend.Next)
		}
	})
}
