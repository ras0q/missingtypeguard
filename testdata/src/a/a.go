package main

import "fmt"

type Animal interface{ Speak() string }

// struct value receiver

type dog struct{}

func (d dog) Speak() string { return "woof" }

var _ Animal = dog{}

type dogMissingTypeGuard struct{} // want "dogMissingTypeGuard is missing a type guard for Animal"

func (d dogMissingTypeGuard) Speak() string { return "woof" }

// struct pointer receiver

type cat struct{}

func (c *cat) Speak() string { return "meow" }

var _ Animal = (*cat)(nil)

type catMissingTypeGuard struct{}

func (c *catMissingTypeGuard) Speak() string { return "meow" }

// defined type value receiver

type bird int

func (b bird) Speak() string { return "tweet" }

var _ Animal = bird(0)

type birdMissingTypeGuard int

func (b birdMissingTypeGuard) Speak() string { return "tweet" }

// defined type pointer receiver

type fish int

func (f *fish) Speak() string { return "blub" }

var _ Animal = (*fish)(nil)

type fishMissingTypeGuard int

func (f *fishMissingTypeGuard) Speak() string { return "blub" }

func speakIfAnimal(a any) {
	if a, ok := a.(Animal); ok {
		fmt.Printf("%T is an animal: %s\n", a, a.Speak())
		return
	}

	fmt.Printf("%T is not an animal\n", a)
}

func main() {
	var (
		dog1  = dog{}
		dog2  = dogMissingTypeGuard{}
		cat1  = cat{}
		cat2  = catMissingTypeGuard{}
		bird1 = bird(0)
		bird2 = birdMissingTypeGuard(0)
		fish1 = fish(0)
		fish2 = fishMissingTypeGuard(0)
	)

	speakIfAnimal(dog1)
	speakIfAnimal(&dog1)
	speakIfAnimal(dog2)
	speakIfAnimal(&dog2)
	speakIfAnimal(cat1)
	speakIfAnimal(&cat1)
	speakIfAnimal(cat2)
	speakIfAnimal(&cat2)
	speakIfAnimal(bird1)
	speakIfAnimal(&bird1)
	speakIfAnimal(bird2)
	speakIfAnimal(&bird2)
	speakIfAnimal(fish1)
	speakIfAnimal(&fish1)
	speakIfAnimal(fish2)
	speakIfAnimal(&fish2)
	speakIfAnimal("gopher")
}
