// Code generated by "stringer -type=Name -output=asset_string.go"; DO NOT EDIT.

package asset

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ant-1]
}

const _Name_name = "Ant"

var _Name_index = [...]uint8{0, 3}

func (i Name) String() string {
	i -= 1
	if i < 0 || i >= Name(len(_Name_index)-1) {
		return "Name(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Name_name[_Name_index[i]:_Name_index[i+1]]
}