package opt

type Status int

const (
	NotTerminated (Status) = 0

	DerivAbsConv = 1
	DerivRelConv = 2

	ObjAbsConv = 3
	ObjRelConv = 4

	XConvAbs = 5
	XConvRel = 6

	Success = 7

	IterLimit = -1
	TimeLimit = -2

	NumericalToleranceReached = -3

	Infeasible = -4

	Fail = -5
)
