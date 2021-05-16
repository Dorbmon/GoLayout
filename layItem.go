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
type Rect struct {
	x1,y1,x2,y2 int
}
//GetRect get the result of the calculation
func (z *LayItem) GetRect() (ret Rect){
	res := C.GetRect(&z.ctx.ctx, z.layId)
	ret.x1 = int(res.x1)
	ret.y1 = int(res.y1)
	ret.x2 = int(res.x2)
	ret.y2 = int(res.y2)
	return
}
func (z *LayItem) SetContain(behave C.uint) {
	C.lay_set_contain(&z.ctx.ctx, z.layId, behave)
}