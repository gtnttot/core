// Code generated by "core generate -add-types"; DO NOT EDIT.

package randx

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/base/randx.Rand", IDName: "rand", Doc: "Rand provides an interface with most of the standard\nrand.Rand methods, to support the use of either the\nglobal rand generator or a separate Rand source."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/base/randx.SysRand", IDName: "sys-rand", Doc: "SysRand supports the system random number generator\nfor either a separate rand.Rand source, or, if that\nis nil, the global rand stream.", Fields: []types.Field{{Name: "Rand", Doc: "if non-nil, use this random number source instead of the global default one"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/base/randx.RandParams", IDName: "rand-params", Doc: "RandParams provides parameterized random number generation according to different distributions\nand variance, mean params", Directives: []types.Directive{{Tool: "git", Directive: "add"}}, Fields: []types.Field{{Name: "Dist", Doc: "distribution to generate random numbers from"}, {Name: "Mean", Doc: "mean of random distribution -- typically added to generated random variants"}, {Name: "Var", Doc: "variability parameter for the random numbers (gauss = standard deviation, not variance; uniform = half-range, others as noted in RandDists)"}, {Name: "Par", Doc: "extra parameter for distribution (depends on each one)"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/base/randx.RandDists", IDName: "rand-dists", Doc: "RandDists are different random number distributions"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/base/randx.Seeds", IDName: "seeds", Doc: "Seeds is a set of random seeds, typically used one per Run"})
