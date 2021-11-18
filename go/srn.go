package main

import "strings"

// SRNs have the following structure:
// <ORG>:<MODEL>:<SLUG>:<VERSION>[:<IDENTIFIER>]

type ParsedSrn struct {
	Srn string
}

func (ps *ParsedSrn) components() []string {
	return strings.Split(ps.Srn, ".")
}

func (ps *ParsedSrn) Org() string {
	return ps.components()[0]
}

func (ps *ParsedSrn) Model() string {
	return ps.components()[1]
}

func (ps *ParsedSrn) Slug() string {
	return ps.components()[2]
}

func (ps *ParsedSrn) Version() string {
	return ps.components()[3]
}

func (ps *ParsedSrn) Identifier() string {
	if len(ps.components()) == 5 {
		return ps.components()[4]
	}
	return ""
}
