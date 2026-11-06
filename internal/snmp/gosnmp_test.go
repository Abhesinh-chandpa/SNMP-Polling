package snmp

import (
 "testing"
 "time"
)

func TestOpenValidation(t *testing.T) {
 cases:=[]Config{
  {},
  {Host:"127.0.0.1",Version:"1"},
  {Host:"127.0.0.1",Version:"2c"},
  {Host:"127.0.0.1",Version:"3"},
  {Host:"127.0.0.1",Version:"3",Username:"u",PrivPassword:"secret"},
  {Host:"127.0.0.1",Version:"3",Username:"u",AuthPassword:"secret",AuthProtocol:"MD5"},
 }
 for _,cfg:=range cases {
  cfg.Timeout=time.Millisecond
  if c,err:=Open(cfg);err==nil { c.Close();t.Errorf("expected invalid config: %+v",cfg) }
 }
}
