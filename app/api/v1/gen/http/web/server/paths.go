// Code generated by goa v3.19.1, DO NOT EDIT.
//
// HTTP request path constructors for the web service.
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

// IndexWebPath returns the URL path to the web service Index HTTP endpoint.
func IndexWebPath() string {
	return "/"
}

// AnotherWebPath returns the URL path to the web service Another HTTP endpoint.
func AnotherWebPath() string {
	return "/another"
}

// LoginGoogleWebPath returns the URL path to the web service LoginGoogle HTTP endpoint.
func LoginGoogleWebPath() string {
	return "/login/google"
}

// LoginGoogleCallbackWebPath returns the URL path to the web service LoginGoogleCallback HTTP endpoint.
func LoginGoogleCallbackWebPath() string {
	return "/login/google/callback"
}

// LogoutWebPath returns the URL path to the web service Logout HTTP endpoint.
func LogoutWebPath() string {
	return "/logout"
}

// SessionTokenWebPath returns the URL path to the web service SessionToken HTTP endpoint.
func SessionTokenWebPath() string {
	return "/session/token"
}