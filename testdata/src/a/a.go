package main

import "fmt"

type Animal interface{ Speak() }

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

func speakIfAnimal(a any) {
	if a, ok := a.(Animal); ok {
		fmt.Println("is an animal")
		a.Speak()
		return
	}

	fmt.Println("is not an animal")
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
