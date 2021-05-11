package requests

type DomainSetHostRequest struct {
	Domain string
	Hosts  []DomainSetHost
}

type DomainSetHost struct {
	HostName   string
	RecordType string
	Address    string
	MXPref     int
}
