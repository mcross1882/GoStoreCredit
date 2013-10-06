package main

import "testing"

func TestFindSolution(t *testing.T) {
    sample := new(SampleCase)
    
    sample.credit    = 200
    sample.itemCount = 5
    sample.store     = append(sample.store, 10, 40, 35, 120, 160)
    
    resA, resB := findSolution(sample)
    
    if resA != 2 || resB != 5 {
        t.Fatal("findSolution(...) failed")
    }
}

func TestConvertToNumber(t *testing.T) {
    line := "1234\n"
    num := convertToNumber(line)
    
    if num != 1234 {
        t.Fatal("convertToNumber(...) failed")
    }
}