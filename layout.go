package Layout

/*
#define LAY_IMPLEMENTATION
#include "layout.h"
*/
import "C"

type Context struct {
	ctx C.lay_context
}

func NewLayout() (ctx Context) {
	C.lay_init_context(&ctx.ctx)
	C.lay_reserve_items_capacity(&ctx.ctx, C.uint(1024))
	return
}

func (z *Context) Calculate() {
	C.lay_run_context(&z.ctx)
}
func (z *Context) ReserveItemsCapacity(size uint) {
	C.lay_reserve_items_capacity(&z.ctx,C.uint(size))
}