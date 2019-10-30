package composite

type Composite struct {
	resumes []Resume
}

func (c *Composite) DisplayEntry() string {
	var statement string
	for _, resume := range c.resumes {
		statement += resume.DisplayEntry() + ", "
	}
	return statement
}

func (c *Composite) AddResume(r Resume) {
	c.resumes = append(c.resumes, r)
}

/* Not yet implemented
func (c Composite) RemoveResume() {}
func (c Composite) GetResume() Resume {}
*/

