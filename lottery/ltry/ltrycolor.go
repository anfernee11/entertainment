package ltry

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	LTRY_P1_NUM int    = 6
	LTRY_P2_NUM int    = 1
	SPLIT_TAG   string = "\t"
	P1_MAX      int    = 33
	P2_MAX      int    = 16
)

func getPickM(a, b int) int {

	k := a*100 + b
	switch k {
	case 400:
		return 10
	case 500:
		return 200
	case 600:
		return 100000
	case 1:
		return 5
	case 101:
		return 5
	case 201:
		return 5
	case 301:
		return 10
	case 401:
		return 200
	case 501:
		return 3000
	case 601:
		return 5000000
	default:
		return 0
	}
	//return 0
}

/*LtryColor : color ball struct*/
type LtryColor struct {
	index int
	des   string
	a     [LTRY_P1_NUM]int
	b     [LTRY_P2_NUM]int
}

/*Pt : print ltry*/
func (ltry *LtryColor) Pt() int {
	fmt.Printf("[%8d] [", ltry.index)
	for i := 0; i < LTRY_P1_NUM; i++ {
		fmt.Printf("%2d ", ltry.a[i])
	}
	fmt.Printf("] [")
	for i := 0; i < LTRY_P2_NUM; i++ {
		fmt.Printf("%2d ", ltry.b[i])
	}
	fmt.Printf("]\n")
	return 0
}

/*Str : form string for pt*/
func (ltry *LtryColor) Str() (s string) {
	s += "[ " + strconv.Itoa(ltry.index) + " ][ "
	for _, v := range ltry.a {
		s += strconv.Itoa(v) + " "
	}
	s += "][ "
	for _, v := range ltry.b {
		s += strconv.Itoa(v) + " "
	}
	s += " ]"
	return
}

/*Getfromstr : get lottey data from*/
func (ltry *LtryColor) Getfromstr(line string) error {
	bline := strings.TrimSpace(line)
	slist := strings.Split(bline, SPLIT_TAG)
	var e error
	for k, v := range slist {
		if k == 0 {
			ltry.index, e = strconv.Atoi(v)
			if e != nil {
				fmt.Printf("Error[%s]\n", e.Error())
				return e
			}
		} else if k <= LTRY_P1_NUM {
			ltry.a[k-1], e = strconv.Atoi(v)
			if e != nil {
				fmt.Printf("Error[%s]\n", e.Error())
				return e
			}
		} else if k <= LTRY_P1_NUM+LTRY_P2_NUM {
			ltry.b[k-1-LTRY_P1_NUM], e = strconv.Atoi(v)
			if e != nil {
				fmt.Printf("Error[%s]\n", e.Error())
				return e
			}
		} else {
			break
		}
	}
	return e

}

/*Check : cal shot number*/
func (lt *LtryColor) Check(baselt Ltry) (int, int) {
	a, b := 0, 0
	switch baselt.(type) {
	case *LtryColor:
		blt, _ := baselt.(*LtryColor)
		// blt.Pt()
		// lt.Pt()
		for _, v := range lt.a {
			for _, vv := range blt.a {
				// fmt.Printf("v:%d, vv:%d\n",v,vv)
				if v == vv {
					a++
				}
			}
		}
		for _, v := range lt.b {
			for _, vv := range blt.b {
				if v == vv {
					b++
				}
			}
		}
		return a, b
	default:
		return 0, 0
	}
	return 0, 0
}

/*CheckM : cal shot value*/
func (lt *LtryColor) CheckM(baselt Ltry) int {
	a, b := lt.Check(baselt)
	return getPickM(a, b)
}

/*AddSelf : ball data add one*/
func (lt *LtryColor) AddSelf() bool {

	if lt.IsMaxSelf() {
		return false
	}
	for i := LTRY_P2_NUM - 1; i >= 0; i-- {
		if lt.b[i] < P2_MAX-i {
			lt.b[i]++
			for j := i + 1; j < LTRY_P2_NUM; j++ {
				lt.b[j] = lt.b[j-1] + 1
			}

			return true
		}
	}
	for i := 0; i < LTRY_P2_NUM; i++ {
		lt.b[i] = i + 1
	}

	for i := LTRY_P1_NUM - 1; i >= 0; i-- {
		if lt.a[i] < P1_MAX-i {
			lt.a[i]++
			for j := i + 1; j < LTRY_P1_NUM; j++ {
				lt.a[j] = lt.a[j-1] + 1
			}
			return true
		}
	}

	return false
}

/*BeginSelf : Init ball data*/
func (lt *LtryColor) BeginSelf() bool {
	for i := 0; i < LTRY_P1_NUM; i++ {
		lt.a[i] = i + 1
	}
	for i := 0; i < LTRY_P2_NUM; i++ {
		lt.b[i] = i + 1
	}

	return true
}

/*IsMaxSelf : check if is the max number*/
func (lt *LtryColor) IsMaxSelf() bool {
	for i := 0; i < LTRY_P1_NUM; i++ {
		if lt.a[i] != P1_MAX-LTRY_P1_NUM+1+i {
			return false
		}
	}

	for i := 0; i < LTRY_P2_NUM; i++ {
		if lt.b[i] != P2_MAX-LTRY_P2_NUM+1+i {
			return false
		}
	}
	return true
}
