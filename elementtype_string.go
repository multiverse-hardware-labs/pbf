// Code generated by "stringer -type=ElementType"; DO NOT EDIT.

package pbf // import "m4o.io/pbf"

import "fmt"

const _ElementType_name = "NODEWAYRELATION"

var _ElementType_index = [...]uint8{0, 4, 7, 15}

func (i ElementType) String() string {
	if i < 0 || i >= ElementType(len(_ElementType_index)-1) {
		return fmt.Sprintf("ElementType(%d)", i)
	}
	return _ElementType_name[_ElementType_index[i]:_ElementType_index[i+1]]
}
