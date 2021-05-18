package Layout

/*
#include "layout.h"

struct Rect {
	int x1,y1,x2,y2;
};
struct Rect GetRect(lay_context* ctx,lay_id id) {
	lay_vec4 result = lay_get_rect(ctx,id);
	struct Rect ret;
	ret.x1 = result[0];
	ret.y1 = result[1];
	ret.x2 = result[2];
	ret.y2 = result[3];
	return ret;
}
struct WH {
	int16_t w,h;
};
struct WH GetSize(lay_context* ctx,lay_id id) {
	struct WH ret;
	lay_vec2 res = lay_get_size(ctx,id);
	ret.w = res[0];
	ret.h = res[1];
	return ret;
}
struct WH GetSizeXY(lay_context* ctx,lay_id id) {
	struct WH ret;
	lay_get_size_xy(ctx,id,&ret.w,&ret.h);
	return ret;
}
struct Margins {
	int16_t l,t,r,b;
};
struct Margins GetMargins(lay_context* ctx,lay_id id) {
	struct Margins ret;
	lay_get_margins_ltrb(ctx,id,&ret.l,&ret.t,&ret.r,&ret.b);
	return ret;
}
void SetMargins(lay_context *ctx,lay_id id, lay_scalar l, lay_scalar t, lay_scalar r, lay_scalar b) {
	lay_set_margins_ltrb(ctx,id,l,t,r,b);
	return;
}
*/
import "C"

type LayItem struct {
	layId C.lay_id
	ctx   *Context
}

func NewLayItem(ctx *Context) (ret LayItem) {
	ret.layId = C.lay_item(&ctx.ctx)
	ret.ctx = ctx
	return
}
func (z *LayItem) SetSizeXY(x, y int32) {
	C.lay_set_size_xy(&z.ctx.ctx, z.layId, C.short(x), C.short(y))
}
func (z *LayItem) SetBehave(behave C.uint) {
	C.lay_set_behave(&z.ctx.ctx, z.layId, behave)
}
func (z *LayItem) Insert(obj *LayItem) {
	C.lay_insert(&obj.ctx.ctx, z.layId, obj.layId)
}

// Calculate only calculate the child of the correct layItem
func (z *LayItem) Calculate() {
	C.lay_run_item(&z.ctx.ctx, z.layId)
}
func (z *LayItem) Append(other *LayItem) {
	C.lay_append(&z.ctx.ctx, z.layId, other.layId)
}
func (z *LayItem) Push(other *LayItem) {
	C.lay_push(&z.ctx.ctx, z.layId, other.layId)
}
func (z *LayItem) GetSize() (w, h int32) {
	res := C.GetSize(&z.ctx.ctx, z.layId)
	return int32(res.w), int32(res.h)
}
func (z *LayItem) GetSizeXY() (w, h int32) {
	res := C.GetSizeXY(&z.ctx.ctx, z.layId)
	return int32(res.w), int32(res.h)
}
func (z *LayItem) GetMargins() (l, t, r, b int16) {
	res := C.GetMargins(&z.ctx.ctx, z.layId)
	return int16(res.l), int16(res.t), int16(res.r), int16(res.b)
}
func (z *LayItem) SetMargins(l, t, r, b int16) {
	C.SetMargins(&z.ctx.ctx, z.layId, C.short(l), C.short(t), C.short(r), C.short(b))
}

type Rect struct {
	X1, Y1, X2, Y2 int
}

//GetRect get the result of the calculation
func (z *LayItem) GetRect() (ret Rect) {
	res := C.GetRect(&z.ctx.ctx, z.layId)
	ret.X1 = int(res.x1)
	ret.Y1 = int(res.y1)
	ret.X2 = int(res.x2)
	ret.Y2 = int(res.y2)
	return
}
func (z *LayItem) SetContain(behave C.uint) {
	C.lay_set_contain(&z.ctx.ctx, z.layId, behave)
}
