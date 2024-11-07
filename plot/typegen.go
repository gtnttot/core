// Code generated by "core generate -add-types"; DO NOT EDIT.

package plot

import (
	"image"

	"cogentcore.org/core/math32/minmax"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Normalizer", IDName: "normalizer", Doc: "Normalizer rescales values from the data coordinate system to the\nnormalized coordinate system.", Methods: []types.Method{{Name: "Normalize", Doc: "Normalize transforms a value x in the data coordinate system to\nthe normalized coordinate system.", Args: []string{"min", "max", "x"}, Returns: []string{"float32"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Axis", IDName: "axis", Doc: "Axis represents either a horizontal or vertical\naxis of a plot.", Fields: []types.Field{{Name: "Min", Doc: "Min and Max are the minimum and maximum data\nvalues represented by the axis."}, {Name: "Max", Doc: "Min and Max are the minimum and maximum data\nvalues represented by the axis."}, {Name: "Axis", Doc: "specifies which axis this is: X or Y"}, {Name: "Label", Doc: "Label for the axis"}, {Name: "Line", Doc: "Line styling properties for the axis line."}, {Name: "Padding", Doc: "Padding between the axis line and the data.  Having\nnon-zero padding ensures that the data is never drawn\non the axis, thus making it easier to see."}, {Name: "TickText", Doc: "has the text style for rendering tick labels, and is shared for actual rendering"}, {Name: "TickLine", Doc: "line style for drawing tick lines"}, {Name: "TickLength", Doc: "length of tick lines"}, {Name: "Ticker", Doc: "Ticker generates the tick marks.  Any tick marks\nreturned by the Marker function that are not in\nrange of the axis are not drawn."}, {Name: "Scale", Doc: "Scale transforms a value given in the data coordinate system\nto the normalized coordinate system of the axis—its distance\nalong the axis as a fraction of the axis range."}, {Name: "AutoRescale", Doc: "AutoRescale enables an axis to automatically adapt its minimum\nand maximum boundaries, according to its underlying Ticker."}, {Name: "ticks", Doc: "cached list of ticks, set in size"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LinearScale", IDName: "linear-scale", Doc: "LinearScale an be used as the value of an Axis.Scale function to\nset the axis to a standard linear scale."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LogScale", IDName: "log-scale", Doc: "LogScale can be used as the value of an Axis.Scale function to\nset the axis to a log scale."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.InvertedScale", IDName: "inverted-scale", Doc: "InvertedScale can be used as the value of an Axis.Scale function to\ninvert the axis using any Normalizer.", Embeds: []types.Field{{Name: "Normalizer"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Valuer", IDName: "valuer", Doc: "Valuer provides an interface for a list of scalar values", Methods: []types.Method{{Name: "Len", Doc: "Len returns the number of values.", Returns: []string{"int"}}, {Name: "Value", Doc: "Value returns a value.", Args: []string{"i"}, Returns: []string{"float32"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Values", IDName: "values", Doc: "Values implements the Valuer interface."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.TensorValues", IDName: "tensor-values", Doc: "TensorValues provides a Valuer interface wrapper for a tensor.", Embeds: []types.Field{{Name: "Tensor"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYer", IDName: "x-yer", Doc: "XYer provides an interface for a list of X,Y data pairs", Methods: []types.Method{{Name: "Len", Doc: "Len returns the number of x, y pairs.", Returns: []string{"int"}}, {Name: "XY", Doc: "XY returns an x, y pair.", Args: []string{"i"}, Returns: []string{"x", "y"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYs", IDName: "x-ys", Doc: "XYs implements the XYer interface."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.TensorXYs", IDName: "tensor-x-ys", Doc: "TensorXYs provides a XYer interface wrapper for a tensor.", Fields: []types.Field{{Name: "X"}, {Name: "Y"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XValues", IDName: "x-values", Doc: "XValues implements the Valuer interface,\nreturning the x value from an XYer.", Embeds: []types.Field{{Name: "XYer"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.YValues", IDName: "y-values", Doc: "YValues implements the Valuer interface,\nreturning the y value from an XYer.", Embeds: []types.Field{{Name: "XYer"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYZer", IDName: "xy-zer", Doc: "XYZer provides an interface for a list of X,Y,Z data triples.\nIt also satisfies the XYer interface for the X,Y pairs.", Methods: []types.Method{{Name: "Len", Doc: "Len returns the number of x, y, z triples.", Returns: []string{"int"}}, {Name: "XYZ", Doc: "XYZ returns an x, y, z triple.", Args: []string{"i"}, Returns: []string{"float32", "float32", "float32"}}, {Name: "XY", Doc: "XY returns an x, y pair.", Args: []string{"i"}, Returns: []string{"float32", "float32"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYZs", IDName: "xy-zs", Doc: "XYZs implements the XYZer interface using a slice."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYZ", IDName: "xyz", Doc: "XYZ is an x, y and z value.", Fields: []types.Field{{Name: "X"}, {Name: "Y"}, {Name: "Z"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.XYValues", IDName: "xy-values", Doc: "XYValues implements the XYer interface, returning\nthe x and y values from an XYZer.", Embeds: []types.Field{{Name: "XYZer"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Labeler", IDName: "labeler", Doc: "Labeler provides an interface for a list of string labels", Methods: []types.Method{{Name: "Label", Doc: "Label returns a label.", Args: []string{"i"}, Returns: []string{"string"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.selection", IDName: "selection", Fields: []types.Field{{Name: "n", Doc: "n is the number of labels selected."}, {Name: "lMin", Doc: "lMin and lMax are the selected min\nand max label values. lq is the q\nchosen."}, {Name: "lMax", Doc: "lMin and lMax are the selected min\nand max label values. lq is the q\nchosen."}, {Name: "lStep", Doc: "lMin and lMax are the selected min\nand max label values. lq is the q\nchosen."}, {Name: "lq", Doc: "lMin and lMax are the selected min\nand max label values. lq is the q\nchosen."}, {Name: "score", Doc: "score is the score for the selection."}, {Name: "magnitude", Doc: "magnitude is the magnitude of the\nlabel step distance."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.weights", IDName: "weights", Doc: "weights is a helper type to calcuate the labelling scheme's total score.", Fields: []types.Field{{Name: "simplicity"}, {Name: "coverage"}, {Name: "density"}, {Name: "legibility"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LegendPosition", IDName: "legend-position", Doc: "LegendPosition specifies where to put the legend", Fields: []types.Field{{Name: "Top", Doc: "Top and Left specify the location of the legend."}, {Name: "Left", Doc: "Top and Left specify the location of the legend."}, {Name: "XOffs", Doc: "XOffs and YOffs are added to the legend's final position,\nrelative to the relevant anchor position"}, {Name: "YOffs", Doc: "XOffs and YOffs are added to the legend's final position,\nrelative to the relevant anchor position"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Legend", IDName: "legend", Doc: "A Legend gives a description of the meaning of different\ndata elements of the plot.  Each legend entry has a name\nand a thumbnail, where the thumbnail shows a small\nsample of the display style of the corresponding data.", Fields: []types.Field{{Name: "TextStyle", Doc: "TextStyle is the style given to the legend entry texts."}, {Name: "Position", Doc: "position of the legend"}, {Name: "ThumbnailWidth", Doc: "ThumbnailWidth is the width of legend thumbnails."}, {Name: "Fill", Doc: "Fill specifies the background fill color for the legend box,\nif non-nil."}, {Name: "Entries", Doc: "Entries are all of the LegendEntries described by this legend."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Thumbnailer", IDName: "thumbnailer", Doc: "Thumbnailer wraps the Thumbnail method, which\ndraws the small image in a legend representing the\nstyle of data.", Methods: []types.Method{{Name: "Thumbnail", Doc: "Thumbnail draws an thumbnail representing\na legend entry.  The thumbnail will usually show\na smaller representation of the style used\nto plot the corresponding data.", Args: []string{"pt"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LegendEntry", IDName: "legend-entry", Doc: "A LegendEntry represents a single line of a legend, it\nhas a name and an icon.", Fields: []types.Field{{Name: "Text", Doc: "text is the text associated with this entry."}, {Name: "Thumbs", Doc: "thumbs is a slice of all of the thumbnails styles"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LineStyle", IDName: "line-style", Doc: "LineStyle has style properties for drawing lines.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "On", Doc: "On indicates whether to plot lines."}, {Name: "Color", Doc: "Color is the stroke color image specification.\nSetting to nil turns line off."}, {Name: "Width", Doc: "Width is the line width, with a default of 1 Pt (point).\nSetting to 0 turns line off."}, {Name: "Dashes", Doc: "Dashes are the dashes of the stroke. Each pair of values specifies\nthe amount to paint and then the amount to skip."}, {Name: "Fill", Doc: "Fill is the color to fill solid regions, in a plot-specific\nway (e.g., the area below a Line plot, the bar color).\nUse nil to disable filling."}, {Name: "NegativeX", Doc: "NegativeX specifies whether to draw lines that connect points with a negative\nX-axis direction; otherwise there is a break in the line.\ndefault is false, so that repeated series of data across the X axis\nare plotted separately."}, {Name: "Step", Doc: "Step specifies how to step the line between points."}}})

// SetOn sets the [LineStyle.On]:
// On indicates whether to plot lines.
func (t *LineStyle) SetOn(v DefaultOffOn) *LineStyle { t.On = v; return t }

// SetColor sets the [LineStyle.Color]:
// Color is the stroke color image specification.
// Setting to nil turns line off.
func (t *LineStyle) SetColor(v image.Image) *LineStyle { t.Color = v; return t }

// SetWidth sets the [LineStyle.Width]:
// Width is the line width, with a default of 1 Pt (point).
// Setting to 0 turns line off.
func (t *LineStyle) SetWidth(v units.Value) *LineStyle { t.Width = v; return t }

// SetDashes sets the [LineStyle.Dashes]:
// Dashes are the dashes of the stroke. Each pair of values specifies
// the amount to paint and then the amount to skip.
func (t *LineStyle) SetDashes(v ...float32) *LineStyle { t.Dashes = v; return t }

// SetFill sets the [LineStyle.Fill]:
// Fill is the color to fill solid regions, in a plot-specific
// way (e.g., the area below a Line plot, the bar color).
// Use nil to disable filling.
func (t *LineStyle) SetFill(v image.Image) *LineStyle { t.Fill = v; return t }

// SetNegativeX sets the [LineStyle.NegativeX]:
// NegativeX specifies whether to draw lines that connect points with a negative
// X-axis direction; otherwise there is a break in the line.
// default is false, so that repeated series of data across the X axis
// are plotted separately.
func (t *LineStyle) SetNegativeX(v bool) *LineStyle { t.NegativeX = v; return t }

// SetStep sets the [LineStyle.Step]:
// Step specifies how to step the line between points.
func (t *LineStyle) SetStep(v StepKind) *LineStyle { t.Step = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.StepKind", IDName: "step-kind", Doc: "StepKind specifies a form of a connection of two consecutive points."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Plot", IDName: "plot", Doc: "Plot is the basic type representing a plot.\nIt renders into its own image.RGBA Pixels image,\nand can also save a corresponding SVG version.\nThe Axis ranges are updated automatically when plots\nare added, so setting a fixed range should happen\nafter that point.  See [UpdateRange] method as well.", Fields: []types.Field{{Name: "Title", Doc: "Title of the plot"}, {Name: "Background", Doc: "Background is the background of the plot.\nThe default is [colors.Scheme.Surface]."}, {Name: "StandardTextStyle", Doc: "standard text style with default options"}, {Name: "X", Doc: "X and Y are the horizontal and vertical axes\nof the plot respectively."}, {Name: "Y", Doc: "X and Y are the horizontal and vertical axes\nof the plot respectively."}, {Name: "Legend", Doc: "Legend is the plot's legend."}, {Name: "Plotters", Doc: "plotters are drawn by calling their Plot method\nafter the axes are drawn."}, {Name: "Size", Doc: "size is the target size of the image to render to"}, {Name: "DPI", Doc: "DPI is the dots per inch for rendering the image.\nLarger numbers result in larger scaling of the plot contents\nwhich is strongly recommended for print (e.g., use 300 for print)"}, {Name: "Paint", Doc: "painter for rendering"}, {Name: "Pixels", Doc: "pixels that we render into"}, {Name: "PlotBox", Doc: "Current plot bounding box in image coordinates, for plotting coordinates"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Plotter", IDName: "plotter", Doc: "Plotter is an interface that wraps the Plot method.\nSome standard implementations of Plotter can be found in plotters.", Methods: []types.Method{{Name: "Plot", Doc: "Plot draws the data to the Plot Paint", Args: []string{"pt"}}, {Name: "XYData", Doc: "returns the data for this plot as X,Y points,\nincluding corresponding pixel data.\nThis allows gui interface to inspect data etc.", Returns: []string{"data", "pixels"}}, {Name: "ApplyStyle", Doc: "ApplyStyle applies any stylers to this element."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.DataRanger", IDName: "data-ranger", Doc: "DataRanger wraps the DataRange method.", Methods: []types.Method{{Name: "DataRange", Doc: "DataRange returns the range of X and Y values.", Args: []string{"pt"}, Returns: []string{"xmin", "xmax", "ymin", "ymax"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.PointStyle", IDName: "point-style", Doc: "PointStyle has style properties for drawing points as different shapes.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "On", Doc: "On indicates whether to plot points."}, {Name: "Shape", Doc: "Shape to draw."}, {Name: "Color", Doc: "Color is the stroke color image specification.\nSetting to nil turns line off."}, {Name: "Fill", Doc: "Fill is the color to fill solid regions, in a plot-specific\nway (e.g., the area below a Line plot, the bar color).\nUse nil to disable filling."}, {Name: "Width", Doc: "Width is the line width, with a default of 1 Pt (point).\nSetting to 0 turns line off."}, {Name: "Size", Doc: "Size of shape to draw for each point.\nDefaults to 4 Pt (point)."}}})

// SetOn sets the [PointStyle.On]:
// On indicates whether to plot points.
func (t *PointStyle) SetOn(v DefaultOffOn) *PointStyle { t.On = v; return t }

// SetShape sets the [PointStyle.Shape]:
// Shape to draw.
func (t *PointStyle) SetShape(v Shapes) *PointStyle { t.Shape = v; return t }

// SetColor sets the [PointStyle.Color]:
// Color is the stroke color image specification.
// Setting to nil turns line off.
func (t *PointStyle) SetColor(v image.Image) *PointStyle { t.Color = v; return t }

// SetFill sets the [PointStyle.Fill]:
// Fill is the color to fill solid regions, in a plot-specific
// way (e.g., the area below a Line plot, the bar color).
// Use nil to disable filling.
func (t *PointStyle) SetFill(v image.Image) *PointStyle { t.Fill = v; return t }

// SetWidth sets the [PointStyle.Width]:
// Width is the line width, with a default of 1 Pt (point).
// Setting to 0 turns line off.
func (t *PointStyle) SetWidth(v units.Value) *PointStyle { t.Width = v; return t }

// SetSize sets the [PointStyle.Size]:
// Size of shape to draw for each point.
// Defaults to 4 Pt (point).
func (t *PointStyle) SetSize(v units.Value) *PointStyle { t.Size = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Shapes", IDName: "shapes"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Style", IDName: "style", Doc: "Style contains the plot styling properties relevant across\nmost plot types. These properties apply both to individual plot elements\nand to the plot as a whole.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "On", Doc: "On specifies whether to plot this item, for cases where it can be turned off."}, {Name: "Range", Doc: "Range is the effective range of data to plot, where either end can be fixed."}, {Name: "Label", Doc: "Label provides an alternative label to use for axis, if set."}, {Name: "NTicks", Doc: "NTicks sets the desired number of ticks for the axis, if > 0."}, {Name: "Line", Doc: "Line has style properties for drawing lines."}, {Name: "Point", Doc: "Point has style properties for drawing points."}, {Name: "Text", Doc: "Text has style properties for rendering text."}, {Name: "Width", Doc: "Width has various plot width properties."}}})

// SetOn sets the [Style.On]:
// On specifies whether to plot this item, for cases where it can be turned off.
func (t *Style) SetOn(v DefaultOffOn) *Style { t.On = v; return t }

// SetRange sets the [Style.Range]:
// Range is the effective range of data to plot, where either end can be fixed.
func (t *Style) SetRange(v minmax.Range32) *Style { t.Range = v; return t }

// SetLabel sets the [Style.Label]:
// Label provides an alternative label to use for axis, if set.
func (t *Style) SetLabel(v string) *Style { t.Label = v; return t }

// SetNTicks sets the [Style.NTicks]:
// NTicks sets the desired number of ticks for the axis, if > 0.
func (t *Style) SetNTicks(v int) *Style { t.NTicks = v; return t }

// SetLine sets the [Style.Line]:
// Line has style properties for drawing lines.
func (t *Style) SetLine(v LineStyle) *Style { t.Line = v; return t }

// SetPoint sets the [Style.Point]:
// Point has style properties for drawing points.
func (t *Style) SetPoint(v PointStyle) *Style { t.Point = v; return t }

// SetText sets the [Style.Text]:
// Text has style properties for rendering text.
func (t *Style) SetText(v TextStyle) *Style { t.Text = v; return t }

// SetWidth sets the [Style.Width]:
// Width has various plot width properties.
func (t *Style) SetWidth(v WidthStyle) *Style { t.Width = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.WidthStyle", IDName: "width-style", Doc: "WidthStyle contains various plot width properties relevant across\ndifferent plot types.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "Cap", Doc: "Cap is the width of the caps drawn at the top of error bars.\nThe default is 10dp"}, {Name: "Offset", Doc: "Offset for Bar plot is the offset added to each X axis value\nrelative to the Stride computed value (X = offset + index * Stride)\nDefaults to 1."}, {Name: "Stride", Doc: "Stride for Bar plot is distance between bars. Defaults to 1."}, {Name: "Width", Doc: "Width for Bar plot is the width of the bars, which should be less than\nthe Stride to prevent bar overlap.\nDefaults to .8"}, {Name: "Pad", Doc: "Pad for Bar plot is additional space at start / end of data range,\nto keep bars from overflowing ends. This amount is subtracted from Offset\nand added to (len(Values)-1)*Stride -- no other accommodation for bar\nwidth is provided, so that should be built into this value as well.\nDefaults to 1."}}})

// SetCap sets the [WidthStyle.Cap]:
// Cap is the width of the caps drawn at the top of error bars.
// The default is 10dp
func (t *WidthStyle) SetCap(v units.Value) *WidthStyle { t.Cap = v; return t }

// SetOffset sets the [WidthStyle.Offset]:
// Offset for Bar plot is the offset added to each X axis value
// relative to the Stride computed value (X = offset + index * Stride)
// Defaults to 1.
func (t *WidthStyle) SetOffset(v float32) *WidthStyle { t.Offset = v; return t }

// SetStride sets the [WidthStyle.Stride]:
// Stride for Bar plot is distance between bars. Defaults to 1.
func (t *WidthStyle) SetStride(v float32) *WidthStyle { t.Stride = v; return t }

// SetWidth sets the [WidthStyle.Width]:
// Width for Bar plot is the width of the bars, which should be less than
// the Stride to prevent bar overlap.
// Defaults to .8
func (t *WidthStyle) SetWidth(v float32) *WidthStyle { t.Width = v; return t }

// SetPad sets the [WidthStyle.Pad]:
// Pad for Bar plot is additional space at start / end of data range,
// to keep bars from overflowing ends. This amount is subtracted from Offset
// and added to (len(Values)-1)*Stride -- no other accommodation for bar
// width is provided, so that should be built into this value as well.
// Defaults to 1.
func (t *WidthStyle) SetPad(v float32) *WidthStyle { t.Pad = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Stylers", IDName: "stylers", Doc: "Stylers is a list of styling functions that set Style properties.\nThese are called in the order added."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.DefaultOffOn", IDName: "default-off-on", Doc: "DefaultOffOn specifies whether to use the default value for a bool option,\nor to override the default and set Off or On."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.TextStyle", IDName: "text-style", Doc: "TextStyle specifies styling parameters for Text elements", Embeds: []types.Field{{Name: "FontRender"}}, Fields: []types.Field{{Name: "Align", Doc: "how to align text along the relevant dimension for the text element"}, {Name: "Padding", Doc: "Padding is used in a case-dependent manner to add space around text elements"}, {Name: "Rotation", Doc: "rotation of the text, in Degrees"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Text", IDName: "text", Doc: "Text specifies a single text element in a plot", Fields: []types.Field{{Name: "Text", Doc: "text string, which can use HTML formatting"}, {Name: "Style", Doc: "styling for this text element"}, {Name: "PaintText", Doc: "PaintText is the [paint.Text] for the text."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Tick", IDName: "tick", Doc: "A Tick is a single tick mark on an axis.", Fields: []types.Field{{Name: "Value", Doc: "Value is the data value marked by this Tick."}, {Name: "Label", Doc: "Label is the text to display at the tick mark.\nIf Label is an empty string then this is a minor tick mark."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.Ticker", IDName: "ticker", Doc: "Ticker creates Ticks in a specified range", Methods: []types.Method{{Name: "Ticks", Doc: "Ticks returns Ticks in a specified range", Args: []string{"min", "max"}, Returns: []string{"Tick"}}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.DefaultTicks", IDName: "default-ticks", Doc: "DefaultTicks is suitable for the Ticker field of an Axis,\nit returns a reasonable default set of tick marks."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.LogTicks", IDName: "log-ticks", Doc: "LogTicks is suitable for the Ticker field of an Axis,\nit returns tick marks suitable for a log-scale axis.", Fields: []types.Field{{Name: "Prec", Doc: "Prec specifies the precision of tick rendering\naccording to the documentation for strconv.FormatFloat."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.ConstantTicks", IDName: "constant-ticks", Doc: "ConstantTicks is suitable for the Ticker field of an Axis.\nThis function returns the given set of ticks."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.TimeTicks", IDName: "time-ticks", Doc: "TimeTicks is suitable for axes representing time values.", Fields: []types.Field{{Name: "Ticker", Doc: "Ticker is used to generate a set of ticks.\nIf nil, DefaultTicks will be used."}, {Name: "Format", Doc: "Format is the textual representation of the time value.\nIf empty, time.RFC3339 will be used"}, {Name: "Time", Doc: "Time takes a float32 value and converts it into a time.Time.\nIf nil, UTCUnixTime is used."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/plot.TickerFunc", IDName: "ticker-func", Doc: "TickerFunc is suitable for the Ticker field of an Axis.\nIt is an adapter which allows to quickly setup a Ticker using a function with an appropriate signature."})
