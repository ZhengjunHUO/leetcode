package main

import (
	"fmt"
	"strings"
	"strconv"
	zstr "github.com/ZhengjunHUO/goutil/strings"
)

func find_mapping(ref_seq string, pdb_seq [][2]string) {
	if len(ref_seq) == 0 {
		return
	}

	var rslt string

	from_idx, _ := strconv.Atoi(strings.Split(pdb_seq[0][0], ".")[1])
	prev_idx := from_idx
	curr_pattern := pdb_seq[0][1]

	start_idx := []int{from_idx}
	patterns := []string{}

	for i := 1; i < len(pdb_seq); i++ {
		curr_idx, _ := strconv.Atoi(strings.Split(pdb_seq[i][0], ".")[1])
		if prev_idx == curr_idx || prev_idx == curr_idx - 1 {
			curr_pattern += pdb_seq[i][1]
			prev_idx = curr_idx
		}else{
			patterns = append(patterns, curr_pattern)
			curr_pattern, from_idx, prev_idx = pdb_seq[i][1], curr_idx, curr_idx
			start_idx = append(start_idx, from_idx)
		}
	}

	if len(curr_pattern) > 0 {
		patterns = append(patterns, curr_pattern)
	}

	//fmt.Println(start_idx)
	//fmt.Println(patterns)

	pdb_seq_idx := 0
	for i := range patterns {
		//mapped_idx := strings.Index(ref_seq, patterns[i])
		pf := zstr.NewPatternFinder(patterns[i])
		mapped_idxs := pf.FindIn(ref_seq)
		mapped_idx := mapped_idxs[0]

		for j:=0; j < len(patterns[i]); j++ {
			rslt += fmt.Sprintf("(%s, %d) ", pdb_seq[pdb_seq_idx][0], mapped_idx)
			mapped_idx++
			pdb_seq_idx++
		}
	}

	fmt.Println(rslt)
}

func main() {
	ref_seq := "FMVFLVLLPSVLSQCNV"
	pdb_seq := [][2]string{[2]string{"c.4.", "F"}, [2]string{"c.5.", "L"}, [2]string{"c.6.", "V"}, [2]string{"c.6.A", "L"}, [2]string{"c.6.B", "L"}, [2]string{"c.7.", "P"}, [2]string{"c.12.", "Q"}, [2]string{"c.13.", "C"}}
	find_mapping(ref_seq, pdb_seq)
}
