// Code generated by "stringer -type=Modifiers"; DO NOT EDIT.

package key

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Shift-0]
	_ = x[Control-1]
	_ = x[Alt-2]
	_ = x[Meta-3]
	_ = x[ModifiersN-4]
}

const _Modifiers_name = "ShiftControlAltMetaModifiersN"

var _Modifiers_index = [...]uint8{0, 5, 12, 15, 19, 29}

func (i Modifiers) String() string {
	if i < 0 || i >= Modifiers(len(_Modifiers_index)-1) {
		return "Modifiers(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Modifiers_name[_Modifiers_index[i]:_Modifiers_index[i+1]]
}

func (i *Modifiers) FromString(s string) error {
	for j := 0; j < len(_Modifiers_index)-1; j++ {
		if s == _Modifiers_name[_Modifiers_index[j]:_Modifiers_index[j+1]] {
			*i = Modifiers(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Modifiers")
}

var _Modifiers_descMap = map[Modifiers]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
}

func (i Modifiers) Desc() string {
	if str, ok := _Modifiers_descMap[i]; ok {
		return str
	}
	return "Modifiers(" + strconv.FormatInt(int64(i), 10) + ")"
}
