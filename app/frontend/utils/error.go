package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MuskHandleErr(err error) {
	if err != nil {
		hlog.Error(err)
	}

}
