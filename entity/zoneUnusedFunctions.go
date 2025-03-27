package entity

import "slices"

func zoneSetup(x, y, width, height int) zones {

	z := zones{
		numX: x, numY: y, count: x * y,
	}
	z.zsum1 = make([]int, z.count)
	z.zrange = make([]int, 0, z.count*2) //{x0,y0,x1,y1...xn,yn}
	xdiv := width / z.numX
	xextra := width % z.numX
	var xadd, yadd []int
	if xextra > 0 {
		xadd = extraPriority(xextra, z.numX)
	}
	ydiv := height / z.numY
	yextra := height % z.numY
	if yextra > 0 {
		yadd = extraPriority(yextra, z.numY)
	}
	xi := 0
	yi := 0
	for n := range z.count { //for each zone

		posX := n % z.numX //zone position x = remainder of zone mod after columns removed
		x := posX * xdiv   //grid x coordinate of first Px in zone
		if xi < len(xadd) && n == xadd[xi] {
			x++
			xi++
		}
		posY := n / z.numX //zone pos y == div by row count
		y := posY * ydiv
		if y < len(yadd) && n == yadd[yi] {
			y++
			yi++
		}
		z.zrange = append(z.zrange, x, y)
	}
	return z
}

// extraPriority prioritizes zone placement of pixel count division remainder
// Messy BUT I think it will work. All using variations on same formula, can be simplified
// -Update: Using a much simpler method inline instead
func extraPriority(extra, tot int) []int {
	if extra >= tot {
		panic("extra>=tot")
	}
	priority := make([]int, extra)
	if tot%2 == 0 { //num zones even
		if extra%2 == 1 {
			extra--
		}
		//*justin case check
		for i := 0; extra > 0; i++ {
			priority[2*i] = (tot / 2) - i
			priority[(2*i)+1] = (tot / 2) + i + 1
			extra -= 2
		}
		//> examples: nz = 10 | nz = 16 > xtra=2: 10/2,+1 | 16/2,+1 >xtra=4: /2,/2+1,/2-1,/2+2,/2-2,etc.
	} else if extra%2 == 0 { //num zones odd, extra even
		st := tot / 2
		for i := 0; extra > 0; i += 2 {
			priority[i] = st
			priority[i+1] = st + 2 + (2 * i) //5,7,4,8,3,9 = a,+2,a-1,+4,a-2,+6
			st--
			extra -= 2
		}
	} else { //num zones odd, extra odd
		priority[0] = (tot / 2) + 1
		if extra > 1 {
			for i := 2; extra > 1; i += 2 {
				priority[i-1] = (tot - i) / 2
				priority[i] = (tot + i) / 2
				extra -= 2
			}
		}
	}
	slices.Sort(priority)
	return priority

}
