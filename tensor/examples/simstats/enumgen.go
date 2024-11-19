// Code generated by "core generate"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/enums"
)

var _TimesValues = []Times{0, 1, 2}

// TimesN is the highest valid value for type Times, plus one.
const TimesN Times = 3

var _TimesValueMap = map[string]Times{`Trial`: 0, `Epoch`: 1, `Run`: 2}

var _TimesDescMap = map[Times]string{0: ``, 1: ``, 2: ``}

var _TimesMap = map[Times]string{0: `Trial`, 1: `Epoch`, 2: `Run`}

// String returns the string representation of this Times value.
func (i Times) String() string { return enums.String(i, _TimesMap) }

// SetString sets the Times value from its string representation,
// and returns an error if the string is invalid.
func (i *Times) SetString(s string) error { return enums.SetString(i, s, _TimesValueMap, "Times") }

// Int64 returns the Times value as an int64.
func (i Times) Int64() int64 { return int64(i) }

// SetInt64 sets the Times value from an int64.
func (i *Times) SetInt64(in int64) { *i = Times(in) }

// Desc returns the description of the Times value.
func (i Times) Desc() string { return enums.Desc(i, _TimesDescMap) }

// TimesValues returns all possible values for the type Times.
func TimesValues() []Times { return _TimesValues }

// Values returns all possible values for the type Times.
func (i Times) Values() []enums.Enum { return enums.Values(_TimesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Times) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Times) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Times") }

var _LoopPhaseValues = []LoopPhase{0, 1}

// LoopPhaseN is the highest valid value for type LoopPhase, plus one.
const LoopPhaseN LoopPhase = 2

var _LoopPhaseValueMap = map[string]LoopPhase{`Start`: 0, `Step`: 1}

var _LoopPhaseDescMap = map[LoopPhase]string{0: `Start is the start of the loop: resets accumulated stats, initializes.`, 1: `Step is each iteration of the loop.`}

var _LoopPhaseMap = map[LoopPhase]string{0: `Start`, 1: `Step`}

// String returns the string representation of this LoopPhase value.
func (i LoopPhase) String() string { return enums.String(i, _LoopPhaseMap) }

// SetString sets the LoopPhase value from its string representation,
// and returns an error if the string is invalid.
func (i *LoopPhase) SetString(s string) error {
	return enums.SetString(i, s, _LoopPhaseValueMap, "LoopPhase")
}

// Int64 returns the LoopPhase value as an int64.
func (i LoopPhase) Int64() int64 { return int64(i) }

// SetInt64 sets the LoopPhase value from an int64.
func (i *LoopPhase) SetInt64(in int64) { *i = LoopPhase(in) }

// Desc returns the description of the LoopPhase value.
func (i LoopPhase) Desc() string { return enums.Desc(i, _LoopPhaseDescMap) }

// LoopPhaseValues returns all possible values for the type LoopPhase.
func LoopPhaseValues() []LoopPhase { return _LoopPhaseValues }

// Values returns all possible values for the type LoopPhase.
func (i LoopPhase) Values() []enums.Enum { return enums.Values(_LoopPhaseValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i LoopPhase) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *LoopPhase) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "LoopPhase")
}
