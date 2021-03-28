package enom

import (
	"log"
	"net/url"
)

type Command struct {
	Name string
	Args url.Values
}

func (s Session) CreateCommand(commandName string) *Command {
	command := &Command{Name: commandName}
	command.SetDefaultArgs(&s)

	return command
}

func (c *Command) ToParams() string {
	return c.Args.Encode()
}

func (c *Command) SetDefaultArgs(session *Session) {
	u, err := url.Parse(session.ResellerURL)
	if err != nil {
		log.Fatal(err)
	}
	c.Args = u.Query()
	c.Args.Set("Command", c.Name)
	c.Args.Set("UID", session.ResellerID)
	c.Args.Set("PW", session.Password)
	c.Args.Set("ResponseType", "XML")
}

func (c *Command) AddParam(key string, value string) {
	c.Args.Set(key, value)
}
