package domain

const (
	Com  = ".com"
	Net  = ".net"
	Org  = ".org"
	Co   = ".co"
	Ai   = ".ai"
	CoUk = ".co.uk"
	Ca   = ".ca"
	Me   = ".me"
	Io   = ".io"
	Dev  = ".dev"
	Eu   = ".eu"
	Xyz  = ".xyz"
)

func AllDomains() []string {
	return []string{Com, Net, Org, Co, Ai, CoUk, Ca, Me, Io, Dev, Eu, Xyz}
}
