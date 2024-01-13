// Code generated by "goki generate"; DO NOT EDIT.

package lsp

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _CompletionKindValues = []CompletionKind{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

// CompletionKindN is the highest valid value
// for type CompletionKind, plus one.
const CompletionKindN CompletionKind = 26

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _CompletionKindNoOp() {
	var x [1]struct{}
	_ = x[CkNone-(0)]
	_ = x[CkText-(1)]
	_ = x[CkMethod-(2)]
	_ = x[CkFunction-(3)]
	_ = x[CkConstructor-(4)]
	_ = x[CkField-(5)]
	_ = x[CkVariable-(6)]
	_ = x[CkClass-(7)]
	_ = x[CkInterface-(8)]
	_ = x[CkModule-(9)]
	_ = x[CkProperty-(10)]
	_ = x[CkUnit-(11)]
	_ = x[CkValue-(12)]
	_ = x[CkEnum-(13)]
	_ = x[CkKeyword-(14)]
	_ = x[CkSnippet-(15)]
	_ = x[Color-(16)]
	_ = x[CkFile-(17)]
	_ = x[CkReference-(18)]
	_ = x[CkFolder-(19)]
	_ = x[CkEnumMember-(20)]
	_ = x[CkConstant-(21)]
	_ = x[CkStruct-(22)]
	_ = x[CkEvent-(23)]
	_ = x[CkOperator-(24)]
	_ = x[CkTypeParameter-(25)]
}

var _CompletionKindNameToValueMap = map[string]CompletionKind{
	`None`:          0,
	`none`:          0,
	`Text`:          1,
	`text`:          1,
	`Method`:        2,
	`method`:        2,
	`Function`:      3,
	`function`:      3,
	`Constructor`:   4,
	`constructor`:   4,
	`Field`:         5,
	`field`:         5,
	`Variable`:      6,
	`variable`:      6,
	`Class`:         7,
	`class`:         7,
	`Interface`:     8,
	`interface`:     8,
	`Module`:        9,
	`module`:        9,
	`Property`:      10,
	`property`:      10,
	`Unit`:          11,
	`unit`:          11,
	`Value`:         12,
	`value`:         12,
	`Enum`:          13,
	`enum`:          13,
	`Keyword`:       14,
	`keyword`:       14,
	`Snippet`:       15,
	`snippet`:       15,
	`Color`:         16,
	`color`:         16,
	`File`:          17,
	`file`:          17,
	`Reference`:     18,
	`reference`:     18,
	`Folder`:        19,
	`folder`:        19,
	`EnumMember`:    20,
	`enummember`:    20,
	`Constant`:      21,
	`constant`:      21,
	`Struct`:        22,
	`struct`:        22,
	`Event`:         23,
	`event`:         23,
	`Operator`:      24,
	`operator`:      24,
	`TypeParameter`: 25,
	`typeparameter`: 25,
}

var _CompletionKindDescMap = map[CompletionKind]string{
	0:  ``,
	1:  ``,
	2:  ``,
	3:  ``,
	4:  ``,
	5:  ``,
	6:  ``,
	7:  ``,
	8:  ``,
	9:  ``,
	10: ``,
	11: ``,
	12: ``,
	13: ``,
	14: ``,
	15: ``,
	16: ``,
	17: ``,
	18: ``,
	19: ``,
	20: ``,
	21: ``,
	22: ``,
	23: ``,
	24: ``,
	25: ``,
}

var _CompletionKindMap = map[CompletionKind]string{
	0:  `None`,
	1:  `Text`,
	2:  `Method`,
	3:  `Function`,
	4:  `Constructor`,
	5:  `Field`,
	6:  `Variable`,
	7:  `Class`,
	8:  `Interface`,
	9:  `Module`,
	10: `Property`,
	11: `Unit`,
	12: `Value`,
	13: `Enum`,
	14: `Keyword`,
	15: `Snippet`,
	16: `Color`,
	17: `File`,
	18: `Reference`,
	19: `Folder`,
	20: `EnumMember`,
	21: `Constant`,
	22: `Struct`,
	23: `Event`,
	24: `Operator`,
	25: `TypeParameter`,
}

// String returns the string representation
// of this CompletionKind value.
func (i CompletionKind) String() string {
	if str, ok := _CompletionKindMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the CompletionKind value from its
// string representation, and returns an
// error if the string is invalid.
func (i *CompletionKind) SetString(s string) error {
	if val, ok := _CompletionKindNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _CompletionKindNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type CompletionKind")
}

// Int64 returns the CompletionKind value as an int64.
func (i CompletionKind) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the CompletionKind value from an int64.
func (i *CompletionKind) SetInt64(in int64) {
	*i = CompletionKind(in)
}

// Desc returns the description of the CompletionKind value.
func (i CompletionKind) Desc() string {
	if str, ok := _CompletionKindDescMap[i]; ok {
		return str
	}
	return i.String()
}

// CompletionKindValues returns all possible values
// for the type CompletionKind.
func CompletionKindValues() []CompletionKind {
	return _CompletionKindValues
}

// Values returns all possible values
// for the type CompletionKind.
func (i CompletionKind) Values() []enums.Enum {
	res := make([]enums.Enum, len(_CompletionKindValues))
	for i, d := range _CompletionKindValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type CompletionKind.
func (i CompletionKind) IsValid() bool {
	_, ok := _CompletionKindMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i CompletionKind) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *CompletionKind) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _SymbolKindValues = []SymbolKind{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}

// SymbolKindN is the highest valid value
// for type SymbolKind, plus one.
const SymbolKindN SymbolKind = 27

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _SymbolKindNoOp() {
	var x [1]struct{}
	_ = x[NoSymbolKind-(0)]
	_ = x[File-(1)]
	_ = x[Module-(2)]
	_ = x[Namespace-(3)]
	_ = x[Package-(4)]
	_ = x[Class-(5)]
	_ = x[Method-(6)]
	_ = x[Property-(7)]
	_ = x[Field-(8)]
	_ = x[Constructor-(9)]
	_ = x[Enum-(10)]
	_ = x[Interface-(11)]
	_ = x[Function-(12)]
	_ = x[Variable-(13)]
	_ = x[Constant-(14)]
	_ = x[String-(15)]
	_ = x[Number-(16)]
	_ = x[Boolean-(17)]
	_ = x[Array-(18)]
	_ = x[Object-(19)]
	_ = x[Key-(20)]
	_ = x[Null-(21)]
	_ = x[EnumMember-(22)]
	_ = x[Struct-(23)]
	_ = x[Event-(24)]
	_ = x[Operator-(25)]
	_ = x[TypeParameter-(26)]
}

var _SymbolKindNameToValueMap = map[string]SymbolKind{
	`NoSymbolKind`: 0,
	`nosymbolkind`: 0,
	`1 in LSP`:     1,
	`1 in lsp`:     1,
	`Module`:       2,
	`module`:       2,
	`Namespace`:    3,
	`namespace`:    3,
	`Package`:      4,
	`package`:      4,
	`Class`:        5,
	`class`:        5,
	`Method`:       6,
	`method`:       6,
	`Property`:     7,
	`property`:     7,
	`Field`:        8,
	`field`:        8,
	`Constructor`:  9,
	`constructor`:  9,
	`Enum`:         10,
	`enum`:         10,
	`Interface`:    11,
	`interface`:    11,
	`Function`:     12,
	`function`:     12,
	`Variable`:     13,
	`variable`:     13,
	`Constant`:     14,
	`constant`:     14,
	`String`:       15,
	`string`:       15,
	`Number`:       16,
	`number`:       16,
	`Boolean`:      17,
	`boolean`:      17,
	`Array`:        18,
	`array`:        18,
	`Object`:       19,
	`object`:       19,
	`Key`:          20,
	`key`:          20,
	`Null`:         21,
	`null`:         21,
	`EnumMember`:   22,
	`enummember`:   22,
	`Struct`:       23,
	`struct`:       23,
	`Event`:        24,
	`event`:        24,
	`Operator`:     25,
	`operator`:     25,
	`26 in LSP`:    26,
	`26 in lsp`:    26,
}

var _SymbolKindDescMap = map[SymbolKind]string{
	0:  ``,
	1:  ``,
	2:  ``,
	3:  ``,
	4:  ``,
	5:  ``,
	6:  ``,
	7:  ``,
	8:  ``,
	9:  ``,
	10: ``,
	11: ``,
	12: ``,
	13: ``,
	14: ``,
	15: ``,
	16: ``,
	17: ``,
	18: ``,
	19: ``,
	20: ``,
	21: ``,
	22: ``,
	23: ``,
	24: ``,
	25: ``,
	26: ``,
}

var _SymbolKindMap = map[SymbolKind]string{
	0:  `NoSymbolKind`,
	1:  `1 in LSP`,
	2:  `Module`,
	3:  `Namespace`,
	4:  `Package`,
	5:  `Class`,
	6:  `Method`,
	7:  `Property`,
	8:  `Field`,
	9:  `Constructor`,
	10: `Enum`,
	11: `Interface`,
	12: `Function`,
	13: `Variable`,
	14: `Constant`,
	15: `String`,
	16: `Number`,
	17: `Boolean`,
	18: `Array`,
	19: `Object`,
	20: `Key`,
	21: `Null`,
	22: `EnumMember`,
	23: `Struct`,
	24: `Event`,
	25: `Operator`,
	26: `26 in LSP`,
}

// String returns the string representation
// of this SymbolKind value.
func (i SymbolKind) String() string {
	if str, ok := _SymbolKindMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the SymbolKind value from its
// string representation, and returns an
// error if the string is invalid.
func (i *SymbolKind) SetString(s string) error {
	if val, ok := _SymbolKindNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _SymbolKindNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type SymbolKind")
}

// Int64 returns the SymbolKind value as an int64.
func (i SymbolKind) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the SymbolKind value from an int64.
func (i *SymbolKind) SetInt64(in int64) {
	*i = SymbolKind(in)
}

// Desc returns the description of the SymbolKind value.
func (i SymbolKind) Desc() string {
	if str, ok := _SymbolKindDescMap[i]; ok {
		return str
	}
	return i.String()
}

// SymbolKindValues returns all possible values
// for the type SymbolKind.
func SymbolKindValues() []SymbolKind {
	return _SymbolKindValues
}

// Values returns all possible values
// for the type SymbolKind.
func (i SymbolKind) Values() []enums.Enum {
	res := make([]enums.Enum, len(_SymbolKindValues))
	for i, d := range _SymbolKindValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type SymbolKind.
func (i SymbolKind) IsValid() bool {
	_, ok := _SymbolKindMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SymbolKind) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SymbolKind) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
