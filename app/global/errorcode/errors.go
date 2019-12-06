package errorcode

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

type apiErrorMethod struct {
	StackTrace   string `json:"stack_trace"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
	TraceID      string `json:"trace_id"`
	ErrorLog     string `json:"error_log"`
}

func ApiError() Error {
	return &apiErrorMethod{}
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error interface {
	Error() string
	SetErrorCode(code string, err error)
	//SetStackTrace(err error, errorCode string)
	SetStackTrace()
	GetErrorCode() int
	GetErrorMessage() string
	GetErrorLog() string
	GetStackTrace() string
}

func (m *apiErrorMethod) Error() string {
	return fmt.Sprintf("%d: %v", m.Code, m.ErrorMessage)
}

//// SetStackTrace set stack trace
//func (e *apiError) SetStackTrace(err error, errorCode string) {
//	e.StackTrace = errors.Wrap(err, errorCode)
//}

// SetErrorCode 設定 errorcode
func (e *apiErrorMethod) SetErrorCode(code string, err error) {
	api, ok := errorCode[code]
	if !ok {
		e.Code = 9999
		e.ErrorMessage = fmt.Sprintf("Undefined Error (%s)", code)
		e.ErrorLog = err.Error()
	} else {
		e.Code = api.ErrorCode
		e.ErrorMessage = fmt.Sprintf(api.ErrorMsg+"(%d)", api.ErrorMsg)
		e.ErrorLog = err.Error()
	}
}

// get stack trace
func (e *apiErrorMethod) SetStackTrace() {
	stackBuf := make([]uintptr, 100)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = trace + fmt.Sprintf("\n\tFile: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}
	e.StackTrace = trace

}

func (e *apiErrorMethod) GetErrorCode() int {
	return e.Code
}

func (e *apiErrorMethod) GetErrorMessage() string {
	return e.ErrorMessage
}

func (e *apiErrorMethod) GetErrorLog() string {
	return e.ErrorLog
}

func (e *apiErrorMethod) GetStackTrace() string {
	return e.StackTrace
}

func WrapError(err error) error {
	return errors.Wrap(err, "failed")
}

// ErrorHandler return interface type
func ErrorHandler(errorCode string, errMsg error) Error {
	apiErr := ApiError()
	apiErr.SetErrorCode(errorCode, errMsg)
	apiErr.SetStackTrace()

	return apiErr
	//return ApiError{ErrorMessage: err, Code: code, StackTrace: getStackTrace()}
}

type APIResult struct {
	Error   ErrorMessage   `json:"error"`
	Success SuccessMessage `json:"success"`
}

// APIResult 回傳API格式
type ErrorMessage struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	StackTrace string `json:"stack_trace"`
}

type SuccessMessage struct {
	Result interface{} `json:"result"`
}

func Failed(err Error) *APIResult {
	res := &APIResult{Error: ErrorMessage{
		Code:       err.GetErrorCode(),
		Message:    err.GetErrorMessage(),
		Log:        err.GetErrorLog(),
		StackTrace: err.GetStackTrace(),
	}}
	res.Success.Result = ""
	return res
}

func Success(result interface{}) *APIResult {
	res := &APIResult{Error: ErrorMessage{
		Code:       1,
		Message:    "NO ERROR",
		Log:        "",
		StackTrace: "",
	}}
	res.Success.Result = result
	return res

}
