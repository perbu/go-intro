// Code generated by "stringer -type=PingerErrorCode"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[connectError-1]
	_ = x[pingError-2]
}

const _PingerErrorCode_name = "connectErrorpingError"

var _PingerErrorCode_index = [...]uint8{0, 12, 21}

func (i PingerErrorCode) String() string {
	i -= 1
	if i < 0 || i >= PingerErrorCode(len(_PingerErrorCode_index)-1) {
		return "PingerErrorCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PingerErrorCode_name[_PingerErrorCode_index[i]:_PingerErrorCode_index[i+1]]
}