package models

type (
	Maths struct{}

	Answer struct {
		Answer int64 `json:"answer"` // Answer
	}
)

func (m Maths) SolveAddition(numA, numB int64) (*Answer, error) {
	answer := new(Answer)

	// Get answer
	answer.Answer = numA + numB

	return answer, nil
}

func (m Maths) SolveSubstraction(numA, numB int64) (*Answer, error) {
	answer := new(Answer)

	// Get answer
	answer.Answer = numA - numB

	return answer, nil
}

func (m Maths) SolveMultiplication(numA, numB int64) (*Answer, error) {
	answer := new(Answer)

	// Get answer
	answer.Answer = numA * numB

	return answer, nil
}

func (m Maths) SolveDivision(numA, numB int64) (*Answer, error) {
	answer := new(Answer)

	if numB == 0 {
		return nil, ErrDivisionByZero
	}
	// Get answer
	answer.Answer = numA / numB

	return answer, nil
}
