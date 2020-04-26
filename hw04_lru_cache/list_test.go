package hw04_lru_cache //nolint:golint,stylecheck

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("len", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)
	})

	t.Run("pushing only front", func(t *testing.T) {
		l := NewList()
		l.PushFront("a")
		l.PushFront("b")
		l.PushFront("c")

		expectL := make([]string, 0, 3)
		for i := l.Front(); i != nil; i = i.Prev {
			expectL = append(expectL, i.Value.(string))
		}
		require.Equal(t, "cba", strings.Join(expectL, ""))

		expectR := make([]string, 0, 3)
		for i := l.Back(); i != nil; i = i.Next {
			expectR = append(expectR, i.Value.(string))
		}
		require.Equal(t, "abc", strings.Join(expectR, ""))
	})

	t.Run("pushing only back", func(t *testing.T) {
		l := NewList()
		l.PushBack("x")
		l.PushBack("y")
		l.PushBack("z")

		expectL := make([]string, 0, 3)
		for i := l.Front(); i != nil; i = i.Prev {
			expectL = append(expectL, i.Value.(string))
		}
		require.Equal(t, "xyz", strings.Join(expectL, ""))

		expectR := make([]string, 0, 3)
		for i := l.Back(); i != nil; i = i.Next {
			expectR = append(expectR, i.Value.(string))
		}
		require.Equal(t, "zyx", strings.Join(expectR, ""))
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)  // [10]
		l.PushFront(20) // [20, 10]
		l.PushBack(30)  // [20, 10, 30]

		middle := l.Back().Next // 10
		l.Remove(middle)        // [20, 30]
		require.Equal(t, l.Len(), 2)
		require.ElementsMatch(t, []interface{}{20, 30}, l.Values())
	})

	t.Run("remove all", func(t *testing.T) {
		l := NewList()
		for _, i := range [3]int{10, 20, 30} {
			l.PushFront(i)
		}
		require.Equal(t, 3, l.Len())

		for li := l.Back(); li != nil; li = li.Next {
			//fmt.Printf("has %+v\n", l)
			//for lit := l.Back(); lit != nil; lit = lit.Next {
			//	fmt.Printf("item %v %+v\n", unsafe.Pointer(lit), lit)
			//}
			//fmt.Printf("removing %v\n", unsafe.Pointer(li))
			l.Remove(li)
		}
		//fmt.Printf("result %+v\n", l)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]

		//fmt.Printf("has %+v\n", l)
		//for lit := l.Back(); lit != nil; lit = lit.Next {
		//	fmt.Printf("item %v %+v\n", unsafe.Pointer(lit), lit)
		//}

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		//fmt.Printf("has %+v\n", l)
		//for lit := l.Back(); lit != nil; lit = lit.Next {
		//	fmt.Printf("item %v %+v\n", unsafe.Pointer(lit), lit)
		//}

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})

	t.Run("remove no listed item", func(t *testing.T) {
		li := ListItem{"who am i", nil, nil}
		require.Panics(t, func() { l := NewList(); l.Remove(&li) })
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()
		lis := make([]*ListItem, 0, 2)
		lis = append(lis, l.PushFront(10)) // [10]
		lis = append(lis, l.PushFront(20)) // [20, 10]
		l.MoveToFront(lis[0])              // [10, 20]
		l.MoveToFront(lis[1])              // [20, 10]
		l.MoveToFront(lis[0])              // [10, 20]
		require.Equal(t, 2, l.Len())
		for i, k := l.Front(), 0; i != nil; i, k = i.Prev, k+1 {
			assert.Equal(t, lis[k], i)
		}

	})
}
