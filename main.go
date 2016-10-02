package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type scanner struct {
	sc *bufio.Scanner
}

func (sc *scanner) next() string {
	sc.sc.Scan()
	return sc.sc.Text()
}

func (sc *scanner) nextInt() int {
	i, err := strconv.Atoi(sc.next())
	if err != nil {
		panic(err)
	}
	return i
}

func newScanner(r io.Reader, split bufio.SplitFunc) *scanner {
	sc := bufio.NewScanner(r)
	sc.Split(split)
	return &scanner{
		sc: sc,
	}
}

var wsc = newScanner(os.Stdin, bufio.ScanWords)
var lsc = newScanner(os.Stdin, bufio.ScanLines)

func next() string {
	return wsc.next()
}

func nextInt() int {
	return wsc.nextInt()
}

func nextLine() string {
	return lsc.next()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func joinInts(a []int, sep string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

type stringStack struct {
	stack []string
}

func (ss *stringStack) push(s string) {
	ss.stack = append(ss.stack, s)
}

func (ss *stringStack) pop() (string, error) {
	l := len(ss.stack)

	if l == 0 {
		return "", errors.New("Empty stack")
	}

	s := ss.stack[l-1]
	ss.stack = ss.stack[:l-1]

	return s, nil
}

func newStringStack(l, c int) *stringStack {
	return &stringStack{make([]string, l, c)}
}

func main() {}
