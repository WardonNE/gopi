package list

type element[V any] struct {
	prev  *element[V]
	next  *element[V]
	list  *LinkedList[V]
	Value V
}
