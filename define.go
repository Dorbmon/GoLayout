package Layout

/*
#include "layout.h"
*/
import "C"

const (
	// LayLeft anchor to left item or left side of parent
	LayLeft = 0x020
	// LayTop anchor to top item or top side of parent
	LayTop = 0x040
	// LayRight anchor to right item or right side of parent
	LayRight = 0x080
	// LayBottom anchor to bottom item or bottom side of parent
	LayBottom = 0x100
	// LayHFILL anchor to both left and right item or parent borders
	LayHFill = 0x0a0
	// LayVFill anchor to both top and bottom item or parent borders
	LayVFill = 0x140
	// LayHCenter center horizontally with left margin as offset
	LayHCenter = 0x000
	// LayVCenter center vertically with top margin as offset
	LayVCenter = 0x000
	// LayCenter center in both directions with left/top margin as offset
	LayCenter = 0x000
	// LayFill anchor to all four directions
	LayFill  = 0x1e0
	// LayBreak Drawing routines can read this via item pointers as needed after
	// performing layout calculations.
	LayBreak = 0x200
)
const (
	// LayRow left to right
	LayRow = 0x002
	// LayColumn top to bottom
	LayColumn = 0x003
	// LayLayout free layout
	LayLayout = 0x000
	// LayFlex flex model
	LayFlex = 0x002

	// LayNoWrap single-line
	LayNoWrap = 0x000
	// LayWrap multi-line wrap left to right
	LayWrap = 0x004

	// LayStart justify-content (start end center space-between)
	// at start of row/column
	LayStart = 0x008
	// LayMiddle at center of row/column
	LayMiddle = 0x000
	// LayEnd at end of row/column
	LayEnd = 0x010
	// LayJustify insert spacing to stretch across whole row/column
	LayJustify = 0x018
)
