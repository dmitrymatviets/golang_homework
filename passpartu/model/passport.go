package model

import "fmt"

const (
	SexMale   = "Male"
	SexFemale = "Female"
)

type Sex string

type BasePassport interface {
	GetName() string
	GetCredentials() string
	GetSex() Sex
	GetMarried() bool
}

type InternationalPassport interface {
	GetName() string
	GetCredentials() string
	GetSex() Sex
}



type AmericanPassport struct {
	FirstName string
	LastName  string
	Sex       Sex
	SSN       string
	IsMarried bool
}

func (p *AmericanPassport) GetName() string {
	return p.FirstName + " " + p.LastName
}

func (p *AmericanPassport) GetCredentials() string {
	return p.SSN
}

func (p *AmericanPassport) GetSex() Sex {
	return p.Sex
}

func (p *AmericanPassport) GetMarried() bool {
	return p.IsMarried
}



type RussianInternationalPassport struct {
	FirstName  string
	MiddleName string
	LastName   string
	Sex        Sex
	Serie      string
	Number     int64
}

func (p *RussianInternationalPassport) GetName() string {
	return fmt.Sprintf("%s %c. %c.", p.LastName, []rune(p.FirstName)[1], []rune(p.MiddleName)[1])
}

func (p *RussianInternationalPassport) GetCredentials() string {
	return fmt.Sprintf("%s %d", p.Serie, p.Number)
}

func (p *RussianInternationalPassport) GetSex() Sex {
	return p.Sex
}



type RussianBasePassport struct {
	FirstName  string
	MiddleName string
	LastName   string
	Sex        Sex
	Serie      string
	Number     int64
	IsMarried  bool
}

func (p *RussianBasePassport) GetName() string {
	return fmt.Sprintf("%s %s %s", p.LastName, p.FirstName, p.MiddleName)
}

func (p *RussianBasePassport) GetCredentials() string {
	return fmt.Sprintf("%s-%d", p.Serie, p.Number)
}

func (p *RussianBasePassport) GetSex() Sex {
	return p.Sex
}

func (p *RussianBasePassport) GetMarried() bool {
	return p.IsMarried
}
