package domain

const (
	Com = ".com"
	Net = ".net"
	Org = ".org"
	Io  = ".io"
)

func AllDomains() []string {
	return []string{Com, Net, Org}
}
