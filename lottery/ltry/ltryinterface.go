package ltry

// import (
//     "sort"
// )

//Ltry base interface of the lottery
type Ltry interface {
	Getfromstr(line string) error
	Pt() int
	Str() string
	Check(baselt Ltry) (int, int)
	CheckM(baselt Ltry) int

	AddSelf() bool
	BeginSelf() bool
	IsMaxSelf() bool
}

//LSet sets of the Ltry
type LSet interface {
	Pt() int
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(lt Ltry) bool

	GetLt(i int) (Ltry, bool)
}
