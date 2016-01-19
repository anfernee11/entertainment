package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "strings"
	// "strconv"
	"entertainment/lottery/ltry"
	"sort"
)

func main() {
	filename := "./test_data.txt"
	ltrylst := new(ltry.LSetColor)

	err := gethistoryfromfile(filename, ltrylst)
	if err != nil {
		fmt.Printf("[%s] read err[%s]\n", filename, err.Error())
	}

	sort.Sort(ltrylst)
	// ltrylst.Pt()

	pValues := checkSets(ltrylst, ltrylst)
	for k, v := range pValues {
		fmt.Printf("idx:[%4d], Value:[%7d]\n", k, v-5000000)
	}

	return

}

func gethistoryfromfile(filename string, ltrylist ltry.LSet) error {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("[%s] open err[%s]\n", filename, err.Error())
		return err
	}

	defer f.Close()

	buff := bufio.NewReader(f)

	for {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		lt := new(ltry.LtryColor)
		err = lt.Getfromstr(line)
		if err != nil {
			fmt.Printf("Error [%s] get from str [%s]\n", line, err.Error())
			continue
		}
		ltrylist.Append(lt)
		//ltrylist.Ltrys = append(ltry_list.Ltrys, *lt)
	}

	return nil

}

func checkSets(lts ltry.LSet, ltsbase ltry.LSet) (pValue []int) {
	pValue = make([]int, lts.Len())

	for idx := 0; idx < lts.Len(); idx++ {
		pv := 0
		lt, err := lts.GetLt(idx)
		if !err {
			continue
		}
		for idx2 := 0; idx2 < ltsbase.Len(); idx2++ {
			ltbase, err2 := ltsbase.GetLt(idx2)
			if !err2 {
				continue
			}
			pv += lt.CheckM(ltbase)
		}
		pValue[idx] = pv
	}

	return
}
