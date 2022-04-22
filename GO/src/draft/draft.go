// the telephone alphabet:
package main

import (
	"fmt"
	"sort"
)

var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	// unsorted: 
	// key: bravo, value:56
	// key: delta, value:87
	// key: echo, value:56
	// key: foxtrot, value:12
	// key: indio, value:87
	// key: juliet, value:65
	// key: kili, value:43
	// key: lima, value:98
	// key: alpha, value:34
	// key: charlie, value:23
	// key: golf, value:34
	// key: hotel, value:16
	
	// sorted: 
	// key: alpha, value: 34
	// key: bravo, value: 56
	// key: charlie, value: 23
	// key: delta, value: 87
	// key: echo, value: 56
	// key: foxtrot, value: 12
	// key: golf, value: 34
	// key: hotel, value: 16
	// key: indio, value: 87
	// key: juliet, value: 65
	// key: kili, value: 43
	// key: lima, value: 98
	fmt.Println("unsorted: ")
	for k,v := range barVal{
		fmt.Printf("key: %v, value:%v\n", k,v)
	}
	keys:=make([]string, len(barVal))
	i:=0
	for k := range barVal {	keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted: ")
	for _, k:= range keys{
		fmt.Printf("key: %v, value: %v\n", k, barVal[k])
	}
}