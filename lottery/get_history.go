package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    // "strings"
    // "strconv"
    "entertainment/lottery/ltry"
    "sort"
)


func main() {
    filename := "./test_data.txt"
    ltry_lst := new(ltry.Ltry_Set_Color)

    err := get_history_from_file(filename, ltry_lst)
    if err!= nil {
        fmt.Printf("[%s] read err[%s]\n", filename, err.Error())
    }

    sort.Sort(ltry_lst)
    ltry_lst.Pt()

    lt, ok := ltry_lst.GetLt(0)
    if !ok {
        fmt.Printf("get [%d] data err\n", 0)
        return
    }
    lt2,ok2 := ltry_lst.GetLt(1)
    if !ok2 {
        fmt.Printf("get [%d] data err\n", 1)
        return
    }

    a,b := lt2.Check(lt)
    fmt.Printf("a:%d, b:%d\n", a,b)

    p := lt2.CheckM(lt)
    fmt.Printf("p:%d\n", p)

    return

}


func get_history_from_file(filename string, ltry_list ltry.Ltry_Set) error  {

    f, err := os.Open(filename)
    if err !=nil {
        fmt.Printf("[%s] open err[%s]\n", filename, err.Error())
        return err
    }

    defer f.Close()

    buff := bufio.NewReader(f)

    for {
        line, err := buff.ReadString('\n')
        if err !=nil || io.EOF == err {
            break
        }

        lt := new(ltry.Ltry_Color)
        err = lt.Get_from_str(line)
        if err!= nil {
            fmt.Printf("Error [%s] get from str [%s]\n", line, err.Error())
            continue
        }
        ltry_list.Append(lt)
        //ltry_list.Ltrys = append(ltry_list.Ltrys, *lt)
    }

    return nil

}
