package vo

type WorkerCertificate struct {
	ID            uint32 `json:"id"`
	CertificateId uint32 `json:"certificateId"`
	CertCode      string `json:"certCode"`
	CertName      string `json:"certName"`
}
