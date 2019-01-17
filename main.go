
package main

import (
  "fmt"
	//"net/http"
	"strconv"
	//"strings"
	//"bufio"
	"os"
	"os/exec"
	//"time"
	//"log"
)




var debug = true // Always have a quick switch to display debug data




// ========== START: Golang Console Colors ========== ========== ========== ==========
// Golang Console Colors
// Example: fmt.Print( cRed + "HelloWorld" + cClr )
var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"
// ========== END: Golang Console Colors ========== ========== ========== ==========




// ========== START: Output Simplification ========== ========== ========== ==========
var breakspace = "\n"
var breakline = breakspace + cBlue + "  ====================================================" + cClr + breakspace
// ========== END: Output Simplification ========== ========== ========== ==========




// ========== START: Console Splash ========== ========== ========== ==========
var appinfo = `
  ` + cBlue + `====================================================` + cBold + cCyan + `
   _____      _   _    _
  |  __ \    | | | |  | |
  | |__) |_ _| |_| |__| | __ _ _   _  __ _  ___ _ __
  |  ___/ _`+"`"+` | __|  __  |/ _`+"`"+` | | | |/ _`+"`"+` |/ _ \ '_ \
  | |  | (_| | |_| |  | | (_| | |_| | (_| |  __/ | | |
  |_|   \__,_|\__|_|  |_|\__,_|\__,_|\__, |\___|_| |_|
                                     __/ |
                                    |___/` + cClr + `
  ` + cCyan + `Interview: ` + cWhite + `citrusbyte` + cClr + `
  ` + cCyan + `https://github.com/` + cYellow + `pathaugen` + cCyan + `/interview-citrusbyte` + cClr + `
  ` + cBlue + `====================================================` + cClr + `
`
// ========== END: Console Splash ========== ========== ========== ==========


func main() {

  // ========== START: Clear Screen ========== ========== ========== ==========
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
  // ========== END: Clear Screen ========== ========== ========== ==========

  fmt.Print( appinfo + breakspace )




  // ========== START: Main Logic ========== ========== ========== ==========
  // Write some code, that will flatten an array of arbitrarily nested arrays of integers into a flat array of integers. e.g. [[1,2,[3]],4] -> [1,2,3,4].

  fmt.Print(
    cBold + cCyan + "  Problem to Solve:" + cClr + breakspace +
    "  Write some code, that will flatten an array of arbitrarily" + breakspace +
    "  nested arrays of integers into a flat array of integers." + breakspace +
    "  e.g. [[1,2,[3]],4] -> [1,2,3,4]" )

  fmt.Print( breakspace + breakline + breakspace )

  fmt.Print(
    cBold + cCyan + "  Solution Function:" + cClr + breakspace +
    "  " + cMagenta + "func " + cBlue + "arrayFlatten" + cClr + "( " + cCyan + "inputArray" + cClr + " ) " + cCyan + "outputArray" + cClr + " {}" + breakspace +
    breakspace +
    cBold + cCyan + "  Input:" + cClr + breakspace +
    "  [[1,2,[3]],4] -> Converted to Golang Multi Dimensional Array" + breakspace ) // input: [[1,2,[3]],4]

  // Generation of the input value ( [[1,2,[3]],4] )
  var input = [][][]int{ { {1}, {2}, {3} }, { {4} } }

  // Debug the input array values (controlled via debug bool)
  if debug {
    fmt.Print( breakspace + cBold + cCyan + "  Debug Output (controlled via debug bool):" + cClr + breakspace )
    fmt.Print( "  input[0][0][0] = " + cYellow + strconv.Itoa( input[0][0][0] ) + cClr + breakspace ) // 1
    fmt.Print( "  input[0][1][0] = " + cYellow + strconv.Itoa( input[0][1][0] ) + cClr + breakspace ) // 2
    fmt.Print( "  input[0][2][0] = " + cYellow + strconv.Itoa( input[0][2][0] ) + cClr + breakspace ) // 3
    fmt.Print( "  input[1][0][0] = " + cYellow + strconv.Itoa( input[1][0][0] ) + cClr + breakspace ) // 4
  }

  fmt.Print(
    breakspace +
    cBold + cCyan + "  Output Utilizing arrayFlatten(input)" + breakspace +
    "  Displayed as string via arrayToString():" + cClr + breakspace +
    "  [" + arrayToString( arrayFlatten(input), "," ) + "]" ) // output expected: [1,2,3,4]

  fmt.Print( breakspace + breakline )

  // Test all functions via > go test -v
  fmt.Print( testCode() )

  fmt.Print( breakline )

  // Benchmark all functions via > go test -bench . -benchtime=1s
  fmt.Print( benchCode() )

  // ========== END: Main Logic ========== ========== ========== ==========




  fmt.Print( breakline )
}
