package link

import (
	"fmt"
	"testing"
)

func TestDLinkInsert(t *testing.T) {
	d := NewDLinkList()

	d.Append(1)
	d.Append(2)
	d.Append(3)
	d.Append(4)
	d.Prepend(9)

	d.Delete(1)

	fmt.Println(d.length)
	d.Traverse()
}
