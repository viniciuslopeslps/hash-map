package pkg

type ArrayList struct {
	Actual string
	Next   *ArrayList
}

func (al *ArrayList) Append(vale string) {
	if al.Next == nil {
		al.Next = &ArrayList{
			Actual: vale,
		}
	} else {
		al.Next.Append(vale)
	}
}

func (al *ArrayList) Get(value string) *string {
	if al.Next == nil {
		return nil
	}

	if al.Actual == value {
		return &value
	}

	return al.Next.Get(value)
}
