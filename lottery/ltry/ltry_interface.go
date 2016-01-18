package ltry

// import (
//     "sort"
// )

type Ltry interface {
    Get_from_str(line string) error
    Pt() int
    Check(baselt Ltry) (int, int)
    CheckM(baselt Ltry) int
}

type Ltry_Set interface {
    Pt() int
    Len() int
    Less (i, j int) bool
    Swap(i, j int)
    Append(lt Ltry) bool

    GetLt(i int) (Ltry,bool)
}
