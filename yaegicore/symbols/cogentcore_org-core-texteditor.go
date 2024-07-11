// Code generated by 'yaegi extract cogentcore.org/core/texteditor'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/texteditor"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/texteditor/texteditor"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AsEditor":                 reflect.ValueOf(texteditor.AsEditor),
		"BufferClosed":             reflect.ValueOf(texteditor.BufferClosed),
		"BufferDelete":             reflect.ValueOf(texteditor.BufferDelete),
		"BufferDone":               reflect.ValueOf(texteditor.BufferDone),
		"BufferInsert":             reflect.ValueOf(texteditor.BufferInsert),
		"BufferMarkupUpdated":      reflect.ValueOf(texteditor.BufferMarkupUpdated),
		"BufferMods":               reflect.ValueOf(texteditor.BufferMods),
		"BufferNew":                reflect.ValueOf(texteditor.BufferNew),
		"BufferSignalsN":           reflect.ValueOf(texteditor.BufferSignalsN),
		"BufferSignalsValues":      reflect.ValueOf(texteditor.BufferSignalsValues),
		"ChromaTagsForLine":        reflect.ValueOf(texteditor.ChromaTagsForLine),
		"ChromaTagsLine":           reflect.ValueOf(texteditor.ChromaTagsLine),
		"ClipHistMax":              reflect.ValueOf(&texteditor.ClipHistMax).Elem(),
		"CompleteEditParse":        reflect.ValueOf(texteditor.CompleteEditParse),
		"CompleteParse":            reflect.ValueOf(texteditor.CompleteParse),
		"CompleteText":             reflect.ValueOf(texteditor.CompleteText),
		"CompleteTextEdit":         reflect.ValueOf(texteditor.CompleteTextEdit),
		"DiffEditorDialog":         reflect.ValueOf(texteditor.DiffEditorDialog),
		"DiffEditorDialogFromRevs": reflect.ValueOf(texteditor.DiffEditorDialogFromRevs),
		"DiffFiles":                reflect.ValueOf(texteditor.DiffFiles),
		"DiffRevertDiffs":          reflect.ValueOf(&texteditor.DiffRevertDiffs).Elem(),
		"DiffRevertLines":          reflect.ValueOf(&texteditor.DiffRevertLines).Elem(),
		"EditNoSignal":             reflect.ValueOf(texteditor.EditNoSignal),
		"EditSignal":               reflect.ValueOf(texteditor.EditSignal),
		"EditorBlinker":            reflect.ValueOf(&texteditor.EditorBlinker).Elem(),
		"EditorSpriteName":         reflect.ValueOf(&texteditor.EditorSpriteName).Elem(),
		"HTMLEscapeBytes":          reflect.ValueOf(texteditor.HTMLEscapeBytes),
		"HTMLEscapeRunes":          reflect.ValueOf(texteditor.HTMLEscapeRunes),
		"LookupParse":              reflect.ValueOf(texteditor.LookupParse),
		"MarkupDelay":              reflect.ValueOf(&texteditor.MarkupDelay).Elem(),
		"MaxLineLen":               reflect.ValueOf(constant.MakeFromLiteral("67108864", token.INT, 0)),
		"MaxNTags":                 reflect.ValueOf(constant.MakeFromLiteral("1024", token.INT, 0)),
		"MaxScopeLines":            reflect.ValueOf(&texteditor.MaxScopeLines).Elem(),
		"NewBuffer":                reflect.ValueOf(texteditor.NewBuffer),
		"NewDiffEditor":            reflect.ValueOf(texteditor.NewDiffEditor),
		"NewDiffTextEditor":        reflect.ValueOf(texteditor.NewDiffTextEditor),
		"NewEditor":                reflect.ValueOf(texteditor.NewEditor),
		"NewSoloEditor":            reflect.ValueOf(texteditor.NewSoloEditor),
		"NewTwinEditors":           reflect.ValueOf(texteditor.NewTwinEditors),
		"PrevISearchString":        reflect.ValueOf(&texteditor.PrevISearchString).Elem(),
		"PrevQReplaceFinds":        reflect.ValueOf(&texteditor.PrevQReplaceFinds).Elem(),
		"PrevQReplaceRepls":        reflect.ValueOf(&texteditor.PrevQReplaceRepls).Elem(),
		"ReplaceMatchCase":         reflect.ValueOf(texteditor.ReplaceMatchCase),
		"ReplaceNoMatchCase":       reflect.ValueOf(texteditor.ReplaceNoMatchCase),
		"TextDialog":               reflect.ValueOf(texteditor.TextDialog),
		"ViewClipHistAdd":          reflect.ValueOf(texteditor.ViewClipHistAdd),
		"ViewClipHistChooseLen":    reflect.ValueOf(&texteditor.ViewClipHistChooseLen).Elem(),
		"ViewClipHistChooseList":   reflect.ValueOf(texteditor.ViewClipHistChooseList),
		"ViewClipHistory":          reflect.ValueOf(&texteditor.ViewClipHistory).Elem(),
		"ViewClipRect":             reflect.ValueOf(&texteditor.ViewClipRect).Elem(),
		"ViewDepthColors":          reflect.ValueOf(&texteditor.ViewDepthColors).Elem(),
		"ViewMaxFindHighlights":    reflect.ValueOf(&texteditor.ViewMaxFindHighlights).Elem(),

		// type definitions
		"Buffer":                 reflect.ValueOf((*texteditor.Buffer)(nil)),
		"BufferSignals":          reflect.ValueOf((*texteditor.BufferSignals)(nil)),
		"DiffEditor":             reflect.ValueOf((*texteditor.DiffEditor)(nil)),
		"DiffTextEditor":         reflect.ValueOf((*texteditor.DiffTextEditor)(nil)),
		"Editor":                 reflect.ValueOf((*texteditor.Editor)(nil)),
		"EditorEmbedder":         reflect.ValueOf((*texteditor.EditorEmbedder)(nil)),
		"HiMarkup":               reflect.ValueOf((*texteditor.HiMarkup)(nil)),
		"ISearch":                reflect.ValueOf((*texteditor.ISearch)(nil)),
		"OutputBuffer":           reflect.ValueOf((*texteditor.OutputBuffer)(nil)),
		"OutputBufferMarkupFunc": reflect.ValueOf((*texteditor.OutputBufferMarkupFunc)(nil)),
		"QReplace":               reflect.ValueOf((*texteditor.QReplace)(nil)),
		"TwinEditors":            reflect.ValueOf((*texteditor.TwinEditors)(nil)),

		// interface wrapper definitions
		"_EditorEmbedder": reflect.ValueOf((*_cogentcore_org_core_texteditor_EditorEmbedder)(nil)),
	}
}

// _cogentcore_org_core_texteditor_EditorEmbedder is an interface wrapper for EditorEmbedder type
type _cogentcore_org_core_texteditor_EditorEmbedder struct {
	IValue    interface{}
	WAsEditor func() *texteditor.Editor
}

func (W _cogentcore_org_core_texteditor_EditorEmbedder) AsEditor() *texteditor.Editor {
	return W.WAsEditor()
}
