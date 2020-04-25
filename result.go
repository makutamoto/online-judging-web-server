package main

type resultType int

const (
	resultAccepted resultType = iota
	resultWrongAnswer
	resultReferenceError
	resultTimeLimitExceeded
	resultCompileError
	resultInternalError
)

func (result resultType) String() string {
	switch result {
	case resultAccepted:
		return "AC"
	case resultInternalError:
		return "IE"
	case resultCompileError:
		return "CE"
	case resultReferenceError:
		return "RE"
	case resultTimeLimitExceeded:
		return "TLE"
	case resultWrongAnswer:
		return "WA"
	}
	return "IE"
}

func (result *resultType) update(new resultType) {
	if *result < new {
		*result = new
	}
}
