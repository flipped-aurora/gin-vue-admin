package service

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

type CertInfo struct {
	Brand          string
	StartAt        *time.Time
	ExpireAt       *time.Time
	ValidityPeriod int
	DaysRemaining  int
	TlsVersion     string
	Fingerprint    string
	IsWildcard     bool
	SANs           []string
}

type DomainInfo struct {
	StartAt       *time.Time
	ExpireAt      *time.Time
	DaysRemaining int
}

func (s *certCertificate) ProbeDomainWhois(domain string) (*DomainInfo, error) {
	// 简单的 WHOIS 探测实现，优先支持 .com 和 .cn
	// 实际生产环境建议使用专门的 WHOIS 解析库
	whoisServer := "whois.iana.org"
	if strings.HasSuffix(domain, ".com") {
		whoisServer = "whois.verisign-grs.com"
	} else if strings.HasSuffix(domain, ".cn") {
		whoisServer = "whois.cnnic.cn"
	}

	conn, err := net.DialTimeout("tcp", whoisServer+":43", 5*time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))
	var res []byte
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			res = append(res, buf[:n]...)
		}
		if err != nil {
			break
		}
	}

	output := string(res)
	info := &DomainInfo{}

	// 解析逻辑（针对常见格式）
	// .com: "Creation Date: 2023-01-01T..." / "Registry Expiry Date: 2024-01-01T..."
	// .cn: "Registration Time: 2023-01-01..." / "Expiration Time: 2024-01-01..."

	patterns := map[string]string{
		"start":  `(?i)(Creation Date|Registration Time|registered):\s*([^\r\n]+)`,
		"expire": `(?i)(Expiry Date|Expiration Time|Registry Expiry Date|expire):\s*([^\r\n]+)`,
	}

	for key, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(output)
		if len(matches) > 2 {
			dateStr := strings.TrimSpace(matches[2])
			// 尝试多种格式解析
			layouts := []string{
				"2006-01-02T15:04:05Z",
				"2006-01-02 15:04:05",
				"2006-01-02",
				time.RFC3339,
			}
			for _, layout := range layouts {
				t, err := time.Parse(layout, dateStr)
				if err == nil {
					timeVal := t
					if key == "start" {
						info.StartAt = &timeVal
					} else {
						info.ExpireAt = &timeVal
					}
					break
				}
			}
		}
	}

	if info.ExpireAt != nil && !info.ExpireAt.IsZero() {
		info.DaysRemaining = int(time.Until(*info.ExpireAt).Hours() / 24)
		return info, nil
	}

	return nil, fmt.Errorf("WHOIS 解析失败")
}

func (s *certCertificate) ProbeCertificate(domain string) (*CertInfo, error) {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         domain,
	})

	if err != nil {
		return nil, fmt.Errorf("TLS拨号失败: %v", err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return nil, fmt.Errorf("未找到对端证书")
	}

	mainCert := state.PeerCertificates[0]

	validityPeriod := int(mainCert.NotAfter.Sub(mainCert.NotBefore).Hours() / 24)
	daysRemaining := int(time.Until(mainCert.NotAfter).Hours() / 24)

	tlsVerMap := map[uint16]string{
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}

	tlsVersion := tlsVerMap[state.Version]
	if tlsVersion == "" {
		tlsVersion = "Unknown"
	}

	brand := "Unknown"
	if len(mainCert.Issuer.Organization) > 0 {
		brand = mainCert.Issuer.Organization[0]
	}

	startAt := mainCert.NotBefore
	expireAt := mainCert.NotAfter

	certBytes := mainCert.Raw
	fingerprint := ""
	if len(certBytes) > 0 {
		hash := sha256.Sum256(certBytes)
		fingerprint = hex.EncodeToString(hash[:])
	}

	isWildcard := false
	sans := []string{}
	if mainCert.DNSNames != nil {
		sans = mainCert.DNSNames
		for _, san := range sans {
			if strings.HasPrefix(san, "*.") {
				isWildcard = true
				break
			}
		}
	}

	return &CertInfo{
		Brand:          brand,
		StartAt:        &startAt,
		ExpireAt:       &expireAt,
		ValidityPeriod: validityPeriod,
		DaysRemaining:  daysRemaining,
		TlsVersion:     tlsVersion,
		Fingerprint:    fingerprint,
		IsWildcard:     isWildcard,
		SANs:           sans,
	}, nil
}
