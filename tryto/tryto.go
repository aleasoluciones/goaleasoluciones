package tryto

func TryTo(wrapped func() (interface{}, error), maximumAttempts int) (interface{}, error) {
	attempts := 1
	result, err := wrapped()
	for err != nil && attempts < maximumAttempts {
		attempts++
		result, err = wrapped()
	}
	return result, err
}
