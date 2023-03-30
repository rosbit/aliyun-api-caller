package sslcert

import (
	"aliyun-api-caller/caller"
	"fmt"
	"os"
	"flag"
)

var sslcertCallerParams = caller.CallerParams {
	Method: "GET",
	Version: "2020-04-07",
	SignatureMethod: "HMAC-SHA1",
	Endpoint: "https://cas.aliyuncs.com",
}

// UploadUserCertificate - 上传证书
// 参考文档：https://help.aliyun.com/document_detail/465111.html#api-detail-22
func UploadUserCertificate() {
	f := flag.NewFlagSet("upload-user-ssl-certs", flag.ExitOnError)
	certName := f.String("cert-name", "", "specify cert name")
	debug := f.Bool("debug", false, "debug mode or not")
	certFileName := f.String("cert-file-name", "", "specify cert file name")
	keyFileName := f.String("key-file-name", "", "specify key file name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*certFileName) == 0 || len(*keyFileName) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s upload-user-ssl-certs -cert-name=xxx -cert-file-name=xxx -key-file-name=xxx -debug=true|false\n", os.Args[0])
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
		"Action": "UploadUserCertificate",
		"Cert": string(cert),
		"Key": string(key),
	}
	if len(*certName) > 0 {
		params["Name"] = *certName
	}

	var res struct {
		CertId uint64
		RequestId string
	}

	if err := caller.CallAliyun(&sslcertCallerParams, params, &res, *debug); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	fmt.Printf("%d\n", res.CertId)
}
