package entity

func init() {
	RegisterRuleset("lvsd", func() Ruleset { return NewLVSD() })
}

// LVSD is the "Light vs Dark" ruleset, the original SimstepLVSD behavior
// extracted from GridEntity. See docs/rules.md.
type LVSD struct {
	cutoff byte
}

// NewLVSD returns an LVSD ruleset with default cutoff.
func NewLVSD() *LVSD { return &LVSD{cutoff: 127} }

func (r *LVSD) Name() string { return "lvsd" }

func (r *LVSD) Init(g *GridEntity) {}

// Step performs one cycle of the center-distance intensity comparison sim.
// calculateInteraction runs twice per cell per tick — intentional perceived-speed hack.
func (r *LVSD) Step(g *GridEntity) {
	g.nticks++
	g.reload()
	clear(g.zone.zsum)
	rv := g.rng[0]
	for i := range g.Area {
		g.pxtozone(i)
		r.calculateInteraction(g, i, g.nticks^rv)
		r.calculateInteraction(g, i, g.nticks-rv)
	}
}

// CutoffUp raises the victory cutoff to see effects in real time.
func (r *LVSD) CutoffUp() {
	t := r.cutoff
	switch {
	case t < 148 && t > 108:
		t += 2
	case t >= 148 && t < 245:
		t += 8
	case t <= 108:
		t += 8
	}
	r.cutoff = t
}

// CutoffDown lowers the victory cutoff to see effects in real time.
func (r *LVSD) CutoffDown() {
	t := r.cutoff
	switch {
	case t < 148 && t > 108:
		t -= 2
	case t >= 148 && t < 254:
		t -= 8
	case t <= 108 && t > 8:
		t -= 8
	}
	r.cutoff = t
}

// CutoffIs returns the current cutoff.
func (r *LVSD) CutoffIs() byte { return r.cutoff }

type outcome = int

const (
	ineut outcome = iota
	ilose
	iwin
	istale
	ifriend
	imine
)

func (r *LVSD) calculateInteraction(g *GridEntity, i int, xor1 byte) {
	iR := i * 4
	dir := xor1 % 4
	var opp, oppR int
	switch dir {
	case 0:
		opp = sidewrap(i, 1, int(g.X))
	case 1:
		opp = sidewrap(i, -1, int(g.X))
	case 2:
		opp = wrap(i-int(g.X), g.Area)
	case 3:
		opp = wrap(i+int(g.X), g.Area)
	}
	oppR = opp * 4
	irng := g.getrng(i)
	iVal := bavg(g.Px[iR], g.Px[iR+1], g.Px[iR+2])
	oVal := bavg(g.Px[oppR], g.Px[oppR+1], g.Px[oppR+2])
	results := r.versus(irng, iVal, oVal)
	r.exec1v1(g, results[0], i, opp, iVal > 128)
}

func (r *LVSD) versus(rng byte, iClr byte, versus ...byte) []outcome {
	wout := make([]outcome, len(versus))
	if iClr == 127 || iClr == 128 {
		return wout
	}
	alignment := iClr > 128
	for i, v := range versus {
		if v == 127 || v == 128 {
			wout[i] = imine
		} else if (v > 128) == alignment {
			wout[i] = ifriend
		} else {
			cutoff := (int(r.cutoff) + int(rng)) / 2
			rval := battlemc(iClr, v, byte(cutoff))
			switch {
			case rval > 0:
				wout[i] = iwin
			case rval < 0:
				wout[i] = ilose
			case rval == 0:
				wout[i] = istale
			}
		}
	}
	return wout
}

// battlemc: sign gives win/lose, magnitude gives by how much.
func battlemc(mainchar, enemy, rng byte) int {
	victoryLine := mainchar + enemy - 128
	mcWin := int(rng) - int(victoryLine)
	if mainchar < 127 {
		return -mcWin
	}
	return mcWin
}

func (r *LVSD) exec1v1(g *GridEntity, o outcome, ipx, epx int, align bool) {
	i := ipx * 4
	e := epx * 4
	_ = align
	switch o {
	case ilose:
		sliceToward(g.Px[e:e+3], g.Px[i:i+3], 200)
	case iwin:
		sliceToward(g.Px[i:i+3], g.Px[e:e+3], 200)
	case ifriend:
		r.interactFriend(g, i, e)
	case imine:
		r.interactMine(g, i, e)
	case istale:
		r.interactStalemate(g, i, e)
	}
}

func (r *LVSD) interactMine(g *GridEntity, i, m int) {
	if g.getrng(i)%10 < 3 {
		ip := g.Px[i : i+4]
		ival := bavg(ip[0], ip[1], ip[2])
		light := ival > 128
		mp := g.Px[m : m+4]
		s := pxisort(mp[:3])
		if light {
			mineLV1 := mp[s[2]] - mp[s[1]]
			mineLV2 := mp[s[1]] - mp[s[0]]
			switch {
			case mineLV1 > 15:
				mp[s[2]], ip[s[2]] = bmov(mp[s[2]], ip[s[2]], 16)
			case mineLV2 > 3:
				mp[s[2]], ip[s[2]] = bmov(mp[s[2]], ip[s[2]], 4)
				mp[s[1]], ip[s[1]] = bmov(mp[s[1]], ip[s[1]], 4)
			default:
				bsladd(ip[:3], 2)
				bslsub(mp[:3], 2)
			}
		} else {
			mineLV1 := (255 - mp[s[0]]) - (255 - mp[s[1]])
			mineLV2 := (255 - mp[s[1]]) - (255 - mp[s[2]])
			switch {
			case mineLV1 > 15:
				ip[s[0]], mp[s[0]] = bmov(ip[s[0]], mp[s[0]], 16)
			case mineLV2 > 3:
				ip[s[0]], mp[s[0]] = bmov(ip[s[0]], mp[s[0]], 4)
				ip[s[1]], mp[s[1]] = bmov(ip[s[1]], mp[s[1]], 4)
			default:
				bsladd(mp[:3], 2)
				bslsub(ip[:3], 2)
			}
		}
	}
}

func (r *LVSD) interactStalemate(g *GridEntity, i, e int) {
	ipx := g.Px[i : i+4]
	epx := g.Px[e : e+4]
	srng := int(g.getrng(i+2)) + int(g.getrng(i+1))
	xtrarng := g.getrng(i + 666)
	iavg := bavg(ipx[0], ipx[1], ipx[2])
	if xtrarng > 252 {
		for n := range ipx {
			ipx[(srng*e+n)%3] ^= epx[(srng*i+n)%3]
			epx[(srng*i+n)%3] ^= ipx[(srng*e+2+n)%3]
		}
	} else if iavg > 128 {
		bsladd(ipx[:3], 16)
		bslsub(epx[:3], 16)
	} else if iavg < 127 {
		bsladd(epx[:3], 16)
		bslsub(ipx[:3], 16)
	}
}

func (r *LVSD) interactFriend(g *GridEntity, i, e int) {
	ip := g.Px[i : i+3]
	ep := g.Px[e : e+3]
	ipavg := acdbavg(ip...)
	epavg := acdbavg(ep...)
	alignment := bavg(ip[0], ip[1], ip[2]) > 128
	if ipavg > epavg {
		sliceToward(ip, ep, 64)
	} else if epavg > ipavg {
		sliceToward(ep, ip, 64)
	} else {
		si := pxisort(ip)
		se := pxisort(ep)
		if alignment {
			if ip[si[2]] > ep[se[2]] {
				temp := ep[si[2]]
				ep[si[2]] = ep[se[2]]
				ep[se[2]] = temp
			} else {
				temp := ip[se[2]]
				ip[se[2]] = ip[si[2]]
				ip[si[2]] = temp
			}
			bsladd(ip, 6)
			bsladd(ep, 6)
		} else {
			if ip[si[0]] < ip[se[0]] {
				temp := ep[si[0]]
				ep[si[0]] = ep[se[0]]
				ep[se[0]] = temp
			} else {
				temp := ip[se[0]]
				ip[se[0]] = ip[si[0]]
				ip[si[0]] = temp
			}
			bslsub(ip, 6)
			bslsub(ep, 6)
		}
	}
}

var (
	overlayRed  = []byte{240, 160, 170, 90}
	overlayBlue = []byte{30, 30, 80, 90}
	overlayMid  = []byte{90, 140, 90, 120}
)

// Overlay produces the LVSD debug overlay: red for light, blue for dark, mid for neutral.
func (r *LVSD) Overlay(g *GridEntity, mode int) []byte {
	_ = mode
	overlay := make([]byte, len(g.Px))
	var rr int
	for i := range g.Area {
		rr = i * 4
		ba := bavg(g.Px[rr], g.Px[rr+1], g.Px[rr+2])
		switch {
		case ba > 128:
			mto(g.Px[rr:rr+4], overlayRed, overlay[rr:rr+4])
		case ba < 127:
			mto(g.Px[rr:rr+4], overlayBlue, overlay[rr:rr+4])
		case ba == 127 || ba == 128:
			mto(g.Px[rr:rr+4], overlayMid, overlay[rr:rr+4])
		}
	}
	return overlay
}
