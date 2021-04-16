package five

import "time"

type takeFive struct {
	five int
}

func Five() int {
	return 5
}

//take
func GimmeFive() takeFive {
	return takeFive{Five()}
}

//The Law of Fives
func (f *takeFive) Law() string {
	return "The Law of Fives states simply that: All things happen in fives, or are divisible by or are multiples of five, or are somehow directly or indirectly appropriate to 5. The Law of Fives is never wrong."
}

//combinatorics
func (f *takeFive) Factorial() int {
	return 120
}

//special bits
func (f *takeFive) UpHigh() string {
	return "⁵"
}

func (f *takeFive) DownLow() string {
	return "₅"
}

func (f *takeFive) TooSlow() int {
	time.Sleep(500 * time.Millisecond)
	return 5
}

func (f *takeFive) Roman() string {
	return "V"
}

func (f *takeFive) Negative() int {
	return -5
}

func (f *takeFive) Loud() string {
	return "FIVE"
}

func (f *takeFive) Smooth() string {
	return "S"
}

func (f *takeFive) MD5() string {
	return ""
}

func (f *takeFive) Golden() float64 {
	return 1.618033988749895
}

//radices
func (f *takeFive) Binary() int {
	return 101
}

func (f *takeFive) Octal() int {
	return 5
}

func (f *takeFive) Hex() int {
	return 5
}

//TODO: Five.Base(int)

//assert
func (f *takeFive) isFive(test int) bool {
	return Five() == test
}
