// Code generated by "stringer -type=Name -output=name_string.go"; DO NOT EDIT.

package action

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Quit-1]
	_ = x[ToggleGrid-2]
	_ = x[ToggleLimitFPS-3]
	_ = x[ToggleLimitLifespans-4]
	_ = x[ToggleBoundingBoxes-5]
	_ = x[Pause-6]
	_ = x[ToggleTextures-7]
	_ = x[ToggleCollisions-8]
	_ = x[ToggleCollisionBoxes-9]
	_ = x[DropFood-10]
}

const _Name_name = "QuitToggleGridToggleLimitFPSToggleLimitLifespansToggleBoundingBoxesPauseToggleTexturesToggleCollisionsToggleCollisionBoxesDropFood"

var _Name_index = [...]uint8{0, 4, 14, 28, 48, 67, 72, 86, 102, 122, 130}

func (i Name) String() string {
	i -= 1
	if i >= Name(len(_Name_index)-1) {
		return "Name(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Name_name[_Name_index[i]:_Name_index[i+1]]
}
