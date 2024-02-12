// Code generated by "core generate"; DO NOT EDIT.

package fi

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/fi.FileInfo", IDName: "file-info", Doc: "FileInfo represents the information about a given file / directory,\nincluding icon, mimetype, etc", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Methods: []gti.Method{{Name: "Duplicate", Doc: "Duplicate creates a copy of given file -- only works for regular files, not\ndirectories.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Returns: []string{"string", "error"}}, {Name: "Delete", Doc: "Delete moves the file to the trash / recycling bin.\nOn mobile and web, it deletes it directly.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Returns: []string{"error"}}, {Name: "Rename", Doc: "Rename renames (moves) this file to given new path name.\nUpdates the FileInfo setting to the new name, although it might\nbe out of scope if it moved into a new path", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"path"}, Returns: []string{"newpath", "err"}}}, Fields: []gti.Field{{Name: "Ic", Doc: "icon for file"}, {Name: "Name", Doc: "name of the file, without any path"}, {Name: "Size", Doc: "size of the file"}, {Name: "Kind", Doc: "type of file / directory; shorter, more user-friendly\nversion of mime type, based on category"}, {Name: "Mime", Doc: "full official mime type of the contents"}, {Name: "Cat", Doc: "functional category of the file, based on mime data etc"}, {Name: "Known", Doc: "known file type"}, {Name: "Mode", Doc: "file mode bits"}, {Name: "ModTime", Doc: "time that contents (only) were last modified"}, {Name: "Vcs", Doc: "version control system status, when enabled"}, {Name: "Path", Doc: "full path to file, including name; for file functions"}}})
