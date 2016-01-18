package ltry

// import (
//     "fmt"
// )

type Ltry_Set_Color struct {
    Ltrys []*Ltry_Color
}

func (ltry_list *Ltry_Set_Color)Pt() int {
    for _, lt :=range ltry_list.Ltrys {
        lt.Pt()
    }
    return 0
}


func (lts *Ltry_Set_Color)Len() int {
    return len(lts.Ltrys)
}

func (lts *Ltry_Set_Color)Less(i, j int) bool {
    return lts.Ltrys[i].index <= lts.Ltrys[j].index
}

func (lts *Ltry_Set_Color)Swap(i, j int) {
    lts.Ltrys[i].index, lts.Ltrys[j].index = lts.Ltrys[j].index,lts.Ltrys[i].index
    for k,_ := range lts.Ltrys[i].a {
        lts.Ltrys[i].a[k] ,lts.Ltrys[j].a[k] = lts.Ltrys[j].a[k], lts.Ltrys[i].a[k]
    }
    for k,_ := range lts.Ltrys[i].b {
        lts.Ltrys[i].b[k] ,lts.Ltrys[j].b[k] = lts.Ltrys[j].b[k], lts.Ltrys[i].b[k]
    }
    return
}

func (lts *Ltry_Set_Color)Append(lt Ltry) bool {
    switch lt.(type) {
    case *Ltry_Color:
        lts.Ltrys = append(lts.Ltrys, lt.(*Ltry_Color))
        return true
    default:
        return false
    }

    return false
}

func (lts *Ltry_Set_Color)GetLt(i int) (Ltry, bool) {

    // fmt.Printf("i:%d, len:%d\n", i, len(lts.Ltrys))
    if i <0 || i> len(lts.Ltrys){
        return nil, false
    }
    lt := lts.Ltrys[i]
    return lt,  true
}
