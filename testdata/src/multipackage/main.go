package main

import (
	"multipackage/sub"
	"multipackage/sub/impl"
)

func main() {
	run(&impl.Hoge{})
	run(&impl.HogeMissingTypeGuard{})
}

func run(i interface{}) {
	if i, ok := i.(sub.Interface); ok {
		i.DoSomething()
	}
}
