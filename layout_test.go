package Layout

import (
	"testing"
)

func TestAll(t *testing.T) {
	ctx := NewLayout()
	ctx.ReserveItemsCapacity(1024)
	root := NewLayItem(&ctx)
	root.SetSizeXY(1280, 720)
	root.SetContain(LayRow)
	master_list := NewLayItem(&ctx)
	root.Insert(&master_list)
	master_list.SetSizeXY(400, 0)
	master_list.SetBehave(LayVFill)
	master_list.SetContain(LayColumn)
	content_view := NewLayItem(&ctx)
	root.Insert(&content_view)
	content_view.SetBehave(LayHFill | LayVFill)
	ctx.Calculate()
	r := master_list.GetRect()
	if r.x1 != 0 || r.y1 != 0 || r.x2 != 400 || r.y2 != 720 {
		t.FailNow()
		return
	}
	r = content_view.GetRect()
	if r.x1 != 400 || r.y1 != 0 || r.x2 != 880 || r.y2 != 720 {
		t.FailNow()
		return
	}
}
