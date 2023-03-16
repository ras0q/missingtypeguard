package a

import "fmt"

type Animal interface{ Speak() }

func f() {
	var a Animal

	a = dog{}
	a.Speak()

	a = dogMissingTypeGuard{}
	a.Speak() // want "dogMissingTypeGuard is used as Animal, but it does not have a type guard"

	a = &cat{}
	a.Speak()

	a = &catMissingTypeGuard{}
	a.Speak() // want "catMissingTypeGuard is used as Animal, but it does not have a type guard"

	a = bird(0)
	a.Speak()

	a = birdMissingTypeGuard(0)
	a.Speak() // want "birdMissingTypeGuard is used as Animal, but it does not have a type guard"

	var f fish
	a = &f
	a.Speak()

	var f2 fishMissingTypeGuard
	a = &f2
	a.Speak() // want "fishMissingTypeGuard is used as Animal, but it does not have a type guard"
}

// struct value receiver

type dog struct{}

func (d dog) Speak() { fmt.Println("woof") }

var _ Animal = dog{}

type dogMissingTypeGuard struct{}

func (d dogMissingTypeGuard) Speak() { fmt.Println("woof") }

// struct pointer receiver

type cat struct{}

func (c *cat) Speak() { fmt.Println("meow") }

var _ Animal = (*cat)(nil)

type catMissingTypeGuard struct{}

func (c *catMissingTypeGuard) Speak() { fmt.Println("meow") }

// defined type value receiver

type bird int

func (b bird) Speak() { fmt.Println("tweet") }

var _ Animal = bird(0)

type birdMissingTypeGuard int

func (b birdMissingTypeGuard) Speak() { fmt.Println("tweet") }

// defined type pointer receiver

type fish int

func (f *fish) Speak() { fmt.Println("blub") }

var _ Animal = (*fish)(nil)

type fishMissingTypeGuard int

func (f *fishMissingTypeGuard) Speak() { fmt.Println("blub") }
