package snmp

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
)

// Config configures a read-only SNMP session. Never log this structure.
type Config struct {
	Host         string
	Port         uint16
	Version      string // "2c" or "3"
	Community    string
	Username     string
	AuthProtocol string // SHA, SHA256
	AuthPassword string
	PrivProtocol string // AES
	PrivPassword string
	Timeout      time.Duration
	Retries      int
}

// Open creates an isolated session for one device. The caller must Close it.
func Open(cfg Config) (Client, error) {
	if cfg.Host == "" {
		return nil, errors.New("SNMP host is required")
	}
	if cfg.Port == 0 {
		cfg.Port = 161
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 2 * time.Second
	}
	if cfg.Retries < 0 {
		return nil, errors.New("retries cannot be negative")
	}
	g := &gosnmp.GoSNMP{Target: cfg.Host, Port: cfg.Port, Timeout: cfg.Timeout, Retries: cfg.Retries}
	switch cfg.Version {
	case "2c":
		if cfg.Community == "" {
			return nil, errors.New("v2c community required")
		}
		g.Version = gosnmp.Version2c
		g.Community = cfg.Community
	case "3":
		if cfg.Username == "" {
			return nil, errors.New("v3 username required")
		}
		g.Version = gosnmp.Version3
		g.SecurityModel = gosnmp.UserSecurityModel
		sec := &gosnmp.UsmSecurityParameters{UserName: cfg.Username}
		if cfg.AuthPassword == "" && cfg.PrivPassword != "" {
			return nil, errors.New("privacy requires authentication")
		}
		if cfg.AuthPassword != "" {
			switch strings.ToUpper(cfg.AuthProtocol) {
			case "SHA":
				sec.AuthenticationProtocol = gosnmp.SHA
			case "SHA256":
				sec.AuthenticationProtocol = gosnmp.SHA256
			default:
				return nil, errors.New("unsupported authentication protocol")
			}
			sec.AuthenticationPassphrase = cfg.AuthPassword
			g.MsgFlags = gosnmp.AuthNoPriv
		} else {
			g.MsgFlags = gosnmp.NoAuthNoPriv
		}
		if cfg.PrivPassword != "" {
			if strings.ToUpper(cfg.PrivProtocol) != "AES" {
				return nil, errors.New("unsupported privacy protocol")
			}
			sec.PrivacyProtocol = gosnmp.AES
			sec.PrivacyPassphrase = cfg.PrivPassword
			g.MsgFlags = gosnmp.AuthPriv
		}
		g.SecurityParameters = sec
	default:
		return nil, fmt.Errorf("unsupported SNMP version %q", cfg.Version)
	}
	if err := g.Connect(); err != nil {
		return nil, fmt.Errorf("SNMP connect: %w", err)
	}
	return &goClient{session: g}, nil
}

type goClient struct {
	session *gosnmp.GoSNMP
}

func (c *goClient) Close() error {
	if c.session.Conn != nil {
		return c.session.Conn.Close()
	}
	return nil
}

func (c *goClient) Get(ctx context.Context, oids []string) ([]Variable, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if len(oids) == 0 {
		return nil, errors.New("at least one OID required")
	}
	packet, err := c.session.Get(oids)
	if err != nil {
		return nil, err
	}
	if packet.Error != gosnmp.NoError {
		return nil, fmt.Errorf("SNMP error: %s", packet.Error.String())
	}
	out := make([]Variable, 0, len(packet.Variables))
	for _, v := range packet.Variables {
		out = append(out, convert(v))
	}
	return out, nil
}

func (c *goClient) Walk(ctx context.Context, root string, visit func(Variable) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if visit == nil {
		return errors.New("visit callback required")
	}
	return c.session.BulkWalk(root, func(p gosnmp.SnmpPDU) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		return visit(convert(p))
	})
}

func convert(p gosnmp.SnmpPDU) Variable {
	v := Variable{OID: p.Name, Type: p.Type.String()}
	switch val := p.Value.(type) {
	case []byte:
		v.Value = string(val)
	case nil:
		v.Value = ""
	case uint64:
		v.Value = strconv.FormatUint(val, 10)
	default:
		v.Value = fmt.Sprint(val)
	}
	return v
}
