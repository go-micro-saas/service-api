package apiutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	clientutil "github.com/go-micro-saas/service-kit/cluster_service_api"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Enum interface {
	String() string
	Number() protoreflect.EnumNumber
}

func CheckAPIResponse(response clientutil.Response, err error) *errors.Error {
	if err != nil {
		return errorpkg.FromError(err)
	}
	if e := clientutil.CheckResponseCode(response); e != nil {
		return e
	}
	return nil
}

func IsReason(err error, reason Enum) bool {
	return errorpkg.IsReason(err, reason.String())
}

// CheckResponseCode 请不要使用 err = CheckResponseCode 函数，因为这个函数返回的是 *errors.Error，而 err = CheckResponseCode 返回的是 error
// 请使用 if e := CheckResponseCode(response); e != nil {return errorpkg.WithStack(err)}
func CheckResponseCode(response clientutil.Response) *errors.Error {
	return clientutil.CheckResponseCode(response)
}

// CheckHTTPResponse 请不要使用 err = CheckHTTPResponse 函数，因为这个函数返回的是 *errors.Error，而 err = CheckHTTPResponse 返回的是 error
// 请使用 if e := CheckHTTPResponse(response); e != nil {return errorpkg.WithStack(err)}
func CheckHTTPResponse(httpCode int, response clientutil.Response) *errors.Error {
	return clientutil.CheckHTTPResponse(httpCode, response)
}

// CheckHTTPStatus 请不要使用 err = CheckHTTPStatus 函数，因为这个函数返回的是 *errors.Error，而 err = CheckHTTPStatus 返回的是 error
// 请使用 if e := CheckHTTPStatus(response); e != nil {return errorpkg.WithStack(err)}
func CheckHTTPStatus(statusCode int) *errors.Error {
	return clientutil.CheckHTTPStatus(statusCode)
}

// CheckResponseStatus 请不要使用 err = CheckResponseStatus 函数，因为这个函数返回的是 *errors.Error，而 err = CheckResponseStatus 返回的是 error
// 请使用 if e := CheckResponseStatus(response); e != nil {return errorpkg.WithStack(err)}
func CheckResponseStatus(resp clientutil.Response) *errors.Error {
	return clientutil.CheckResponseStatus(resp)
}
