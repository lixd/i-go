package certs

import (
	"crypto/x509"
	"net"
	"testing"
)

const CommonName = "foo.example.com"

// TestSelfSignedCertHasSAN verifies the existing of
// a SAN on the generated self-signed certificate.
// a SAN ensures that the certificate is considered
// valid by default in go 1.15 and above, which
// turns off fallback to Common Name by default.
func TestSelfSignedCertHasSAN(t *testing.T) {
	key, err := NewPrivateKey(x509.RSA)
	if err != nil {
		t.Fatalf("rsa key failed to generate: %v", err)
	}
	selfSignedCert, err := NewSelfSignedCACert(Config{CommonName: CommonName}, key)
	if err != nil {
		t.Fatalf("self signed certificate failed to generate: %v", err)
	}
	if len(selfSignedCert.DNSNames) == 0 {
		t.Fatalf("self signed certificate has zero DNS names.")
	}
}

func TestGenerateCert(t *testing.T) {
	// 	1. Generate ca
	caKey, err := NewPrivateKey(x509.RSA)
	if err != nil {
		t.Fatalf("rsa key failed to generate: %v", err)
	}
	ca, err := NewSelfSignedCACert(Config{CommonName: "webhook.kube-system.svc"}, caKey)
	if err != nil {
		t.Fatalf("self signed certificate failed to generate: %v", err)
	}
	// 	2. Generate a self-signed certificate and key.
	c := &Config{
		CommonName:   "webhook.kube-system.svc",
		Organization: nil,
		AltNames: AltNames{
			DNSNames: []string{"webhook.kube-system.svc"},
			IPs:      []net.IP{net.ParseIP("192.168.10.89"), net.ParseIP("172.20.148.199")}},
		Usages:             []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		NotAfter:           nil,
		PublicKeyAlgorithm: 0,
	}
	cert, key, err := NewCertAndKey(ca, caKey, c)
	if err != nil {
		t.Fatalf("cert and key failed to generate: %v", err)
	}
	err = WriteCertAndKey("", "server", cert, key)
	if err != nil {
		t.Fatalf("cert and key failed to write: %v", err)
	}

	c = &Config{
		CommonName:   "webhook.kube-system.svc",
		Organization: nil,
		AltNames: AltNames{
			DNSNames: []string{"webhook.kube-system.svc"},
			IPs:      []net.IP{net.ParseIP("192.168.20.163")},
		},
		Usages:             []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		NotAfter:           nil,
		PublicKeyAlgorithm: 0,
	}
	cert, key, err = NewCertAndKey(ca, caKey, c)
	if err != nil {
		t.Fatalf("cert and key failed to generate: %v", err)
	}
	err = WriteCertAndKey("", "client", cert, key)
	if err != nil {
		t.Fatalf("cert and key failed to write: %v", err)
	}
	err = WriteCert("", "ca", ca)
	if err != nil {
		t.Fatalf("ca failed to write: %v", err)
	}
}
