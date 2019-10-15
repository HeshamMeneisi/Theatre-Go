module theatre-mgr

go 1.13

require (
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/gotestyourself/gotest.tools v2.2.0+incompatible
	github.com/pkg/errors v0.8.1 // indirect
	gotest.tools v2.2.0+incompatible
	theatre-mgr/generators v0.0.0
	theatre-mgr/theatre v0.0.0
)

replace theatre-mgr/theatre => ./theatre

replace theatre-mgr/generators => ./generators
