package client

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
)

// SyslogMessageDTO represents a Syslog message
type SyslogMessageDTO struct {
	Timestamp string `xml:"timestamp,attr"`
	Content   []byte `xml:",chardata"`
}

// SyslogMessageLogDTO represents a collection of Syslog messages
type SyslogMessageLogDTO struct {
	XMLName       xml.Name           `xml:"syslog-message-log"`
	SystemID      string             `xml:"system-id,attr"`
	Location      string             `xml:"location,attr"`
	SourceAddress string             `xml:"source-address,attr"`
	SourcePort    int                `xml:"source-port,attr"`
	Messages      []SyslogMessageDTO `xml:"messages"`
}

func (dto SyslogMessageLogDTO) String() string {
	jsonmap := make(map[string]interface{})
	jsonmap["systemId"] = dto.SystemID
	jsonmap["location"] = dto.Location
	jsonmap["sourceAddress"] = dto.SourceAddress
	jsonmap["sourcePort"] = dto.SourcePort
	messages := make(map[string]string)
	for _, m := range dto.Messages {
		data, _ := base64.StdEncoding.DecodeString(string(m.Content))
		messages[m.Timestamp] = string(data)
	}
	jsonmap["messages"] = messages
	bytes, _ := json.MarshalIndent(jsonmap, "", "  ")
	return string(bytes)
}
