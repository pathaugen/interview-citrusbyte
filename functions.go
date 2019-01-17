
package main

import (
	"fmt"
	//"os"
	"os/exec"

	//"strconv" // https://golang.org/pkg/strconv/
  "strings"

  //"reflect" // Determine element type
)


// ========== START: arrayFlatten ========== ========== ========== ==========
// Write some code, that will flatten an array of arbitrarily nested arrays of integers into a flat array of integers. e.g. [[1,2,[3]],4] -> [1,2,3,4].
func arrayFlatten( inputArray [][][]int ) []int {
  // Example Input: [][][]int{ { {1}, {2}, {3} }, { {4} } }

  returnArray := []int{}

  // Could make a function that calls itself for recursion, but we have a predictable multidimensional array here
  for i := 0; i < len(inputArray); i++ {
    for j := 0; j < len(inputArray[i]); j++ {
      for k := 0; k < len(inputArray[i][j]); k++ {
        returnArray = append(returnArray, inputArray[i][j][k])
      }
    }
  }

  //returnArray := []int{1,2,3,4,5,6} // Debug
  return returnArray
}
// ========== END: arrayFlatten ========== ========== ========== ==========


// ========== START: arrayToString ========== ========== ========== ==========
func arrayToString(a []int, delim string) string {
  // Convert array of ints to a string - three different ways to show how benchmarking is the way to decide how to proceed for any function/application

  // NOTE: With multiple ways to perform the same task, benchmarking each is the proper way to see which is fastest and to utilize
  //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]") // Benchmark result per op (input []int{1,2,3,4,5,6}) = 2828 ns/op (slowest of three)
  //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]") // Benchmark result per op (input []int{1,2,3,4,5,6}) = 2736 ns/op
  return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]") // Benchmark result per op (input []int{1,2,3,4,5,6}) = 2486 ns/op (fastest of three)
}
// ========== START: arrayToString ========== ========== ========== ==========





// ========== START: statusColorize ========== ========== ========== ==========
// Colorize the testing output for console
func statusColorize(input string) string {
  output := input

  output = strings.Replace( output, "RUN", cCyan + "RUN" + cClr, -1 )
  output = strings.Replace( output, "PASS", cBold + cGreen + "PASS" + cClr, -1 )
  output = strings.Replace( output, "FAIL", cRed + "FAIL" + cClr, -1 )
  output = strings.Replace( output, "ok ", cGreen + "ok " + cClr, -1 )

  return output
}
// ========== END: statusColorize ========== ========== ========== ==========




// ========== START: testCode - show the Golang tests in the UI ========== ========== ========== ==========
func testCode() string {
  consoleCommand := `go test -v` // Single place to alter the command line command
  output := ``

  output += breakspace + cBold + cCyan + "  Golang Function Tests:" + cClr
  output += breakspace + cYellow + "  > " + cClr + consoleCommand + breakspace

  output += breakspace

  testOutput, _ := exec.Command( "cmd", "/c", consoleCommand ).Output()

  // Alternative code:
  //cmd.Stdout = os.Stdout
  //cmd.Run()

  return output + statusColorize( string(testOutput) )
}
// ========== END: testCode - show the Golang tests in the UI ========== ========== ========== ==========




// ========== START: benchCode - show the Golang benchmarks in the UI ========== ========== ========== ==========
func benchCode() string {
  consoleCommand := `go test -bench . -benchtime=1s` // Single place to alter the command line command
  output := ``

  output += breakspace + cBold + cCyan + "  Golang Function Benchmarks:" + cClr
  output += breakspace + cYellow + "  > " + cClr + consoleCommand + breakspace

  output += breakspace

  testOutput, _ := exec.Command( "cmd", "/c", consoleCommand ).Output()

  // Alternative code:
  //cmd.Stdout = os.Stdout
  //cmd.Run()

  return output + statusColorize( string(testOutput) )
}
// ========== END: benchCode - show the Golang benchmarks in the UI ========== ========== ========== ==========
