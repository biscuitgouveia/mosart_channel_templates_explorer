package templatetypes

import (
	"encoding/xml"
	"fmt"
)

type Avconfig struct {
	XMLName          xml.Name         `xml:"avconfig"`
	Origin           string           `xml:"origin,attr"`
	Provider         string           `xml:"provider,attr"`
	ConnectionString string           `xml:"connectionString,attr"`
	ChannelTemplates ChannelTemplates `xml:"channeltemplates"`
}

func (a Avconfig) String() string {
	return fmt.Sprintf("Origin: %s, Provider: %s, ConnectionString: %s", a.Origin, a.Provider, a.ConnectionString)
}

type ChannelTemplates struct {
	XMLName  xml.Name   `xml:"channeltemplates"`
	Revision string     `xml:"revision,attr"`
	Channels []Channels `xml:"channels"`
}

func (c ChannelTemplates) String() string {
	return fmt.Sprintf("Revision: %s", c.Revision)
}

type Channels struct {
	XMLName      xml.Name `xml:"channels"`
	Name         string   `xml:"name,attr"`
	Default      string   `xml:"default,attr"`
	Created      string   `xml:"created,attr"`
	Shared       string   `xml:"shared,attr"`
	BasedOn      string   `xml:"based_on,attr"`
	HideFromUser string   `xml:"hide_from_user,attr"`
}

func (c Channels) String() string {
	return fmt.Sprintf("Name: %s, Default: %s, Created: %s, Shared: %s, BasedOn: %s, HideFromUser: %s", c.Name, c.Default, c.Created, c.Shared, c.BasedOn, c.HideFromUser)
}
