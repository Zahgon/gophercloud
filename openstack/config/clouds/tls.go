package clouds

import (
	"crypto/tls"
)

func computeTLSConfig(cloud Cloud, options cloudOpts) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveTilde(p string) (string, error) { _ = "STUB: not implemented"; return "", nil }
