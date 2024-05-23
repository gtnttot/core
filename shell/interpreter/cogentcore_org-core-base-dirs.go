// Code generated by 'yaegi extract cogentcore.org/core/base/dirs'. DO NOT EDIT.

package interpreter

import (
	"cogentcore.org/core/base/dirs"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/base/dirs/dirs"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AllFiles":         reflect.ValueOf(dirs.AllFiles),
		"AllFilesGlob":     reflect.ValueOf(dirs.AllFilesGlob),
		"DirAndFile":       reflect.ValueOf(dirs.DirAndFile),
		"DirFS":            reflect.ValueOf(dirs.DirFS),
		"Dirs":             reflect.ValueOf(dirs.Dirs),
		"ExtFilenames":     reflect.ValueOf(dirs.ExtFilenames),
		"ExtFilenamesFS":   reflect.ValueOf(dirs.ExtFilenamesFS),
		"ExtFiles":         reflect.ValueOf(dirs.ExtFiles),
		"FileExists":       reflect.ValueOf(dirs.FileExists),
		"FileExistsFS":     reflect.ValueOf(dirs.FileExistsFS),
		"FindFilesOnPaths": reflect.ValueOf(dirs.FindFilesOnPaths),
		"GoSrcDir":         reflect.ValueOf(dirs.GoSrcDir),
		"HasFile":          reflect.ValueOf(dirs.HasFile),
		"LatestMod":        reflect.ValueOf(dirs.LatestMod),
		"RelFilePath":      reflect.ValueOf(dirs.RelFilePath),
		"SplitExt":         reflect.ValueOf(dirs.SplitExt),
	}
}
