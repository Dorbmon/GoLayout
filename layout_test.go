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
	if r.X1 != 0 || r.Y1 != 0 || r.X2 != 400 || r.Y2 != 720 {
		t.FailNow()
		return
	}
	r = content_view.GetRect()
	if r.X1 != 400 || r.Y1 != 0 || r.X2 != 1280 || r.Y2 != 720 {
		t.FailNow()
		return
	}
}
