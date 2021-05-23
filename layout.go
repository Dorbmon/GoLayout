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
	return
}

func (z *Context) Calculate() {
	C.lay_run_context(&z.ctx)
}
func (z *Context) ReserveItemsCapacity(size int) {
	C.lay_reserve_items_capacity(&z.ctx, C.int(size))
}
func (z *Context) Reset() {
	C.lay_reset_context(&z.ctx)
}
func (z *Context) Destroy() {
	C.lay_destroy_context(&z.ctx)
}
