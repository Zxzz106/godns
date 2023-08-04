package provider

import (
	"fmt"

	"github.com/Zxzz106/godns/internal/provider/alidns"
	"github.com/Zxzz106/godns/internal/provider/cloudflare"
	"github.com/Zxzz106/godns/internal/provider/dnspod"
	"github.com/Zxzz106/godns/internal/provider/dreamhost"
	"github.com/Zxzz106/godns/internal/provider/duck"
	"github.com/Zxzz106/godns/internal/provider/dynv6"
	"github.com/Zxzz106/godns/internal/provider/google"
	"github.com/Zxzz106/godns/internal/provider/he"
	"github.com/Zxzz106/godns/internal/provider/hetzner"
	"github.com/Zxzz106/godns/internal/provider/linode"
	"github.com/Zxzz106/godns/internal/provider/loopiase"
	"github.com/Zxzz106/godns/internal/provider/noip"
	"github.com/Zxzz106/godns/internal/provider/ovh"
	"github.com/Zxzz106/godns/internal/provider/scaleway"
	"github.com/Zxzz106/godns/internal/provider/strato"
	"github.com/Zxzz106/godns/internal/settings"
	"github.com/Zxzz106/godns/internal/utils"
)

func GetProvider(conf *settings.Settings) (IDNSProvider, error) {
	var provider IDNSProvider

	switch conf.Provider {
	case utils.CLOUDFLARE:
		provider = &cloudflare.DNSProvider{}
	case utils.DNSPOD:
		provider = &dnspod.DNSProvider{}
	case utils.DREAMHOST:
		provider = &dreamhost.DNSProvider{}
	case utils.HE:
		provider = &he.DNSProvider{}
	case utils.ALIDNS:
		provider = &alidns.DNSProvider{}
	case utils.GOOGLE:
		provider = &google.DNSProvider{}
	case utils.DUCK:
		provider = &duck.DNSProvider{}
	case utils.NOIP:
		provider = &noip.DNSProvider{}
	case utils.SCALEWAY:
		provider = &scaleway.DNSProvider{}
	case utils.DYNV6:
		provider = &dynv6.DNSProvider{}
	case utils.LINODE:
		provider = &linode.DNSProvider{}
	case utils.STRATO:
		provider = &strato.DNSProvider{}
	case utils.LOOPIASE:
		provider = &loopiase.DNSProvider{}
	case utils.HETZNER:
		provider = &hetzner.DNSProvider{}
	case utils.OVH:
		provider = &ovh.DNSProvider{}
	default:
		return nil, fmt.Errorf("Unknown provider '%s'", conf.Provider)
	}

	provider.Init(conf)

	return provider, nil
}
