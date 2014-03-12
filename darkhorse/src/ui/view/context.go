package view

import (

)

////////////////////
// NOT USED - REMOVE !!!!!!!!!!!!!!!!!! <<<--------
////////////////////



//type Context struct {
//	surface *Surface
//	Width   float64
//	Height  float64
//	x, y    float64	
//}
//
////func (self *Context) Save() {
////	self.surface.save(self.context)
////}
////
////func (self *Context) Restore() {
////	self.surface.restore(self.context)
////}
//
//func (self *Context) PushGroup() {
//	self.surface.PopGroup()
//}
//
//func (self *Context) PushGroupWithContent(content Content) {
//	self.surface.PushGroupWithContent(content)
//}
//
//func (self *Context) PopGroup() *Pattern {
//	return self.surface.PopGroup()
//}
//
//func (self *Context) PopGroupToSource() {
//	self.surface.PopGroupToSource()
//}
//
//func (self *Context) SetOperator(operator Operator) {
//	self.surface.SetOperator(operator)
//}
//
//func (self *Context) SetSource(pattern *Pattern) {
//	self.surface.SetSource(pattern)
//}
//
//func (self *Context) SetSourceRGB(red, green, blue float64) {
//	self.surface.SetSourceRGB(red, green, blue)
//}
//
//func (self *Context) SetSourceRGBA(red, green, blue, alpha float64) {
//	self.surface.SetSourceRGBA(red, green, blue, alpha)
//}
//
//func (self *Context) SetSourceSurface(surface *Surface, x, y float64) {
//	self.surface.SetSourceSurface(surface, x + self.x, y + self.y)
//}
//
//func (self *Context) SetTolerance(tolerance float64) {
//	self.surface.SetTolerance(tolerance)
//}
//
//func (self *Context) SetAntialias(antialias Antialias) {
//	self.surface.SetAntialias(antialias)
//}
//
//func (self *Context) SetFillRule(fill_rule FillRule) {
//	self.surface.SetFillRule(fill_rule)
//}
//
//func (self *Context) SetLineWidth(width float64) {
//	self.surface.SetLineWidth(width)
//}
//
//func (self *Context) SetLineCap(line_cap LineCap) {
//	self.surface.SetLineCap(line_cap)
//}
//
//func (self *Context) SetLineJoin(line_join LineJoin) {
//	self.surface.SetLineJoin(line_join)
//}
//
//func (self *Context) SetDash(dashes []float64, num_dashes int, offset float64) {
//	self.surface.SetDash(dashes, num_dashes, offset)
//}
//
//func (self *Context) SetMiterLimit(limit float64) {
//	self.surface.SetMiterLimit(limit)
//}
//
//func (self *Context) Translate(tx, ty float64) {
//	self.surface.translate(self.context, C.double(tx), C.double(ty))
//}
//
//func (self *Context) Scale(sx, sy float64) {
//	self.surface.scale(self.context, C.double(sx), C.double(sy))
//}
//
//func (self *Context) Rotate(angle float64) {
//	self.surface.rotate(self.context, C.double(angle))
//}
//
//func (self *Context) Transform(matrix Matrix) {
//	self.surface.transform(self.context, matrix.cairo_matrix_t())
//}
//
//func (self *Context) SetMatrix(matrix Matrix) {
//	self.surface.set_matrix(self.context, matrix.cairo_matrix_t())
//}
//
//func (self *Context) IdentityMatrix() {
//	self.surface.identity_matrix(self.context)
//}
//
//func (self *Context) UserToDevice(x, y float64) (float64, float64) {
//	self.surface.user_to_device(self.context, (*C.double)(&x), (*C.double)(&y))
//	return x, y
//}
//
//func (self *Context) UserToDeviceDistance(dx, dy float64) (float64, float64) {
//	self.surface.user_to_device_distance(self.context, (*C.double)(&dx), (*C.double)(&dy))
//	return dx, dy
//}
//
//// path creation methods
//
//func (self *Context) NewPath() {
//	self.surface.new_path(self.context)
//}
//
//func (self *Context) MoveTo(x, y float64) {
//	self.surface.move_to(self.context, C.double(x), C.double(y))
//}
//
//func (self *Context) NewSubPath() {
//	self.surface.new_sub_path(self.context)
//}
//
//func (self *Context) LineTo(x, y float64) {
//	self.surface.line_to(self.context, C.double(x), C.double(y))
//}
//
//func (self *Context) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
//	self.surface.curve_to(self.context,
//		C.double(x1), C.double(y1),
//		C.double(x2), C.double(y2),
//		C.double(x3), C.double(y3))
//}
//
//func (self *Context) Arc(xc, yc, radius, angle1, angle2 float64) {
//	self.surface.arc(self.context,
//		C.double(xc), C.double(yc),
//		C.double(radius),
//		C.double(angle1), C.double(angle2))
//}
//
//func (self *Context) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
//	self.surface.arc_negative(self.context,
//		C.double(xc), C.double(yc),
//		C.double(radius),
//		C.double(angle1), C.double(angle2))
//}
//
//func (self *Context) RelMoveTo(dx, dy float64) {
//	self.surface.rel_move_to(self.context, C.double(dx), C.double(dy))
//}
//
//func (self *Context) RelLineTo(dx, dy float64) {
//	self.surface.rel_line_to(self.context, C.double(dx), C.double(dy))
//}
//
//func (self *Context) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
//	self.surface.rel_curve_to(self.context,
//		C.double(dx1), C.double(dy1),
//		C.double(dx2), C.double(dy2),
//		C.double(dx3), C.double(dy3))
//}
//
//func (self *Context) Rectangle(x, y, width, height float64) {
//	self.surface.rectangle(self.context,
//		C.double(x), C.double(y),
//		C.double(width), C.double(height))
//}
//
//func (self *Context) RoundedRectangle(x, y, width, height, radiusUL, radiusUR, radiusLR, radiusLL float64) {
//	degrees := math.Pi / 180.0;
//	self.NewSubPath();
//	self.Arc(x + radiusUL, y + radiusUL, radiusUL, 180 * degrees, 270 * degrees)
//	self.Arc(x + width - radiusUR, y + radiusUR, radiusUR, -90 * degrees, 0 * degrees)
//	self.Arc(x + width - radiusLR, y + height - radiusLR, radiusLR, 0 * degrees, 90 * degrees)
//	self.Arc(x + radiusLL, y + height - radiusLL, radiusLL, 90 * degrees, 180 * degrees)
//	
//	self.ClosePath()
//}
//
//func (self *Context) ClosePath() {
//	self.surface.close_path(self.context)
//}
//
//func (self *Context) PathExtents() (left, top, right, bottom float64) {
//	self.surface.path_extents(self.context,
//		(*C.double)(&left), (*C.double)(&top),
//		(*C.double)(&right), (*C.double)(&bottom))
//	return left, top, right, bottom
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Painting methods
//
//func (self *Context) Paint() {
//	self.surface.paint(self.context)
//}
//
//func (self *Context) PaintWithAlpha(alpha float64) {
//	self.surface.paint_with_alpha(self.context, C.double(alpha))
//}
//
//func (self *Context) Mask(pattern Pattern) {
//	self.surface.mask(self.context, pattern.pattern)
//}
//
//func (self *Context) MaskSurface(surface *Context, surface_x, surface_y float64) {
//	self.surface.mask_surface(self.context, surface.surface, C.double(surface_x), C.double(surface_y))
//}
//
//func (self *Context) Stroke() {
//	self.surface.stroke(self.context)
//}
//
//func (self *Context) StrokePreserve() {
//	self.surface.stroke_preserve(self.context)
//}
//
//func (self *Context) Fill() {
//	self.surface.fill(self.context)
//}
//
//func (self *Context) FillPreserve() {
//	self.surface.fill_preserve(self.context)
//}
//
//func (self *Context) CopyPage() {
//	self.surface.copy_page(self.context)
//}
//
//func (self *Context) ShowPage() {
//	self.surface.show_page(self.context)
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Insideness testing
//
//func (self *Context) InStroke(x, y float64) bool {
//	return self.surface.in_stroke(self.context, C.double(x), C.double(y)) != 0
//}
//
//func (self *Context) InFill(x, y float64) bool {
//	return self.surface.in_fill(self.context, C.double(x), C.double(y)) != 0
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Rectangular extents
//
//func (self *Context) StrokeExtents() (left, top, right, bottom float64) {
//	self.surface.stroke_extents(self.context,
//		(*C.double)(&left), (*C.double)(&top),
//		(*C.double)(&right), (*C.double)(&bottom))
//	return left, top, right, bottom
//}
//
//func (self *Context) FillExtents() (left, top, right, bottom float64) {
//	self.surface.fill_extents(self.context,
//		(*C.double)(&left), (*C.double)(&top),
//		(*C.double)(&right), (*C.double)(&bottom))
//	return left, top, right, bottom
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Clipping methods
//
//func (self *Context) ResetClip() {
//	self.surface.reset_clip(self.context)
//}
//
//func (self *Context) Clip() {
//	self.surface.clip(self.context)
//}
//
//func (self *Context) ClipPreserve() {
//	self.surface.clip_preserve(self.context)
//}
//
//func (self *Context) ClipExtents() (left, top, right, bottom float64) {
//	self.surface.clip_extents(self.context,
//		(*C.double)(&left), (*C.double)(&top),
//		(*C.double)(&right), (*C.double)(&bottom))
//	return left, top, right, bottom
//}
//
//func (self *Context) ClipRectangleList() ([]Rectangle, Status) {
//	list := self.surface.copy_clip_rectangle_list(self.context)
//	defer self.surface.rectangle_list_destroy(list)
//	rects := make([]Rectangle, int(list.num_rectangles))
//	C.memcpy(unsafe.Pointer(&rects[0]), unsafe.Pointer(list.rectangles), C.size_t(list.num_rectangles*8))
//	return rects, Status(list.status)
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Font/Text methods
//
//func (self *Context) SelectFontFace(name string, font_slant_t, font_weight_t int) {
//	s := C.CString(name)
//	self.surface.select_font_face(self.context, s, self.surface.font_slant_t(font_slant_t), self.surface.font_weight_t(font_weight_t))
//	C.free(unsafe.Pointer(s))
//}
//
//func (self *Context) SetFontSize(size float64) {
//	self.surface.set_font_size(self.context, C.double(size))
//}
//
//func (self *Context) SetFontMatrix(matrix Matrix) {
//	self.surface.set_font_matrix(self.context, matrix.cairo_matrix_t())
//}
//
//func (self *Context) SetFontOptions(fontOptions *FontOptions) {
//	panic("not implemented") // todo
//}
//
//func (self *Context) GetFontOptions() *FontOptions {
//	panic("not implemented") // todo
//	return nil
//}
//
//func (self *Context) SetFontFace(fontFace *FontFace) {
//	panic("not implemented") // todo
//}
//
//func (self *Context) GetFontFace() *FontFace {
//	panic("not implemented") // todo
//	return nil
//}
//
//func (self *Context) SetScaledFont(scaledFont *ScaledFont) {
//	panic("not implemented") // todo
//}
//
//func (self *Context) GetScaledFont() *ScaledFont {
//	panic("not implemented") // todo
//	return nil
//}
//
//func (self *Context) ShowText(text string) {
//	cs := C.CString(text)
//	self.surface.show_text(self.context, cs)
//	C.free(unsafe.Pointer(cs))
//}
//
//func (self *Context) ShowGlyphs(glyphs []Glyph) {
//	panic("not implemented") // todo
//}
//
//func (self *Context) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flags TextClusterFlag) {
//}
//
//func (self *Context) TextPath(text string) {
//	cs := C.CString(text)
//	self.surface.text_path(self.context, cs)
//	C.free(unsafe.Pointer(cs))
//}
//
//func (self *Context) GlyphPath(glyphs []Glyph) {
//	panic("not implemented") // todo
//}
//
//func (self *Context) TextExtents(text string) *TextExtents {
//	cte := self.surface.text_extents_t{}
//	cs := C.CString(text)
//	self.surface.text_extents(self.context, cs, &cte)
//	C.free(unsafe.Pointer(cs))
//	te := &TextExtents{
//		Xbearing: float64(cte.x_bearing),
//		Ybearing: float64(cte.y_bearing),
//		Width:    float64(cte.width),
//		Height:   float64(cte.height),
//		Xadvance: float64(cte.x_advance),
//		Yadvance: float64(cte.y_advance),
//	}
//	return te
//}
//
//func (self *Context) GlyphExtents(glyphs []Glyph) *TextExtents {
//	panic("not implemented") // todo
//	//self.surface.text_extents
//	return nil
//}
//
//func (self *Context) FontExtents() *FontExtents {
//	panic("not implemented") // todo
//	//self.surface.text_extents
//	return nil
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Error status queries
//
//func (self *Context) Status() Status {
//	return Status(self.surface.status(self.context))
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Backend device manipulation
//
/////////////////////////////////////////////////////////////////////////////////
//// Surface manipulation
//
//func (self *Context) CreateForRectangle(x, y, width, height float64) *Context {
//	return &Surface{
//		context: self.context,
//		surface: self.surface.surface_create_for_rectangle(self.surface,
//			C.double(x), C.double(y), C.double(width), C.double(height)),
//	}
//}
//
//func (self *Context) Finish() {
//	self.surface.surface_finish(self.surface)
//}
//
//func (self *Context) GetReferenceCount() int {
//	return int(self.surface.surface_get_reference_count(self.surface))
//}
//
//func (self *Context) GetStatus() Status {
//	return Status(self.surface.surface_status(self.surface))
//}
//
//func (self *Context) GetType() SurfaceType {
//	return SurfaceType(self.surface.surface_get_type(self.surface))
//}
//
//func (self *Context) GetContent() Content {
//	return Content(self.surface.surface_get_content(self.surface))
//}
//
//func (self *Context) WriteToPNG(filename string) {
//	cs := C.CString(filename)
//	self.surface.surface_write_to_png(self.surface, cs)
//	C.free(unsafe.Pointer(cs))
//}
//
//// Already implemented via context split context/surface?
//// func (self *Context) GetFontOptions() *FontOptions {
//// 	// todo
//// 	// self.surface.surface_get_font_options (cairo_surface_t      *Context,				cairo_font_options_t *options);
//// 	return nil
//// }
//
//func (self *Context) Flush() {
//	self.surface.surface_flush(self.surface)
//}
//
//// Tells cairo that drawing has been done to Surface 
//// using means other than cairo, and that cairo should 
//// reread any cached areas. Note that you must call 
//// flush() before doing such drawing.
//func (self *Context) MarkDirty() {
//	self.surface.surface_mark_dirty(self.surface)
//}
//
////  x (int) – X coordinate of dirty rectangle
////  y (int) – Y coordinate of dirty rectangle
////  width (int) – width of dirty rectangle
////  height (int) – height of dirty rectangle
////
//// Like mark_dirty(), but drawing has been done only 
//// to the specified rectangle, so that cairo can 
//// retain cached contents for other parts of the 
//// surface.
//// 
//// Any cached clip set on the Surface will be reset 
//// by this function, to make sure that future cairo 
//// calls have the clip set that they expect.
//func (self *Context) MarkDirtyRectangle(x, y, width, height int) {
//	self.surface.surface_mark_dirty_rectangle(self.surface,
//		C.int(x), C.int(y), C.int(width), C.int(height))
//}
//
//// x_offset (float) – the offset in the X direction, 
////                    in device units
//// y_offset (float) – the offset in the Y direction, 
////                    in device units
////
//// Sets an offset that is added to the device 
//// coordinates determined by the CTM when drawing 
//// to Surface. One use case for this function is 
//// when we want to create a Surface that redirects 
//// drawing for a portion of an onscreen surface to 
//// an offscreen surface in a way that is completely 
//// invisible to the user of the cairo API. Setting a
//// transformation via Context.translate() isn’t 
//// sufficient to do this, since functions like 
//// Context.device_to_user() will expose the hidden 
//// offset.
////
//// Note that the offset affects drawing to the surface 
//// as well as using the surface in a source pattern.
//func (self *Context) SetDeviceOffset(x, y float64) {
//	self.surface.surface_set_device_offset(self.surface, C.double(x), C.double(y))
//}
//
//func (self *Context) GetDeviceOffset() (x, y float64) {
//	self.surface.surface_get_device_offset(self.surface, (*C.double)(&x), (*C.double)(&y))
//	return x, y
//}
//
//func (self *Context) SetFallbackResolution(xPixelPerInch, yPixelPerInch float64) {
//	self.surface.surface_set_fallback_resolution(self.surface,
//		C.double(xPixelPerInch), C.double(yPixelPerInch))
//}
//
//func (self *Context) GetFallbackResolution() (xPixelPerInch, yPixelPerInch float64) {
//	self.surface.surface_get_fallback_resolution(self.surface,
//		(*C.double)(&xPixelPerInch), (*C.double)(&yPixelPerInch))
//	return xPixelPerInch, yPixelPerInch
//}
//
//// Already defined for context
//// func (self *Context) CopyPage() {
//// 	self.surface.surface_copy_page(self.surface)
//// }
//
//// func (self *Context) ShowPage() {
//// 	self.surface.surface_show_page(self.surface)
//// }
//
//func (self *Context) HasShowTextGlyphs() bool {
//	return self.surface.surface_has_show_text_glyphs(self.surface) != 0
//}
//
//// GetData returns a copy of the surfaces raw pixel data.
//// This method also calls Flush.
//func (self *Context) GetData() []byte {
//	self.Flush()
//	dataPtr := self.surface.image_surface_get_data(self.surface)
//	if dataPtr == nil {
//		panic("cairo.Surface.GetData(): can't access surface pixel data")
//	}
//	stride := self.surface.image_surface_get_stride(self.surface)
//	height := self.surface.image_surface_get_height(self.surface)
//	return C.GoBytes(unsafe.Pointer(dataPtr), stride*height)
//}
//
//// SetData sets the surfaces raw pixel data.
//// This method also calls Flush and MarkDirty.
//func (self *Context) SetData(data []byte) {
//	self.Flush()
//	dataPtr := unsafe.Pointer(self.surface.image_surface_get_data(self.surface))
//	if dataPtr == nil {
//		panic("cairo.Surface.SetData(): can't access surface pixel data")
//	}
//	stride := self.surface.image_surface_get_stride(self.surface)
//	height := self.surface.image_surface_get_height(self.surface)
//	if len(data) != int(stride*height) {
//		panic("cairo.Surface.SetData(): invalid data size")
//	}
//	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(stride*height))
//	self.MarkDirty()
//}
//
//func (self *Context) GetFormat() Format {
//	return Format(self.surface.image_surface_get_format(self.surface))
//}
//
//func (self *Context) GetWidth() int {
//	return int(self.surface.image_surface_get_width(self.surface))
//}
//
//func (self *Context) GetHeight() int {
//	return int(self.surface.image_surface_get_height(self.surface))
//}
//
//func (self *Context) GetStride() int {
//	return int(self.surface.image_surface_get_stride(self.surface))
//}
//
/////////////////////////////////////////////////////////////////////////////////
//// Pattern creation methods
//
/////////////////////////////////////////////////////////////////////////////////
//// image.Image methods
//
//func (self *Context) GetImage() image.Image {
//	width := self.GetWidth()
//	height := self.GetHeight()
//	stride := self.GetStride()
//	data := self.GetData()
//
//	switch self.GetFormat() {
//	case FORMAT_ARGB32:
//		return &extimage.BGRA{
//			Pix:    data,
//			Stride: stride,
//			Rect:   image.Rect(0, 0, width, height),
//		}
//
//	case FORMAT_RGB24:
//		return &extimage.BGRN{
//			Pix:    data,
//			Stride: stride,
//			Rect:   image.Rect(0, 0, width, height),
//		}
//
//	case FORMAT_A8:
//		return &image.Alpha{
//			Pix:    data,
//			Stride: stride,
//			Rect:   image.Rect(0, 0, width, height),
//		}
//
//	case FORMAT_A1:
//		panic("Unsuppored surface format cairo.FORMAT_A1")
//
//	case FORMAT_RGB16_565:
//		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")
//
//	case FORMAT_RGB30:
//		panic("Unsuppored surface format cairo.FORMAT_RGB30")
//
//	case FORMAT_INVALID:
//		panic("Invalid surface format")
//	}
//	panic("Unknown surface format")
//}
//
//// SetImage set the data from an image.Image with identical size.
//func (self *Context) SetImage(img image.Image) {
//	width := self.GetWidth()
//	height := self.GetHeight()
//	stride := self.GetStride()
//
//	switch self.GetFormat() {
//	case FORMAT_ARGB32:
//		if i, ok := img.(*extimage.BGRA); ok {
//			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
//				self.SetData(i.Pix)
//				return
//			}
//		}
//		surfImg := self.GetImage().(*extimage.BGRA)
//		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
//		self.SetData(surfImg.Pix)
//
//	case FORMAT_RGB24:
//		if i, ok := img.(*extimage.BGRN); ok {
//			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
//				self.SetData(i.Pix)
//				return
//			}
//		}
//		surfImg := self.GetImage().(*extimage.BGRN)
//		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
//		self.SetData(surfImg.Pix)
//
//	case FORMAT_A8:
//		if i, ok := img.(*image.Alpha); ok {
//			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
//				self.SetData(i.Pix)
//				return
//			}
//		}
//		surfImg := self.GetImage().(*image.Alpha)
//		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
//		self.SetData(surfImg.Pix)
//
//	case FORMAT_A1:
//		panic("Unsuppored surface format cairo.FORMAT_A1")
//
//	case FORMAT_RGB16_565:
//		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")
//
//	case FORMAT_RGB30:
//		panic("Unsuppored surface format cairo.FORMAT_RGB30")
//
//	case FORMAT_INVALID:
//		panic("Invalid surface format")
//	}
//	panic("Unknown surface format")
//}
//
//
//
//
//
