package client

import (
	"encoding/xml"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSnmp(t *testing.T) {
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
	fmt.Println(trap.String())
}
