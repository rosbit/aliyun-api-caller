package dcdn

import (
	"aliyun-api-caller/caller"
	"fmt"
	"os"
	"flag"
)

var dcdnCallerParams = caller.CallerParams {
	Method: "GET",
	Version: "2018-01-15",
	SignatureMethod: "HMAC-SHA1",
	Endpoint: "https://DCDN.aliyuncs.com",
}

// SetDcdnDomainCertificate: 设置加速域名的证书功能
// 参考文档：https://help.aliyun.com/document_detail/131404.html
func SetByCertName() {
	f := flag.NewFlagSet("set-dcdn-certs-by-certname", flag.ExitOnError)
	domain := f.String("domain", "", "specify domain name")
	certName := f.String("cert-name", "", "specify cert name")
	debug := f.Bool("debug", false, "debug mode or not")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*domain) == 0 || len(*certName) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s set-dcdn-certs-by-certname -domain=xxx -cert-name=xxx -debug=true|false\n", os.Args[0])
		os.Exit(4)
	}

	var res struct {
		RequestId string
	}

	if err := caller.CallAliyun(&dcdnCallerParams, map[string]string{
		"Action": "SetDcdnDomainCertificate",
		"DomainName": *domain,
		"SSLProtocol": "on",
		"CertName": *certName,
		"CertType": "cas",
	}, &res, *debug); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	// fmt.Printf("OK\n")
}

// SetDcdnDomainCertificate: 上传并设置加速域名的证书功能
// 参考文档：https://help.aliyun.com/document_detail/131404.html
// [注] 本功能可以拆解为 sslcert.UploadUserCertificate() + dcdn.SetByCertName()，无需单独使用
func SetByUploadingCerts() {
	f := flag.NewFlagSet("set-dcdn-certs-by-uploading-certs", flag.ExitOnError)
	domain := f.String("domain", "", "specify domain name")
	certName := f.String("cert-name", "", "specify cert name")
	debug := f.Bool("debug", false, "debug mode or not")
	certFileName := f.String("cert-file-name", "", "specify cert file name")
	keyFileName := f.String("key-file-name", "", "specify key file name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*domain) == 0 || len(*certFileName) == 0 || len(*keyFileName) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s set-dcdn-certs-by-uploading-certs -domain=xxx -cert-name=xxx -cert-file-name=xxx -key-file-name=xxx -debug=true|false\n", os.Args[0])
		os.Exit(4)
	}
	cert, err := os.ReadFile(*certFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	key, err := os.ReadFile(*keyFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(6)
	}

	params := map[string]string{
		"Action": "SetDcdnDomainCertificate",
		"DomainName": *domain,
		"SSLProtocol": "on",
		"SSLPub": string(cert),
		"SSLPri": string(key),
		"CertType": "upload",
	}
	if len(*certName) > 0 {
		params["CertName"] = *certName
		params["ForceSet"] = "1"
	}

	var res struct {
		RequestId string
	}

	if err := caller.CallAliyun(&dcdnCallerParams, params, &res, *debug); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	// fmt.Printf("OK\n")
}
