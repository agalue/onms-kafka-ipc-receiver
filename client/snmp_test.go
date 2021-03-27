// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestSnmpV1(t *testing.T) {
	var data = []byte(`<trap-message-log location="Local" system-id="agalue-mbp.local" trap-address="192.168.105.1">
   <messages>
      <agent-address>127.0.0.1</agent-address>
      <community>public</community>
      <version>v1</version>
      <timestamp>0</timestamp>
      <creation-time>1604427231000</creation-time>
      <pdu-length>5</pdu-length>
      <trap-identity enterprise-id=".1.3.6.1.4.1.674.10892.5.3.2.1" generic="6" specific="2177"></trap-identity>
      <results>
         <result>
            <base>.1.3.6.1.4.1.674.10892.5.3.1.2.0</base>
            <value type="4">c29tZXRoaW5nIHdlbnQgd3Jvbmcu</value>
         </result>
         <result>
            <base>.1.3.6.1.4.1.674.10892.5.3.1.3.0</base>
            <value type="2">BQ==</value>
         </result>
      </results>
   </messages>
</trap-message-log>`)
	trap := new(TrapLogDTO)
	err := xml.Unmarshal(data, trap)
	assert.NilError(t, err)
	text := trap.String()

	fmt.Println(text)
	assert.Equal(t, "agalue-mbp.local", trap.SystemID)
	assert.Equal(t, 1, len(trap.Messages))
	assert.Assert(t, strings.Contains(text, `"value": "something went wrong."`))
	assert.Assert(t, strings.Contains(text, `"value": "5"`))
}

func TestSnmpV2(t *testing.T) {
	var data = []byte(`<trap-message-log system-id="docker-minion" location="Docker" trap-address="172.18.0.1">
   <messages>
      <agent-address>172.18.0.1</agent-address>
      <community>public</community>
      <version>v2</version>
      <timestamp>22318014</timestamp>
      <pdu-length>6</pdu-length>
      <creation-time>1615470557029</creation-time>
      <trap-identity generic="6" specific="2" enterprise-id=".1.3.6.1.4.1.9.9.171.2"/>
      <results>
         <result>
            <base>.1.3.6.1.6.3.18.1.3.0</base>
            <value type="64">CwAABA==</value>
         </result>
         <result>
            <base>.1.3.6.1.4.1.9.9.171.1.2.2.1.6</base>
            <value type="4">0EewAg==</value>
         </result>
         <result>
            <base>.1.3.6.1.4.1.9.9.171.1.2.2.1.7</base>
            <value type="4">0EewAg==</value>
         </result>
         <result>
            <base>.1.3.6.1.4.1.9.9.171.1.2.3.1.16</base>
            <value type="2">Cg==</value>
         </result>
      </results>
   </messages>
</trap-message-log>`)
	trap := new(TrapLogDTO)
	err := xml.Unmarshal(data, trap)
	assert.NilError(t, err)
	text := trap.String()

	fmt.Println(text)
}
