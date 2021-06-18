package entity

type LoginInfo struct {
	LoginToken string `json:"login_token"`
	Form       string `json:"json"`
}
type DdnsInfo struct {
	LoginInfo
	Domain     string `json:"domain_id"`
	RecordID   string `json:"record_id"`
	SubDomain  string `json:"sub_domain"`
	RecordLine string `json:"record_line"`
	Value      string `json:"value"`
}
