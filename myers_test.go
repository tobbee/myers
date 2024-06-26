package myers

import (
	"reflect"
	t "testing"
)

type TestCase struct {
	l1  []string
	l2  []string
	exp []Op
}

func TestDiff(t *t.T) {
	A := "A"
	B := "B"
	C := "C"
	testCases := []TestCase{
		{[]string{}, []string{}, []Op{}},
		{[]string{}, []string{"foo"}, []Op{{OpInsert, 0, 0, "foo"}}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "bar", "baz"}, []Op{}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "baz"}, []Op{{OpDelete, 1, -1, "bar"}}},
		{[]string{"baz"}, []string{"foo", "baz"}, []Op{{OpInsert, 0, 0, "foo"}}},
		{[]string{"bar", "baz"}, []string{"foo", "baz"}, []Op{{OpDelete, 0, -1, "bar"}, {OpInsert, 1, 0, "foo"}}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "bar"}, []Op{{OpDelete, 2, -1, "baz"}}},
		{[]string{A, B, C, A, B, B, A}, []string{C, B, A, B, A, C},
			[]Op{{OpDelete, 0, -1, A}, {OpInsert, 1, 0, C}, {OpDelete, 2, -1, C}, {OpDelete, 5, -1, B}, {OpInsert, 7, 5, C}}},
		{[]string{C, A, B, A, B, A, B, A, B, A, B, A, B, C},
			[]string{B, A, B, A, B, A, B, A, B, A, B, A, B, A},
			[]Op{{OpDelete, 0, -1, C}, {OpInsert, 1, 0, B}, {OpDelete, 13, -1, C}, {OpInsert, 14, 13, A}}},
	}
	for _, c := range testCases {
		act := DiffStr(c.l1, c.l2)
		if !reflect.DeepEqual(c.exp, act) {
			t.Errorf("Failed diff, expected %v actual %v\n", c.exp, act)
		}
	}
}
