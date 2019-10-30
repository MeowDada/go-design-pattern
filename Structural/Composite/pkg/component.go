package composite

type Contest struct {}

func (c Contest) DisplayEntry() string {
	return "1st prize of a music contest"
}

type Skill struct {}

func (s Skill) DisplayEntry() string {
	return "Nice cooking skill"
}

type Experience struct {}

func (e Experience) DisplayEntry() string {
	return "5 years of related working experiences"
}