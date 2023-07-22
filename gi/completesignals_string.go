// Code generated by "stringer -type=CompleteSignals"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompleteSelect-0]
	_ = x[CompleteExtend-1]
}

const _CompleteSignals_name = "CompleteSelectCompleteExtend"

var _CompleteSignals_index = [...]uint8{0, 14, 28}

func (i CompleteSignals) String() string {
	if i < 0 || i >= CompleteSignals(len(_CompleteSignals_index)-1) {
		return "CompleteSignals(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompleteSignals_name[_CompleteSignals_index[i]:_CompleteSignals_index[i+1]]
}

func (i *CompleteSignals) FromString(s string) error {
	for j := 0; j < len(_CompleteSignals_index)-1; j++ {
		if s == _CompleteSignals_name[_CompleteSignals_index[j]:_CompleteSignals_index[j+1]] {
			*i = CompleteSignals(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: CompleteSignals")
}

var _CompleteSignals_descMap = map[CompleteSignals]string{
	0: `CompleteSelect means the user chose one of the possible completions`,
	1: `CompleteExtend means user has requested that the seed extend if all completions have a common prefix longer than current seed`,
}

func (i CompleteSignals) Desc() string {
	if str, ok := _CompleteSignals_descMap[i]; ok {
		return str
	}
	return "CompleteSignals(" + strconv.FormatInt(int64(i), 10) + ")"
}
