
package main

import (
  "testing"
  "fmt"
  //"strings"
  //"reflect"
)


// In this file we define tests on the functions as well as benchmarks to check which solutions are in fact the fastest and most cost efeective (from a processing per hour standpoint)
// NOTE: Functions in _test.go files are utilized only for testing and don't affect the main application


// ========== START: Testing arrayFlatten ========== ========== ========== ==========
// Make the base test private (lower case letter) then have the other functions invoke it for testing
func testArrayFlatten( input [][][]int, expectedResult []int, t *testing.T ) {
  //input := [][][]int{ {{1}}, {{2}}, {{3}} }
  //expectedResult := []int{1,2,3}

  t.Log( "Testing arrayFlatten: Using input [][][]int{" + arrayToString( arrayFlatten(input), ",") + "}... (expected result: []int{" + arrayToString(expectedResult, ",") + "})" )

  //result := Add(12,28)
  result := arrayFlatten(input)
  t.Log( "DEBUG: result = [" + arrayToString(result, ",") + "]" ) // Debug  

  // Reflect is one way to check array equality, but we'll be using our arrayToString for comparisons
  //if reflect.DeepEqual(expectedResult, result) {
  if arrayToString(result, ",") != arrayToString(expectedResult, ",") {
    t.Errorf( "Expected result of [" + arrayToString(expectedResult, ",") + "], but it was [" + arrayToString(result, ",") + "] instead." )
  } else {
    //t.Log( "[" + arrayToString(expectedResult, ",") + "] == [" + arrayToString(result, ",") + "]" ) // Debug
  }

}
func TestArrayFlatten1(t *testing.T) { testArrayFlatten( [][][]int{ { {1}, {2}, {3} }, { {4} } }, []int{1,2,3,4}, t ) }
func TestArrayFlatten2(t *testing.T) { testArrayFlatten( [][][]int{ { {1}, {2}, {3} }, { {4} }, { {5, 6} } }, []int{1,2,3,4,5,6}, t ) }
// ========== END: Testing arrayFlatten ========== ========== ========== ==========

// ========== START: Benchmark arrayFlatten ========== ========== ========== ==========
// Make the base benchmark private (lower case letter) then have the other functions invoke it for testing
func benchmarkArrayFlatten( s string, b *testing.B ) {
  // Put any initialization before timer reset
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    fmt.Sprintf("hello, " + s)
  }
}
func BenchmarkArrayFlatten1(b *testing.B) { benchmarkArrayFlatten("111111111111111111111111111111111111111111111111111111111111111111111111111111", b) }
func BenchmarkArrayFlatten2(b *testing.B) { benchmarkArrayFlatten("222", b) }
// ========== END: Benchmark arrayFlatten ========== ========== ========== ==========



// ========== START: Testing arrayToString ========== ========== ========== ==========
// Make the base test private (lower case letter) then have the other functions invoke it for testing
func testArrayToString( input []int, expectedResult string, t *testing.T ) {
  t.Log("Testing arrayToString: Using input []int{" + arrayToString(input, ",") + "}... (expected result: [" + expectedResult + "])")
  result := arrayToString(input, ",")
  if result != expectedResult {
    t.Errorf( "Expected result of [" + expectedResult + "], but it was [" + result + "] instead." )
  }
}
func TestArrayToString1(t *testing.T) { testArrayToString( []int{1,2,3}, "1,2,3", t ) }
func TestArrayToString2(t *testing.T) { testArrayToString( []int{1,2,3,4,5,6}, "1,2,3,4,5,6", t ) }
// ========== END: Testing arrayToString ========== ========== ========== ==========

// ========== START: Benchmark arrayToString ========== ========== ========== ==========
// Make the base benchmark private (lower case letter) then have the other functions invoke it for testing
func benchmarkArrayToString( input []int, b *testing.B ) {
  // Put any initialization before timer reset
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    arrayToString(input, ",")
  }
}
func BenchmarkArrayToString1(b *testing.B) { benchmarkArrayToString( []int{1,2,3}, b ) }
func BenchmarkArrayToString2(b *testing.B) { benchmarkArrayToString( []int{1,2,3,4,5,6}, b ) }
// ========== END: Benchmark arrayToString ========== ========== ========== ==========
