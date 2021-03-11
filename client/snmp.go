// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math/big"
	"strings"
)

// SNMPValueDTO represents an SNMP value
type SNMPValueDTO struct {
	XMLName xml.Name `xml:"value" json:"-"`
	Type    int      `xml:"type,attr" json:"type"`
	Value   string   `xml:",chardata" json:"content"`
}

func (dto SNMPValueDTO) isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

// MarshalJSON converts SNMP Value to JSON
func (dto SNMPValueDTO) MarshalJSON() ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(dto.Value))
	var content string = string(dto.Value)
	if dto.Type == 4 || dto.Type == 6 { // OCTETSTRING(4), OID(6)
		if err == nil {
			content = string(data)
			if !dto.isASCII(content) { // TODO parse hex-content
				content = fmt.Sprintf("0x%x", string(data))
			}
		} else {
			log.Printf("[error] cannot decode base64 value: %v", err)
		}
	} else if dto.Type == 64 { // IPADDRESS
		elements := make([]string, len(data))
		for i, n := range data {
			elements[i] = new(big.Int).SetBytes([]byte{n}).String()
		}
		content = strings.Join(elements, ".")
	} else { // Assume numeric value
		content = new(big.Int).SetBytes(data).String()
	}
	return []byte(fmt.Sprintf(`{"type": %d, "value": "%s"}`, dto.Type, content)), nil
}

// SNMPResultDTO represents an SNMP result
type SNMPResultDTO struct {
	XMLName  xml.Name     `xml:"result" json:"-"`
	Base     string       `xml:"base" json:"base"`
	Instance string       `xml:"instance,omitempty" json:"instance,omitempty"`
	Value    SNMPValueDTO `xml:"value" json:"value"`
}

// SNMPResults represents a collection of SNMP result instances
type SNMPResults struct {
	Results []SNMPResultDTO `xml:"result" json:"varbinds"`
}

// TrapIdentityDTO represents the SNMP Trap Identity
type TrapIdentityDTO struct {
	EnterpriseID string `xml:"enterprise-id,attr" json:"enterpriseID"`
	Generic      int    `xml:"generic,attr" json:"generic"`
	Specific     int    `xml:"specific,attr" json:"specific"`
}

// TrapDTO represents an SNMP Trap
type TrapDTO struct {
	AgentAddress string           `xml:"agent-address" json:"agentAddress"`
	Community    string           `xml:"community" json:"community,omitempty"`
	Version      string           `xml:"version" json:"version"`
	Timestamp    int64            `xml:"timestamp" json:"timestamp"`
	CreationTime int64            `xml:"creation-time" json:"creationTime"`
	PDULength    int              `xml:"pdu-length" json:"pduLength"`
	RawMessage   []byte           `xml:"raw-message,omitempty" json:"rawMessage,omitempty"`
	TrapIdentity *TrapIdentityDTO `xml:"trap-identity" json:"trapIdentity"`
	Results      *SNMPResults     `xml:"results" json:"results"`
}

// TrapLogDTO represents a collection of SNMP Trap messages
type TrapLogDTO struct {
	XMLName     xml.Name  `xml:"trap-message-log" json:"-"`
	Location    string    `xml:"location,attr" json:"location"`
	SystemID    string    `xml:"system-id,attr" json:"systemId"`
	TrapAddress string    `xml:"trap-address,attr" json:"trapAddress"`
	Messages    []TrapDTO `xml:"messages" json:"messages"`
}

func (dto TrapLogDTO) String() string {
	bytes, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		log.Printf("[error] cannot generate JSON for SNMP trap: %v", err)
	}
	return string(bytes)
}
