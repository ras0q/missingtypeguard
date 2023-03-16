package impl

import "multipackage"

type hoge struct{}

func (*hoge) DoSomething() {}

var _ multipackage.Interface = (*hoge)(nil)

type hogeMissingTypeGuard struct{} // want "the pointer of multipackage/impl.hogeMissingTypeGuard is missing a type guard for multipackage.Interface"

func (*hogeMissingTypeGuard) DoSomething() {}
