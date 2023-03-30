package waf20

import (
	"aliyun-api-caller/caller"
	"fmt"
	"os"
	"flag"
)

var waf20CallerParams = caller.CallerParams {
	Method: "GET",
	Version: "2019-09-10",
	SignatureMethod: "HMAC-SHA1",
	Endpoint: "https://wafopenapi.cn-hangzhou.aliyuncs.com",
}

// CreateCertificateByCertificateId，根据证书ID为指定域名添加SSL证书
// 参考文档：https://help.aliyun.com/document_detail/160786.html
func CreateCertificateByCertificateId() {
	f := flag.NewFlagSet("set-waf2.0-domain-certs-by-id", flag.ExitOnError)
	domain := f.String("domain", "", "specify domain")
	certId := f.Uint64("cert-id", 0, "specify certifacte id")
	instanceId := f.String("instance-id", "", "specify instance id")
	debug := f.Bool("debug", false, "debug mode or not")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*domain) == 0 || *certId == 0 || len(*instanceId) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s set-waf2.0-domain-certs-by-id -domain=xxx -cert-id=xxx -instance-id=xxx -debug=true|false\n", os.Args[0])
		os.Exit(4)
	}

	params := map[string]string {
		"Action": "CreateCertificateByCertificateId",
		"Domain": *domain,
		"CertificateId": fmt.Sprintf("%d", *certId),
		"InstanceId": *instanceId,
	}

	var res struct {
		CertificateId string
	}

	if err := caller.CallAliyun(&waf20CallerParams, params, &res, *debug); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	fmt.Printf("%s\n", res.CertificateId)
}

// DescribeInstanceInfo 查询当前阿里云账号下WAF实例的详情
// 参考文档：https://help.aliyun.com/document_detail/140857.htm
func DescribeInstanceInfo() {
	f := flag.NewFlagSet("query-waf2.0-instance-id", flag.ExitOnError)
	instanceId := f.String("instance-id", "", "specify instance id")
	debug := f.Bool("debug", false, "debug mode or not")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	params := map[string]string {
		"Action": "DescribeInstanceInfo",
	}
	if len(*instanceId) > 0 {
		params["InstanceId"] = *instanceId
	}

	var res struct {
		InstanceInfo struct {
			Status int
			EndDate int64
			Version string
			RemainDay int
			Region string
			PayType int
			InDebt int
			InstanceId string
		}
	}

	if err := caller.CallAliyun(&waf20CallerParams, params, &res, *debug); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	fmt.Printf("%s\n", res.InstanceInfo.InstanceId)
}
