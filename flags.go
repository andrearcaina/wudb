package main

import (
	"flag"
	"fmt"
)

var (
	aof          = flag.Bool("aof", false, "persistence mode")
	aofShort     = flag.Bool("a", false, "persistence mode (shorthand)")     // shorthand for -aof
	aofPath      = flag.String("path", "wudb.aof", "aof file path")          // only used if -aof or -a is true
	aofPathShort = flag.String("p", "wudb.aof", "aof file path (shorthand)") // shorthand for -path
	addr         = flag.String("addr", "8080", "server address")
)

func firstTrue(values ...bool) bool {
	for _, v := range values {
		if v {
			return true
		}
	}
	return false
}

func firstNonDefault(value, fallback, defaultVal string) string {
	if value != defaultVal {
		return value
	}
	return fallback
}

func ParseFlags() (string, bool, string) {
	flag.Parse()

	persistence := firstTrue(*aof, *aofShort)

	finalAOFPath := firstNonDefault(*aofPathShort, *aofPath, "wudb.aof")

	if !persistence && finalAOFPath != "wudb.aof" {
		fmt.Printf("Warning: AOF path specified (%s) but AOF flag is set to %t. Ignoring path.\n", finalAOFPath, persistence)
		return *addr, false, ""
	}

	return *addr, persistence, finalAOFPath
}
