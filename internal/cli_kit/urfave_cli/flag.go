package urfave_cli

import (
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"
)

func in(target string, strArray []string) bool {
	sort.Strings(strArray)

	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}

	return false
}

// UrfaveCliAppendCliFlag
//
//	append cli.Flag
//	if flag name exists will panic: do UrfaveCliAppendCliFlag err, flag exists name xxx at [xxx]
func UrfaveCliAppendCliFlag(target []cli.Flag, elem []cli.Flag) []cli.Flag {
	if len(elem) == 0 {
		return target
	}

	var appendFlagName []string
	for _, flag := range elem {
		appendFlagName = append(appendFlagName, flag.Names()...)
	}

	if len(target) > 0 { // check target name exists
		for _, flag := range target {
			targetNames := flag.Names()
			if len(targetNames) > 0 {
				for _, name := range targetNames {
					if in(name, appendFlagName) {
						panic(
							fmt.Errorf(
								"do UrfaveCliAppendCliFlag err, flag exists name %s at %v",
								name,
								targetNames,
							),
						)
					}
				}
			}
		}
	}

	return append(target, elem...)
}

// UrfaveCliAppendCliFlags
//
//	append cli.Flag
func UrfaveCliAppendCliFlags(target []cli.Flag, elems ...[]cli.Flag) []cli.Flag {
	if len(elems) == 0 {
		return target
	}

	for _, elem := range elems {
		if len(elem) > 0 {
			target = UrfaveCliAppendCliFlag(target, elem)
		}
	}

	return target
}
