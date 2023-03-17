package impl

import "multipackage/sub"

type Hoge struct{}

func (*Hoge) DoSomething() {}

var _ sub.Interface = (*Hoge)(nil)

type HogeMissingTypeGuard struct{} // want "the pointer of multipackage/sub/impl.HogeMissingTypeGuard is missing a type guard for multipackage/sub.Interface"

func (*HogeMissingTypeGuard) DoSomething() {}
