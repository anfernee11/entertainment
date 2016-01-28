package ltry

// import (
//     "fmt"
// )

//LSetColor Ltrys' set
type LSetColor struct {
	Ltrys []*LtryColor
}

//Pt Print the LSetCoor
func (lts *LSetColor) Pt() int {
	for _, lt := range lts.Ltrys {
		lt.Pt()
	}
	return 0
}

func (lts *LSetColor) Len() int {
	return len(lts.Ltrys)
}

func (lts *LSetColor) Less(i, j int) bool {
	return lts.Ltrys[i].index <= lts.Ltrys[j].index
}

func (lts *LSetColor) Swap(i, j int) {
	lts.Ltrys[i].index, lts.Ltrys[j].index = lts.Ltrys[j].index, lts.Ltrys[i].index
	for k := range lts.Ltrys[i].a {
		lts.Ltrys[i].a[k], lts.Ltrys[j].a[k] = lts.Ltrys[j].a[k], lts.Ltrys[i].a[k]
	}
	for k := range lts.Ltrys[i].b {
		lts.Ltrys[i].b[k], lts.Ltrys[j].b[k] = lts.Ltrys[j].b[k], lts.Ltrys[i].b[k]
	}
	return
}

//Append add a ltry into the set
func (lts *LSetColor) Append(lt Ltry) bool {
	switch lt.(type) {
	case *LtryColor:
		lts.Ltrys = append(lts.Ltrys, lt.(*LtryColor))
		return true
	default:
		return false
	}

	//return false
}

//GetLt get one of the Ltry by index
func (lts *LSetColor) GetLt(i int) (Ltry, bool) {

	// fmt.Printf("i:%d, len:%d\n", i, len(lts.Ltrys))
	if i < 0 || i > len(lts.Ltrys) {
		return nil, false
	}
	lt := lts.Ltrys[i]
	return lt, true
}
