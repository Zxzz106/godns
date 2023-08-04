package provider

import (
	"github.com/Zxzz106/godns/internal/settings"
)

type IDNSProvider interface {
	Init(conf *settings.Settings)
	UpdateIP(domainName, subdomainName, ip string) error
}
