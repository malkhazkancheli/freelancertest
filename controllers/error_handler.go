package controllers

import (
	"encoding/json"
	"fmt"
	"freelancertest/models"
	"github.com/go-openapi/errors"
	"net/http"
	"reflect"
	"strings"
)

// DefaultHTTPCode is used when the error Code cannot be used as an HTTP code.
var DefaultHTTPCode = http.StatusUnprocessableEntity

// Error represents a error interface all swagger framework errors implement
type Error interface {
	error
	Code() int32
}

type apiError struct{
	problem models.ProblemDetails
}

func (a *apiError) Error() string {
	return a.problem.Detail
}

func (a *apiError) Code() int32 {
	return a.problem.Status
}

// New creates a new API error with a code and a message
func New(problem models.ProblemDetails) Error {

	return &apiError{problem: problem}
}

// NotFound creates a new not found error
func NotFound(message string, args ...interface{}) Error {
	if message == "" {
		message = "Not found"
	}

	problem:= makeNewError(http.StatusNotFound,message,"Not found").problem
	return New(problem)
}

// NotImplemented creates a new not implemented error
func NotImplemented(message string) Error {

	problem:= makeNewError(http.StatusNotImplemented,message,"Method not implemented").problem
	return New(problem)
}

// MethodNotAllowedError represents an error for when the path matches but the method doesn't
type MethodNotAllowedError apiError

func (m *MethodNotAllowedError) Error() string {
	return m.problem.Detail
}

// Code the error code
func (m *MethodNotAllowedError) Code() int32 {
	return m.problem.Status
}

func errorAsJSON(err apiError) []byte {
	b, _ := json.Marshal(err.problem)
	return b
}

func flattenComposite(errs *errors.CompositeError) *errors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *errors.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}
	return errors.CompositeValidationError(res...)
}

// MethodNotAllowed creates a new method not allowed error
func MethodNotAllowed(requested string, allow []string) *apiError {
	msg := fmt.Sprintf("method %s is not allowed, but [%s] are", requested, strings.Join(allow, ","))
	return makeNewError(http.StatusMethodNotAllowed,msg,"Method not allowed")

}

// ServeError the error handler interface implementation
func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *errors.CompositeError:
		er := flattenComposite(e)
		// strips composite errors to first element only
		if len(er.Errors) > 0 {
			ServeError(rw, r, er.Errors[0])
		} else {
			// guard against empty CompositeError (invalid construct)
			ServeError(rw, r, nil)
		}
	case Error:
		value := reflect.ValueOf(e)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write(errorAsJSON(*makeNewError(http.StatusInternalServerError,e.Error(),"Internal server error")))
			return
		}
		rw.WriteHeader(asHTTPCode(int(e.Code())))
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(*makeNewError(http.StatusInternalServerError,e.Error(),"Internal server error")))
		}
	case nil:
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write(errorAsJSON(*makeNewError(http.StatusInternalServerError,e.Error(),"Unknown error")))
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(*makeNewError(http.StatusInternalServerError,e.Error(),"Internal server error")))
		}
	}
}

func asHTTPCode(input int) int {
	if input >= 600 {
		return DefaultHTTPCode
	}
	return input
}

func makeNewError(code int, message string, title string) *apiError {
	return &apiError{
		models.ProblemDetails{
			Code: http.StatusText( code),
			Detail: message,
			Instance: "",
			Status: int32(code),
			Title: title,
			Type: "",
		},}
}
