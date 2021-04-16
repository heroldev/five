package five

import (
	"testing"
)

func TestFiveMath(test *testing.T) {

	test.Run("five addition", func(t *testing.T) {

		got := Five() + Five()

		want := 10

		assertInt(t, got, want)

	})

	test.Run("five subtraction", func(t *testing.T) {

		got := Five() - Five()
		want := 0

		assertInt(t, got, want)

	})

	test.Run("five multiplication", func(t *testing.T) {

		got := Five() * Five()
		want := 25

		assertInt(t, got, want)

	})

	test.Run("five division", func(t *testing.T) {

		got := Five() / Five()
		want := 1

		assertInt(t, got, want)

	})

	test.Run("five factorial", func(t *testing.T) {

		five := GimmeFive()
		got := five.Factorial()
		want := 120

		assertInt(t, got, want)

	})

}

func TestAssert(test *testing.T) {
	test.Run("test assert true", func(t *testing.T) {

		five := GimmeFive()
		got := five.isFive(Five())
		want := true

		if got != want {
			t.Errorf("five assert function broken")
		}

	})

	test.Run("test assert false", func(t *testing.T) {

		five := GimmeFive()
		got := five.isFive(10)
		want := false

		if got != want {
			t.Errorf("five assert function very broken")
		}

	})
}

func assertInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}
}
