package main

import (
	"aliyun-api-caller/conf"
	sslcert "aliyun-api-caller/ssl-cert"
	waf20 "aliyun-api-caller/waf-2.0"
	"aliyun-api-caller/dcdn"
	"fmt"
	"os"
)

func init() {
	if err := conf.CheckGlobalConf(); err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <command> <options>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Where <command> are:\n")
	fmt.Fprintf(os.Stderr, "  upload-user-ssl-certs\n")
	fmt.Fprintf(os.Stderr, "  query-waf2.0-instance-id\n")
	fmt.Fprintf(os.Stderr, "  set-waf2.0-domain-certs-by-id\n")
	fmt.Fprintf(os.Stderr, "  set-dcdn-certs-by-certname\n")
	fmt.Fprintf(os.Stderr, "  set-dcdn-certs-by-uploading-certs\n")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "upload-user-ssl-certs":
		sslcert.UploadUserCertificate()
	case "query-waf2.0-instance-id":
		waf20.DescribeInstanceInfo()
	case "set-waf2.0-domain-certs-by-id":
		waf20.CreateCertificateByCertificateId()
	case "set-dcdn-certs-by-certname":
		dcdn.SetByCertName()
	case "set-dcdn-certs-by-uploading-certs":
		dcdn.SetByUploadingCerts()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[1])
		os.Exit(2)
	}
}

