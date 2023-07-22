// Code generated by "stringer -type=MethViewFlags"; DO NOT EDIT.

package giv

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MethViewConfirm-0]
	_ = x[MethViewShowReturn-1]
	_ = x[MethViewNoUpdateAfter-2]
	_ = x[MethViewHasSubMenu-3]
	_ = x[MethViewHasSubMenuVal-4]
	_ = x[MethViewKeyFun-5]
	_ = x[MethViewFlagsN-6]
}

const _MethViewFlags_name = "MethViewConfirmMethViewShowReturnMethViewNoUpdateAfterMethViewHasSubMenuMethViewHasSubMenuValMethViewKeyFunMethViewFlagsN"

var _MethViewFlags_index = [...]uint8{0, 15, 33, 54, 72, 93, 107, 121}

func (i MethViewFlags) String() string {
	if i < 0 || i >= MethViewFlags(len(_MethViewFlags_index)-1) {
		return "MethViewFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MethViewFlags_name[_MethViewFlags_index[i]:_MethViewFlags_index[i+1]]
}

func (i *MethViewFlags) FromString(s string) error {
	for j := 0; j < len(_MethViewFlags_index)-1; j++ {
		if s == _MethViewFlags_name[_MethViewFlags_index[j]:_MethViewFlags_index[j+1]] {
			*i = MethViewFlags(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: MethViewFlags")
}

var _MethViewFlags_descMap = map[MethViewFlags]string{
	0: `MethViewConfirm confirms action before proceeding`,
	1: `MethViewShowReturn shows the return value from the method`,
	2: `MethViewNoUpdateAfter means do not update window after method runs (default is to do so)`,
	3: `MethViewHasSubMenu means that this action has a submenu option -- argument values will be selected from the auto-generated submenu`,
	4: `MethViewHasSubMenuVal means that this action was called using a submenu and the SubMenuVal has the selected value`,
	5: `MethViewKeyFun means this action's only function is to emit the key fun`,
	6: ``,
}

func (i MethViewFlags) Desc() string {
	if str, ok := _MethViewFlags_descMap[i]; ok {
		return str
	}
	return "MethViewFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
