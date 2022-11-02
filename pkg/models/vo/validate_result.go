package vo

type ValidateResult struct {
	WorkerSame      uint8 `json:"workerSame"`
	CareerSame      uint8 `json:"careerSame"`
	CertificateSame uint8 `json:"certificateSame"`
}
