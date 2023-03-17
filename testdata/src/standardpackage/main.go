package main

import "io"

type myWriter struct{} // want "the pointer of standardpackage.myWriter is missing a type guard for io.Writer"

func (w *myWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func run(a any) {
	if a, ok := a.(io.Writer); ok {
		a.Write(nil)
	}
}

func main() {
	run(&myWriter{})
}
