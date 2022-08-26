package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CF691C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	olddot, newdot := len(s), -1
	as := make([]rune, 0)
	for i, c := range s {
		if c == '.' {
			olddot = i
		} else if c > '0' && newdot == -1 {
			newdot = i + 1
			as = append(as, c)
			as = append(as, '.')
		} else if newdot != -1 {
			as = append(as, c)
		}
	}

	i := len(as) - 1
	for ; i >= 0; i-- {
		if as[i] == '.' {
			i--
			break
		} else if as[i] != '0' {
			break
		}
	}
	as = as[:i+1]

	var ans strings.Builder
	ans.WriteString(string(as))

	if olddot-newdot != 0 {
		ans.WriteRune('E')
		if olddot-newdot < 0 {
			olddot++
		}

		ans.WriteString(strconv.Itoa(olddot - newdot))
	}
	fmt.Fprintln(out, ans.String())
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF691C(in, os.Stdout)
}
