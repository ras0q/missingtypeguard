package constructor

type I interface {
	DoSomething()
}

// missing type guard or constructor

type s struct{} // want "the pointer of constructor.s is missing a type guard for constructor.I"

func (*s) DoSomething() {}

// with constructor
// TODO: allow to omit type guard if the type has a constructor

type sWithConstructor struct{} // want "the pointer of constructor.sWithConstructor is missing a type guard for constructor.I"

func (s *sWithConstructor) DoSomething() {}

func NewSWithConstructor() I {
	return &sWithConstructor{}
}

// with constructor which has assignment

type sWithConstructorWithAssignment struct{}

func (s *sWithConstructorWithAssignment) DoSomething() {}

func NewSWithConstructorWithAssignment() I {
	var s I = &sWithConstructorWithAssignment{}
	return s
}
