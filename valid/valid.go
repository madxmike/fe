package valid

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}