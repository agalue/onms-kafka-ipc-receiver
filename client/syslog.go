// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

// SyslogMessageDTO represents a Syslog message
type SyslogMessageDTO struct {
	Timestamp string `xml:"timestamp,attr" json:"timestamp"`
	Content   []byte `xml:",chardata" json:"content"`
}

// MarshalJSON converts Syslog message to JSON
func (dto *SyslogMessageDTO) MarshalJSON() ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(string(dto.Content))
	content := strings.TrimSuffix(string(bytes), "\n")
	if err != nil {
		log.Printf("[error] cannot decode base64 value: %v", err)
	}
	return []byte(fmt.Sprintf(`{"timestamp": "%s", "content": "%s"}`, dto.Timestamp, content)), nil
}

// SyslogMessageLogDTO represents a collection of Syslog messages
type SyslogMessageLogDTO struct {
	XMLName       xml.Name           `xml:"syslog-message-log" json:"-"`
	SystemID      string             `xml:"system-id,attr" json:"systemId"`
	Location      string             `xml:"location,attr" json:"location"`
	SourceAddress string             `xml:"source-address,attr" json:"sourceAddress"`
	SourcePort    int                `xml:"source-port,attr" json:"sourcePort"`
	Messages      []SyslogMessageDTO `xml:"messages" json:"messages"`
}

func (dto SyslogMessageLogDTO) String() string {
	bytes, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		log.Printf("[error] cannot generate JSON for syslog message: %v", err)
	}
	return string(bytes)
}
