# missingtypeguard

The development assignments for Gopher Enablement Internship day2&3

missingtypeguard is a tool that checks if types that implement an interface have a type guard for it.

## Example

```go
package main

import "fmt"

type Animal interface{ Speak() string }

type dog struct{}
func (d dog) Speak() string { return "woof" }
var _ Animal = dog{} // 😃 dog has a type guard for Animal

type dogMissingTypeGuard struct{} // 😡 dogMissingTypeGuard is missing atype guard for Animal"
func (d dogMissingTypeGuard) Speak() string { return "woof" }

func speakIfAnimal(a any) {
    if a, ok := a.(Animal); ok {
        fmt.Printf("%T is an animal: %s\n", a, a.Speak())
    }
}

func main() {
    speakIfAnimal(dog{})
    speakIfAnimal(dogMissingTypeGuard{})
}
```

## Usage

```bash
go install github.com/ras0q/missingtypeguard/cmd/missingtypeguard@latest
go vet -vettool=$(which missingtypeguard) ./...
```
