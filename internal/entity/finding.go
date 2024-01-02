package entity

type Finding struct {
	Name   string
	Desc   string
	Status bool
}

func CreateSuccessFinding(name string, desc string) *Finding {
	return &Finding{
		Desc:   desc,
		Status: true,
	}
}

func CreateFailureFinding(name string, desc string) *Finding {
	return &Finding{
		Desc:   desc,
		Status: false,
	}
}
