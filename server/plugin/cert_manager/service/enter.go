package service

var ServiceGroupApp = new(serviceGroup)

type serviceGroup struct {
	CertCertificate    certCertificate
	CertCategory       certCategory
	SubdomainDiscovery subdomainDiscovery
	CertAssociation    certAssociation
	ConfigService      configService
	DomainAssetService domainAssetService
}
