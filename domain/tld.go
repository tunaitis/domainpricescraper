package domain

const (
	Com = ".com"
	Net = ".net"
	Org = ".org"
	Io  = ".io"
	Dev = ".dev"
)

func AllDomains() []string {
	return []string{Com, Net, Org, Io, Dev}
}
