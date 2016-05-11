package readline

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Get(out interface{}) {
	in := bufio.NewReader(os.Stdin)
	switch typedout := out.(type) {
	case *string:
		// Error usually EOF, which shouldn't even happen with Stdin
		// default return value should be the right thing nevertheless.
		input, _ := in.ReadString('\n')
		*typedout = strings.TrimRight(input, "\n")
	case *int:
		// when ParseInt returns an error, the default value it returns
		// is still acceptable for a game situation, so ignore it
		input, _ := in.ReadString('\n')
		intval, _ := strconv.ParseInt(regexp.MustCompile(`[^0-9]`).ReplaceAllString(input, ""), 10, 0)

		*typedout = int(intval) // conversion from int64 to int, should work because of the last 0 argument in ParseInt
	default:
		panic("readline invoked with non-pointer type or pointer to unsupported type.")
	}
}
