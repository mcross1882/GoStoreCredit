/**
 * My first GO program. It solves the problem at 
 * http://code.google.com/codejam/contest/351101/dashboard#s=p0
 *
 * @since  0.1
 * @author Matthew Cross <matthew@pmg.co>
 */
package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
)

/**
 * SampleCase represents a single "sample" in the test file
 * 
 * @since  0.1
 */
type SampleCase struct {
    credit    int
    itemCount int
    store     []int
}

/**
 * Finds the solution (if any) for a given SampleCase
 *
 * @since  0.1
 * @param  *SampleCase
 * @return int, int
 */
func findSolution(s *SampleCase) (resA int, resB int) {
    resA, resB = -1, -1
    for x := 0; x < s.itemCount; x++ {
        for y := 0; y < s.itemCount; y++ {
            if x == y { continue }
            
            if (s.store[x] + s.store[y]) == s.credit {
                resA, resB = x+1, y+1
                
                if resA > resB {
                    resA, resB = resB, resA
                }
                break
            }
        }
    }
    return
}

/**
 * Strips whitespace and converts a string literal
 * to a number using Atoi
 *
 * @since  0.1
 * @param  string
 * @return int
 */
func convertToNumber(line string) (res int) {
    res, err := strconv.Atoi(strings.TrimSpace(line))
    if err != nil {
        fmt.Printf("Failed to extract number from " + line + "\n")
        return -1
    }
    return
}

/**
 * Extracts a line from a bufio reader
 *
 * @since  0.1
 * @param  bufio.Reader
 * @return string
 */
func extractLine(r *bufio.Reader) (line string) {
    line, err := r.ReadString('\n')
    if err != nil {
        fmt.Printf("%s", err)
        return ""
    }
    return 
}

func parseRow(s *SampleCase, r *bufio.Reader) {
    line := extractLine(r)
    s.credit = convertToNumber(line)
    
    line = extractLine(r)
    s.itemCount = convertToNumber(line)
    
    items := strings.Split(extractLine(r), " ")
    s.store = make([]int, s.itemCount);
    
    for i := 0; i < s.itemCount; i++ {
        s.store[i] = convertToNumber(items[i])
    }
    
    solutionA, solutionB := findSolution(s)
    
    if solutionA == -1 || solutionB == -1 {
        fmt.Printf("No solutions\n")
    } else {
        fmt.Printf("%d %d\n", solutionA, solutionB)
    }
}

func main() {
    args := os.Args[1:]
    
    if len(args) <= 0 {
        log.Fatal("Not enough arguments given")
        return
    }
    
    f, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
        return
    }
    
    defer f.Close()
    r := bufio.NewReader(f)
    
    numOfRecords := convertToNumber(extractLine(r))
    fmt.Printf("Number of Records: %d\n", numOfRecords)
    
    sample := new(SampleCase)
    for i := 0; i < numOfRecords; i++ {
        fmt.Printf("Case #%d: ", i+1)
        parseRow(sample, r)
    }
    
    fmt.Printf("Finished!")
}