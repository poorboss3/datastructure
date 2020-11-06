package list

import "testing"

func Test_node_Insert(t *testing.T) {
	var list = NewShiplist()
	list.Insert(1)
	list.Insert(2)
	list.Insert(4)
	list.PrintAll()
}
