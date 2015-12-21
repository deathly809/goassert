
package goassert

import "log"

// Assert takes in a test and if the test does
// not pass we output the message given and 
// terminate
func Assert(passed bool,msg ...interface{}) {
    if !passed {
        log.Fatal(msg...)
    }
}