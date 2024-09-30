// Code generated by 'yaegi extract cogentcore.org/core/tensor/stats/stats'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/tensor/stats/stats"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/tensor/stats/stats/stats"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AsStatsFunc":        reflect.ValueOf(stats.AsStatsFunc),
		"Binarize":           reflect.ValueOf(stats.Binarize),
		"BinarizeOut":        reflect.ValueOf(stats.BinarizeOut),
		"Clamp":              reflect.ValueOf(stats.Clamp),
		"ClampOut":           reflect.ValueOf(stats.ClampOut),
		"Count":              reflect.ValueOf(stats.Count),
		"CountOut":           reflect.ValueOf(stats.CountOut),
		"CountOut64":         reflect.ValueOf(stats.CountOut64),
		"Describe":           reflect.ValueOf(stats.Describe),
		"DescribeTable":      reflect.ValueOf(stats.DescribeTable),
		"DescribeTableAll":   reflect.ValueOf(stats.DescribeTableAll),
		"DescriptiveStats":   reflect.ValueOf(&stats.DescriptiveStats).Elem(),
		"GroupAll":           reflect.ValueOf(stats.GroupAll),
		"GroupDescribe":      reflect.ValueOf(stats.GroupDescribe),
		"GroupStats":         reflect.ValueOf(stats.GroupStats),
		"Groups":             reflect.ValueOf(stats.Groups),
		"Max":                reflect.ValueOf(stats.Max),
		"MaxAbs":             reflect.ValueOf(stats.MaxAbs),
		"MaxAbsOut":          reflect.ValueOf(stats.MaxAbsOut),
		"MaxOut":             reflect.ValueOf(stats.MaxOut),
		"Mean":               reflect.ValueOf(stats.Mean),
		"MeanOut":            reflect.ValueOf(stats.MeanOut),
		"MeanOut64":          reflect.ValueOf(stats.MeanOut64),
		"Median":             reflect.ValueOf(stats.Median),
		"MedianOut":          reflect.ValueOf(stats.MedianOut),
		"Min":                reflect.ValueOf(stats.Min),
		"MinAbs":             reflect.ValueOf(stats.MinAbs),
		"MinAbsOut":          reflect.ValueOf(stats.MinAbsOut),
		"MinOut":             reflect.ValueOf(stats.MinOut),
		"NFunc":              reflect.ValueOf(stats.NFunc),
		"NormL1":             reflect.ValueOf(stats.NormL1),
		"NormL1Out":          reflect.ValueOf(stats.NormL1Out),
		"NormL2":             reflect.ValueOf(stats.NormL2),
		"NormL2Out":          reflect.ValueOf(stats.NormL2Out),
		"NormL2Out64":        reflect.ValueOf(stats.NormL2Out64),
		"Prod":               reflect.ValueOf(stats.Prod),
		"ProdOut":            reflect.ValueOf(stats.ProdOut),
		"Q1":                 reflect.ValueOf(stats.Q1),
		"Q1Out":              reflect.ValueOf(stats.Q1Out),
		"Q3":                 reflect.ValueOf(stats.Q3),
		"Q3Out":              reflect.ValueOf(stats.Q3Out),
		"Quantiles":          reflect.ValueOf(stats.Quantiles),
		"QuantilesOut":       reflect.ValueOf(stats.QuantilesOut),
		"Sem":                reflect.ValueOf(stats.Sem),
		"SemOut":             reflect.ValueOf(stats.SemOut),
		"SemPop":             reflect.ValueOf(stats.SemPop),
		"SemPopOut":          reflect.ValueOf(stats.SemPopOut),
		"StatCount":          reflect.ValueOf(stats.StatCount),
		"StatMax":            reflect.ValueOf(stats.StatMax),
		"StatMaxAbs":         reflect.ValueOf(stats.StatMaxAbs),
		"StatMean":           reflect.ValueOf(stats.StatMean),
		"StatMedian":         reflect.ValueOf(stats.StatMedian),
		"StatMin":            reflect.ValueOf(stats.StatMin),
		"StatMinAbs":         reflect.ValueOf(stats.StatMinAbs),
		"StatNormL1":         reflect.ValueOf(stats.StatNormL1),
		"StatNormL2":         reflect.ValueOf(stats.StatNormL2),
		"StatProd":           reflect.ValueOf(stats.StatProd),
		"StatQ1":             reflect.ValueOf(stats.StatQ1),
		"StatQ3":             reflect.ValueOf(stats.StatQ3),
		"StatSem":            reflect.ValueOf(stats.StatSem),
		"StatSemPop":         reflect.ValueOf(stats.StatSemPop),
		"StatStd":            reflect.ValueOf(stats.StatStd),
		"StatStdPop":         reflect.ValueOf(stats.StatStdPop),
		"StatSum":            reflect.ValueOf(stats.StatSum),
		"StatSumSq":          reflect.ValueOf(stats.StatSumSq),
		"StatVar":            reflect.ValueOf(stats.StatVar),
		"StatVarPop":         reflect.ValueOf(stats.StatVarPop),
		"StatsN":             reflect.ValueOf(stats.StatsN),
		"StatsValues":        reflect.ValueOf(stats.StatsValues),
		"Std":                reflect.ValueOf(stats.Std),
		"StdOut":             reflect.ValueOf(stats.StdOut),
		"StdOut64":           reflect.ValueOf(stats.StdOut64),
		"StdPop":             reflect.ValueOf(stats.StdPop),
		"StdPopOut":          reflect.ValueOf(stats.StdPopOut),
		"StripPackage":       reflect.ValueOf(stats.StripPackage),
		"Sum":                reflect.ValueOf(stats.Sum),
		"SumOut":             reflect.ValueOf(stats.SumOut),
		"SumOut64":           reflect.ValueOf(stats.SumOut64),
		"SumSq":              reflect.ValueOf(stats.SumSq),
		"SumSqDevOut64":      reflect.ValueOf(stats.SumSqDevOut64),
		"SumSqOut":           reflect.ValueOf(stats.SumSqOut),
		"SumSqOut64":         reflect.ValueOf(stats.SumSqOut64),
		"SumSqScaleOut64":    reflect.ValueOf(stats.SumSqScaleOut64),
		"TableGroupDescribe": reflect.ValueOf(stats.TableGroupDescribe),
		"TableGroupStats":    reflect.ValueOf(stats.TableGroupStats),
		"TableGroups":        reflect.ValueOf(stats.TableGroups),
		"UnitNorm":           reflect.ValueOf(stats.UnitNorm),
		"UnitNormOut":        reflect.ValueOf(stats.UnitNormOut),
		"Var":                reflect.ValueOf(stats.Var),
		"VarOut":             reflect.ValueOf(stats.VarOut),
		"VarOut64":           reflect.ValueOf(stats.VarOut64),
		"VarPop":             reflect.ValueOf(stats.VarPop),
		"VarPopOut":          reflect.ValueOf(stats.VarPopOut),
		"VarPopOut64":        reflect.ValueOf(stats.VarPopOut64),
		"Vec2inFunc":         reflect.ValueOf(stats.Vec2inFunc),
		"Vec2outFunc":        reflect.ValueOf(stats.Vec2outFunc),
		"VecFunc":            reflect.ValueOf(stats.VecFunc),
		"Vectorize2Out64":    reflect.ValueOf(stats.Vectorize2Out64),
		"VectorizeOut64":     reflect.ValueOf(stats.VectorizeOut64),
		"ZScore":             reflect.ValueOf(stats.ZScore),
		"ZScoreOut":          reflect.ValueOf(stats.ZScoreOut),

		// type definitions
		"Stats":        reflect.ValueOf((*stats.Stats)(nil)),
		"StatsFunc":    reflect.ValueOf((*stats.StatsFunc)(nil)),
		"StatsOutFunc": reflect.ValueOf((*stats.StatsOutFunc)(nil)),
	}
}
