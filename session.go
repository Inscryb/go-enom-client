package enom

type Session struct {
	ResellerURL string
	ResellerID  string
	Password    string
}

// CreateSession will create a Enom sessions for services
func CreateSession(resellerURL string, resellerID string, password string) Session {
	return Session{
		ResellerURL: resellerURL,
		ResellerID:  resellerID,
		Password:    password,
	}
}
