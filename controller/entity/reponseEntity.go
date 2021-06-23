package entity

import "encoding/xml"

type Dnspod struct {
	XMLName xml.Name `xml:"dnspod"`
	Status  struct {
		Code      string `xml:"code"`
		Message   string `xml:"message"`
		CreatedAt string `xml:"created_at"`
	} `xml:"status"`
	Domain struct {
		ID        string `xml:"id"`
		Name      string `xml:"name"`
		Punycode  string `xml:"punycode"`
		Grade     string `xml:"grade"`
		Owner     string `xml:"owner"`
		ExtStatus string `xml:"ext_status"`
		Ttl       string `xml:"ttl"`
		MinTtl    string `xml:"min_ttl"`
		DnspodNs  struct {
			Item []string `xml:"item"`
		} `xml:"dnspod_ns"`
		Status string `xml:"status"`
	} `xml:"domain"`
	Info struct {
		SubDomains  string `xml:"sub_domains"`
		RecordTotal string `xml:"record_total"`
		RecordsNum  string `xml:"records_num"`
	} `xml:"info"`
	Records struct {
		Item []struct {
			ID            string `xml:"id"`
			Ttl           string `xml:"ttl"`
			Value         string `xml:"value"`
			Enabled       string `xml:"enabled"`
			Status        string `xml:"status"`
			UpdatedOn     string `xml:"updated_on"`
			Name          string `xml:"name"`
			Line          string `xml:"line"`
			LineID        string `xml:"line_id"`
			Type          string `xml:"type"`
			Weight        string `xml:"weight"`
			MonitorStatus string `xml:"monitor_status"`
			Remark        string `xml:"remark"`
			UseAqb        string `xml:"use_aqb"`
			Mx            string `xml:"mx"`
		} `xml:"item"`
	} `xml:"records"`
}
