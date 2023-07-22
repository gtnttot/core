// Code generated by "stringer -type=Buttons"; DO NOT EDIT.

package mouse

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoButton-0]
	_ = x[Left-1]
	_ = x[Middle-2]
	_ = x[Right-3]
	_ = x[ButtonsN-4]
}

const _Buttons_name = "NoButtonLeftMiddleRightButtonsN"

var _Buttons_index = [...]uint8{0, 8, 12, 18, 23, 31}

func (i Buttons) String() string {
	if i < 0 || i >= Buttons(len(_Buttons_index)-1) {
		return "Buttons(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Buttons_name[_Buttons_index[i]:_Buttons_index[i+1]]
}

func (i *Buttons) FromString(s string) error {
	for j := 0; j < len(_Buttons_index)-1; j++ {
		if s == _Buttons_name[_Buttons_index[j]:_Buttons_index[j+1]] {
			*i = Buttons(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Buttons")
}

var _Buttons_descMap = map[Buttons]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
}

func (i Buttons) Desc() string {
	if str, ok := _Buttons_descMap[i]; ok {
		return str
	}
	return "Buttons(" + strconv.FormatInt(int64(i), 10) + ")"
}
