package golang

// Container -
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

// List -
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Sort(comparator Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})

	Container
}

// Set -
type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool

	Container
}

// Stack -
type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	Container
}

// Map -
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}

	Container
}

// Tree -
type Tree interface {
	Container
}

// Trie -
type Trie interface {
	Put(route string, value interface{}) error
	Get(route string) (interface{}, bool)
	Remove(route string)
	Keys() []string

	Container
}

// Comparator -
type Comparator func(a, b interface{}) int
