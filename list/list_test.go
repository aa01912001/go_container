package list

import (
	"reflect"
	"testing"
)

func TestPushFront(t *testing.T) {
	type testCase[T any] struct {
		set       []T
		wantFront T
		wantBack  T
	}

	testCases := []testCase[interface{}]{
		{
			set:       []interface{}{3, 1, 6, 5, 1},
			wantFront: 1,
			wantBack:  3,
		},
		{
			set:       []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			wantFront: 1.1,
			wantBack:  3.1,
		},
		{
			set:       []interface{}{"e", "d", "c", "b", "a"},
			wantFront: "a",
			wantBack:  "e",
		},
	}

	for _, tc := range testCases {
		l := Init[interface{}]()
		for _, v := range tc.set {
			l.PushFront(v)
		}
		if l.Front().Value != tc.wantFront {
			t.Errorf("wantFront =  %v, but result = %v", tc.wantFront, l.Front().Value)
		}
		if l.Back().Value != tc.wantBack {
			t.Errorf("wantBack =  %v, but result = %v", tc.wantBack, l.Back().Value)
		}
	}
}

func TestPushBack(t *testing.T) {
	type testCase[T any] struct {
		set       []T
		wantFront T
		wantBack  T
	}

	testCases := []testCase[interface{}]{
		{
			set:       []interface{}{3, 1, 6, 5, 1},
			wantFront: 3,
			wantBack:  1,
		},
		{
			set:       []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			wantFront: 3.1,
			wantBack:  1.1,
		},
		{
			set:       []interface{}{"e", "d", "c", "b", "a"},
			wantFront: "e",
			wantBack:  "a",
		},
	}

	for _, tc := range testCases {
		l := Init[interface{}]()
		for _, v := range tc.set {
			l.PushBack(v)
		}
		if l.Front().Value != tc.wantFront {
			t.Errorf("wantFront =  %v, but result = %v", tc.wantFront, l.Front().Value)
		}
		if l.Back().Value != tc.wantBack {
			t.Errorf("wantBack =  %v, but result = %v", tc.wantBack, l.Back().Value)
		}
	}
}

func TestInsertBefore(t *testing.T) {
	type testCase[T any] struct {
		set  []T
		want []T
	}

	testCases := []testCase[interface{}]{
		{
			set:  []interface{}{3, 1, 6, 5, 1},
			want: []interface{}{1, 5, 6, 1, 3},
		},
		{
			set:  []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			want: []interface{}{1.1, 5.7, 6.1, 1.2, 3.1},
		},
		{
			set:  []interface{}{"e", "d", "c", "b", "a"},
			want: []interface{}{"a", "b", "c", "d", "e"},
		},
	}

	for _, tc := range testCases {
		l := Init[interface{}]()

		n := l.PushFront(tc.set[0])
		for i := 1; i < len(tc.set); i++ {
			n = l.InsertBefore(tc.set[i], n)
		}

		var result []interface{}
		for node := l.Front(); node != nil; node = node.Next() {
			result = append(result, node.Value)
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, result)
		}

	}
}

func TestInsertAfter(t *testing.T) {
	type testCase[T any] struct {
		set  []T
		want []T
	}

	testCases := []testCase[interface{}]{
		{
			set:  []interface{}{3, 1, 6, 5, 1},
			want: []interface{}{3, 1, 6, 5, 1},
		},
		{
			set:  []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			want: []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
		},
		{
			set:  []interface{}{"e", "d", "c", "b", "a"},
			want: []interface{}{"e", "d", "c", "b", "a"},
		},
	}

	for _, tc := range testCases {
		l := Init[interface{}]()

		n := l.PushFront(tc.set[0])
		for i := 1; i < len(tc.set); i++ {
			n = l.insertAfter(tc.set[i], n)
		}

		var result []interface{}
		for node := l.Front(); node != nil; node = node.Next() {
			result = append(result, node.Value)
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, result)
		}

	}
}

func TestRemove(t *testing.T) {
	type testCase[T any] struct {
		set  []T
		want []T
	}

	testCases := []testCase[interface{}]{
		{
			set:  []interface{}{3, 1, 6, 5, 1},
			want: []interface{}{3, 1, 5, 1},
		},
		{
			set:  []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			want: []interface{}{3.1, 1.2, 5.7, 1.1},
		},
		{
			set:  []interface{}{"e", "d", "c", "b", "a"},
			want: []interface{}{"e", "d", "b", "a"},
		},
	}

	for _, tc := range testCases {
		l := Init[interface{}]()

		var node *Node[interface{}]
		for i := 0; i < len(tc.set); i++ {
			if i == 2 {
				node = l.PushBack(tc.set[i])
			} else {
				l.PushBack(tc.set[i])
			}
		}

		l.Remove(node)

		var result []interface{}
		for node := l.Front(); node != nil; node = node.Next() {
			result = append(result, node.Value)
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, result)
		}

	}
}
