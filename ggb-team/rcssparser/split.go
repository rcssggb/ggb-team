package rcssparser

import "strings"

func breakArgs(argStr string) []string {
	argStr = strings.TrimPrefix(argStr, "(")
	argStr = strings.TrimSuffix(argStr, ")")
	return strings.Split(argStr, " ")
}
