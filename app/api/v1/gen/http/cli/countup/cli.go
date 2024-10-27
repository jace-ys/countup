// Code generated by goa v3.19.1, DO NOT EDIT.
//
// countup HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	apic "github.com/jace-ys/countup/api/v1/gen/http/api/client"
	teapotc "github.com/jace-ys/countup/api/v1/gen/http/teapot/client"
	webc "github.com/jace-ys/countup/api/v1/gen/http/web/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `api (auth-token|counter-get|counter-increment)
web (index|another|login-google|login-google-callback|logout|session-token)
teapot echo
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` api auth-token --body '{
      "access_token": "Omnis cumque est asperiores dolorem.",
      "provider": "google"
   }'` + "\n" +
		os.Args[0] + ` web index` + "\n" +
		os.Args[0] + ` teapot echo --body '{
      "text": "Tempora repellendus."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		apiFlags = flag.NewFlagSet("api", flag.ContinueOnError)

		apiAuthTokenFlags    = flag.NewFlagSet("auth-token", flag.ExitOnError)
		apiAuthTokenBodyFlag = apiAuthTokenFlags.String("body", "REQUIRED", "")

		apiCounterGetFlags     = flag.NewFlagSet("counter-get", flag.ExitOnError)
		apiCounterGetTokenFlag = apiCounterGetFlags.String("token", "REQUIRED", "")

		apiCounterIncrementFlags     = flag.NewFlagSet("counter-increment", flag.ExitOnError)
		apiCounterIncrementBodyFlag  = apiCounterIncrementFlags.String("body", "REQUIRED", "")
		apiCounterIncrementTokenFlag = apiCounterIncrementFlags.String("token", "REQUIRED", "")

		webFlags = flag.NewFlagSet("web", flag.ContinueOnError)

		webIndexFlags = flag.NewFlagSet("index", flag.ExitOnError)

		webAnotherFlags = flag.NewFlagSet("another", flag.ExitOnError)

		webLoginGoogleFlags = flag.NewFlagSet("login-google", flag.ExitOnError)

		webLoginGoogleCallbackFlags             = flag.NewFlagSet("login-google-callback", flag.ExitOnError)
		webLoginGoogleCallbackCodeFlag          = webLoginGoogleCallbackFlags.String("code", "REQUIRED", "")
		webLoginGoogleCallbackStateFlag         = webLoginGoogleCallbackFlags.String("state", "REQUIRED", "")
		webLoginGoogleCallbackSessionCookieFlag = webLoginGoogleCallbackFlags.String("session-cookie", "REQUIRED", "")

		webLogoutFlags             = flag.NewFlagSet("logout", flag.ExitOnError)
		webLogoutSessionCookieFlag = webLogoutFlags.String("session-cookie", "REQUIRED", "")

		webSessionTokenFlags             = flag.NewFlagSet("session-token", flag.ExitOnError)
		webSessionTokenSessionCookieFlag = webSessionTokenFlags.String("session-cookie", "REQUIRED", "")

		teapotFlags = flag.NewFlagSet("teapot", flag.ContinueOnError)

		teapotEchoFlags    = flag.NewFlagSet("echo", flag.ExitOnError)
		teapotEchoBodyFlag = teapotEchoFlags.String("body", "REQUIRED", "")
	)
	apiFlags.Usage = apiUsage
	apiAuthTokenFlags.Usage = apiAuthTokenUsage
	apiCounterGetFlags.Usage = apiCounterGetUsage
	apiCounterIncrementFlags.Usage = apiCounterIncrementUsage

	webFlags.Usage = webUsage
	webIndexFlags.Usage = webIndexUsage
	webAnotherFlags.Usage = webAnotherUsage
	webLoginGoogleFlags.Usage = webLoginGoogleUsage
	webLoginGoogleCallbackFlags.Usage = webLoginGoogleCallbackUsage
	webLogoutFlags.Usage = webLogoutUsage
	webSessionTokenFlags.Usage = webSessionTokenUsage

	teapotFlags.Usage = teapotUsage
	teapotEchoFlags.Usage = teapotEchoUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "api":
			svcf = apiFlags
		case "web":
			svcf = webFlags
		case "teapot":
			svcf = teapotFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "api":
			switch epn {
			case "auth-token":
				epf = apiAuthTokenFlags

			case "counter-get":
				epf = apiCounterGetFlags

			case "counter-increment":
				epf = apiCounterIncrementFlags

			}

		case "web":
			switch epn {
			case "index":
				epf = webIndexFlags

			case "another":
				epf = webAnotherFlags

			case "login-google":
				epf = webLoginGoogleFlags

			case "login-google-callback":
				epf = webLoginGoogleCallbackFlags

			case "logout":
				epf = webLogoutFlags

			case "session-token":
				epf = webSessionTokenFlags

			}

		case "teapot":
			switch epn {
			case "echo":
				epf = teapotEchoFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "api":
			c := apic.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "auth-token":
				endpoint = c.AuthToken()
				data, err = apic.BuildAuthTokenPayload(*apiAuthTokenBodyFlag)
			case "counter-get":
				endpoint = c.CounterGet()
				data, err = apic.BuildCounterGetPayload(*apiCounterGetTokenFlag)
			case "counter-increment":
				endpoint = c.CounterIncrement()
				data, err = apic.BuildCounterIncrementPayload(*apiCounterIncrementBodyFlag, *apiCounterIncrementTokenFlag)
			}
		case "web":
			c := webc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "index":
				endpoint = c.Index()
			case "another":
				endpoint = c.Another()
			case "login-google":
				endpoint = c.LoginGoogle()
			case "login-google-callback":
				endpoint = c.LoginGoogleCallback()
				data, err = webc.BuildLoginGoogleCallbackPayload(*webLoginGoogleCallbackCodeFlag, *webLoginGoogleCallbackStateFlag, *webLoginGoogleCallbackSessionCookieFlag)
			case "logout":
				endpoint = c.Logout()
				data, err = webc.BuildLogoutPayload(*webLogoutSessionCookieFlag)
			case "session-token":
				endpoint = c.SessionToken()
				data, err = webc.BuildSessionTokenPayload(*webSessionTokenSessionCookieFlag)
			}
		case "teapot":
			c := teapotc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "echo":
				endpoint = c.Echo()
				data, err = teapotc.BuildEchoPayload(*teapotEchoBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// apiUsage displays the usage of the api command and its subcommands.
func apiUsage() {
	fmt.Fprintf(os.Stderr, `Service is the api service interface.
Usage:
    %[1]s [globalflags] api COMMAND [flags]

COMMAND:
    auth-token: AuthToken implements AuthToken.
    counter-get: CounterGet implements CounterGet.
    counter-increment: CounterIncrement implements CounterIncrement.

Additional help:
    %[1]s api COMMAND --help
`, os.Args[0])
}
func apiAuthTokenUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api auth-token -body JSON

AuthToken implements AuthToken.
    -body JSON: 

Example:
    %[1]s api auth-token --body '{
      "access_token": "Omnis cumque est asperiores dolorem.",
      "provider": "google"
   }'
`, os.Args[0])
}

func apiCounterGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api counter-get -token STRING

CounterGet implements CounterGet.
    -token STRING: 

Example:
    %[1]s api counter-get --token "Voluptatem aliquid in quaerat ut nihil."
`, os.Args[0])
}

func apiCounterIncrementUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api counter-increment -body JSON -token STRING

CounterIncrement implements CounterIncrement.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s api counter-increment --body '{
      "user": "Quae cupiditate."
   }' --token "Harum error iste ipsam."
`, os.Args[0])
}

// webUsage displays the usage of the web command and its subcommands.
func webUsage() {
	fmt.Fprintf(os.Stderr, `Service is the web service interface.
Usage:
    %[1]s [globalflags] web COMMAND [flags]

COMMAND:
    index: Index implements Index.
    another: Another implements Another.
    login-google: LoginGoogle implements LoginGoogle.
    login-google-callback: LoginGoogleCallback implements LoginGoogleCallback.
    logout: Logout implements Logout.
    session-token: SessionToken implements SessionToken.

Additional help:
    %[1]s web COMMAND --help
`, os.Args[0])
}
func webIndexUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web index

Index implements Index.

Example:
    %[1]s web index
`, os.Args[0])
}

func webAnotherUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web another

Another implements Another.

Example:
    %[1]s web another
`, os.Args[0])
}

func webLoginGoogleUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web login-google

LoginGoogle implements LoginGoogle.

Example:
    %[1]s web login-google
`, os.Args[0])
}

func webLoginGoogleCallbackUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web login-google-callback -code STRING -state STRING -session-cookie STRING

LoginGoogleCallback implements LoginGoogleCallback.
    -code STRING: 
    -state STRING: 
    -session-cookie STRING: 

Example:
    %[1]s web login-google-callback --code "Odit cum blanditiis ut." --state "Deleniti non." --session-cookie "Rerum eum dolor."
`, os.Args[0])
}

func webLogoutUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web logout -session-cookie STRING

Logout implements Logout.
    -session-cookie STRING: 

Example:
    %[1]s web logout --session-cookie "Dolorem ullam magnam."
`, os.Args[0])
}

func webSessionTokenUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] web session-token -session-cookie STRING

SessionToken implements SessionToken.
    -session-cookie STRING: 

Example:
    %[1]s web session-token --session-cookie "Rerum aut tenetur."
`, os.Args[0])
}

// teapotUsage displays the usage of the teapot command and its subcommands.
func teapotUsage() {
	fmt.Fprintf(os.Stderr, `Service is the teapot service interface.
Usage:
    %[1]s [globalflags] teapot COMMAND [flags]

COMMAND:
    echo: Echo implements Echo.

Additional help:
    %[1]s teapot COMMAND --help
`, os.Args[0])
}
func teapotEchoUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] teapot echo -body JSON

Echo implements Echo.
    -body JSON: 

Example:
    %[1]s teapot echo --body '{
      "text": "Tempora repellendus."
   }'
`, os.Args[0])
}