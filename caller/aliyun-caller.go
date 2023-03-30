package caller

import (
	"github.com/rosbit/aliyun-sign"
	"github.com/rosbit/gnet"
	"aliyun-api-caller/conf"
	"net/http"
	"fmt"
	"os"
)

type CallerParams struct {
	Method string
	Version string
	SignatureMethod string
	Endpoint string
}

func CallAliyun(callerParams *CallerParams, apiParams map[string]string, res interface{}, debug bool) error {
	ak := &conf.ServiceConf.AccessKey
	params := aysign.CreateParamsWithSignature(ak.AccessKeyId, ak.AccessSecret, callerParams.Method, aysign.CommonParams{
			Format: "JSON",
			Version: callerParams.Version,
			SignatureMethod: callerParams.SignatureMethod,
		}, apiParams, debug,
	)

	o := func(*gnet.Options){}
	if debug {
		o = gnet.BodyLogger(os.Stderr)
	}
	status, err := gnet.HttpCallJ(callerParams.Endpoint, res, gnet.Params(params), o)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return fmt.Errorf("status: %d\n", status)
	}
	return nil
}

