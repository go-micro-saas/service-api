package apiutil

import (
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
)

func CheckAPIResponse(response clientutil.Response, err error) error {
	if err != nil {
		return err
	}
	e := clientutil.CheckResponseCode(response)
	if e != nil {
		return e
	}
	return nil
}
