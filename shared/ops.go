package shared

type OpKind byte

const (
	// OpPlus Increments the value of the current cell by 1.
	OpPlus OpKind = '+'

	// OpMinus Decrements the value of the current cell by 1.
	OpMinus OpKind = '-'

	// OpLeftArrow Moves the pointer to the previous cell (to the left of the current cell).
	OpLeftArrow OpKind = '<'

	// OpRightArrow Moves the pointer to the next cell (to the right of the current cell).
	OpRightArrow OpKind = '>'

	// OpDot Writes (outputs) the value of the current cell.
	OpDot OpKind = '.'

	// OpLeftBracket Jumps to the matching ] instruction if the value of the current cell is zero.
	OpLeftBracket OpKind = '['

	// OpRightBracket Jumps to the matching [ instruction if the value of the current cell is not zero.
	OpRightBracket OpKind = ']'

	// OpComma Reads the next byte from an input stream and replaces the value of the current cell with that new value.
	OpComma = ','
)
