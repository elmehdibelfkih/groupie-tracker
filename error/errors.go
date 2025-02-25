package error

import (
	"html/template"
	"net/http"
)

const ERROR_TMPL_PATH = "./templates/error.html"

type Error struct {
	StatusCode       int
	StatusText       string
	ErrorMessage     string
	ErrorTitle       string
	ErrorDescription string
}

var ERRORTMPL *template.Template

func NotFoundError(w http.ResponseWriter) {
	notFoundError := Error{
		StatusCode:       404,
		StatusText:       "Not Found",
		ErrorMessage:     "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		ErrorTitle:       "Oops! Page Not Found",
		ErrorDescription: "We couldn't find the page you were looking for. Please check the URL for any mistakes or go back to the homepage.",
	}
	w.WriteHeader(notFoundError.StatusCode)
	ERRORTMPL.Execute(w, notFoundError)
}

func MethodNotAllowed(w http.ResponseWriter) {
	methodNotAllowed := Error{
		StatusCode:       405,
		StatusText:       "Method not allowed",
		ErrorMessage:     "The HTTP method used for this request is not allowed on this resource. Please use a different method.",
		ErrorTitle:       "Oops! method not allowed",
		ErrorDescription: "Please send a POST request to this endpoint with the required data for processing. Other HTTP methods are not allowed.",
	}
	w.WriteHeader(methodNotAllowed.StatusCode)
	ERRORTMPL.Execute(w, methodNotAllowed)
}

func InternalServerError(w http.ResponseWriter) {
	internalServerError := Error{
		StatusCode:       500,
		StatusText:       "Internal Server Error",
		ErrorMessage:     "An unexpected error occurred while processing your request.",
		ErrorTitle:       "Oops! Internal Server Error",
		ErrorDescription: "Something went wrong on our end. Please try again later or contact support if the issue persists.",
	}
	w.WriteHeader(internalServerError.StatusCode)
	ERRORTMPL.Execute(w, internalServerError)
}

func BadRequest(w http.ResponseWriter) {
	badRequest := Error{
		StatusCode:       400,
		StatusText:       "Bad Request",
		ErrorMessage:     "The Server recieved a Bad request!!",
		ErrorTitle:       "Oops! Bad Request",
		ErrorDescription: "Please make sure you respect the input limits",
	}
	w.WriteHeader(badRequest.StatusCode)
	ERRORTMPL.Execute(w, badRequest)
}

func BadGateway(w http.ResponseWriter) {
	badGateway := Error{
		StatusCode:       502,
		StatusText:       "Bad Gateway",
		ErrorMessage:     "The server received an invalid response from the upstream server.",
		ErrorTitle:       "Oops! Bad Gateway",
		ErrorDescription: "The server encountered an issue while communicating with another service.",
	}
	w.WriteHeader(badGateway.StatusCode)
	ERRORTMPL.Execute(w, badGateway)
}
