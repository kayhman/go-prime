package main

import (
	"flag"
)


var maxPtr = flag.Int64("max", 100000, "Max prime size")
var msgPtr = flag.Int64("msg", 666, "Message to encrypt")
