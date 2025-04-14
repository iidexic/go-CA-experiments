package entity

/*-Replaced by calculateInteraction.
func (grid *GridEntity) processInteraction(i int) {
	irng := grid.getrng(i)
	up := wrap(i-int(grid.X), grid.Area)
	lft := sidewrap(i, -1, int(grid.X))
	iR := i * 4
	upR := up * 4
	lftR := lft * 4

	ival := bavg(grid.Px[iR : iR+3]...) // averaged value of pixel colors
	uval := bavg(grid.Px[upR : upR+3]...)
	lval := bavg(grid.Px[lftR : lftR+3]...)

	results := versusLVSD(irng, ival, uval, lval) // (currently) return slice of lightWin bools.
	standinrng := uval ^ lval
	if standinrng > 127 {
		grid.exec1v1(results[1], i, lft, ival > 128)
		grid.exec1v1(results[0], i, up, ival > 128)
	} else {
		grid.exec1v1(results[0], i, up, ival > 128)
		grid.exec1v1(results[1], i, lft, ival > 128)
	}
}
*/
