package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	// "strings"
	"entertainment/lottery/ltry"
	"sort"
)

func main() {

	fmt.Printf("Curent Path:[%s]\n\n", getCurrPath())

	filename := "./test_data.txt"
	ltrylst := new(ltry.LSetColor)

	checkBeg := 4000
	checkEnd := 10000
	var err error

	argNum := len(os.Args)
	if argNum >= 3 {
		checkBeg, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("para err arg1 [%s]\n", os.Args[1])
			return
		}
		checkEnd, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("para err arg2 [%s]\n", os.Args[2])
			return
		}
	}

	err = gethistoryfromfile(filename, ltrylst)
	if err != nil {
		fmt.Printf("[%s] read err[%s]\n", filename, err.Error())
	}

	sort.Sort(ltrylst)
	// ltrylst.Pt()

	// pValues := checkSets(ltrylst, ltrylst)
	// for k, v := range pValues {
	// 	fmt.Printf("idx:[%4d], Value:[%7d]\n", k, v-5000000)
	// }

	lt := new(ltry.LtryColor)
	lt.BeginSelf()
	totalCount := 0

	for i := 0; i < 100000000; i++ {
		pv := checkOneLt(lt, ltrylst)

		if pv <= checkEnd && pv >= checkBeg {
			fmt.Printf("%s, PV is [%d]\n", lt.Str(), pv)
			totalCount++
		}

		if !lt.AddSelf() {
			fmt.Printf("It's scaned all data already!!, Total Count [%d]\n", totalCount)
			return
		}
	}

	return

}

func checkOneLt(lt ltry.Ltry, ltsBase ltry.LSet) (pv int) {
	for idx2 := 0; idx2 < ltsBase.Len(); idx2++ {
		ltbase, err2 := ltsBase.GetLt(idx2)
		if !err2 {
			continue
		}
		pv += lt.CheckM(ltbase)
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

func getCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}
