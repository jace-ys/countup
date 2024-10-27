// Code generated by goa v3.19.1, DO NOT EDIT.
//
// countup gRPC client CLI support package
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package cli

import (
	"flag"
	"fmt"
	"os"

	apic "github.com/jace-ys/countup/api/v1/gen/grpc/api/client"
	teapotc "github.com/jace-ys/countup/api/v1/gen/grpc/teapot/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `api (auth-token|counter-get|counter-increment)
teapot echo
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` api auth-token --message '{
      "access_token": "Et eum.",
      "provider": "google"
   }'` + "\n" +
		os.Args[0] + ` teapot echo --message '{
      "text": "Quidem inventore."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		apiFlags = flag.NewFlagSet("api", flag.ContinueOnError)

		apiAuthTokenFlags       = flag.NewFlagSet("auth-token", flag.ExitOnError)
		apiAuthTokenMessageFlag = apiAuthTokenFlags.String("message", "", "")

		apiCounterGetFlags     = flag.NewFlagSet("counter-get", flag.ExitOnError)
		apiCounterGetTokenFlag = apiCounterGetFlags.String("token", "REQUIRED", "")

		apiCounterIncrementFlags       = flag.NewFlagSet("counter-increment", flag.ExitOnError)
		apiCounterIncrementMessageFlag = apiCounterIncrementFlags.String("message", "", "")
		apiCounterIncrementTokenFlag   = apiCounterIncrementFlags.String("token", "REQUIRED", "")

		teapotFlags = flag.NewFlagSet("teapot", flag.ContinueOnError)

		teapotEchoFlags       = flag.NewFlagSet("echo", flag.ExitOnError)
		teapotEchoMessageFlag = teapotEchoFlags.String("message", "", "")
	)
	apiFlags.Usage = apiUsage
	apiAuthTokenFlags.Usage = apiAuthTokenUsage
	apiCounterGetFlags.Usage = apiCounterGetUsage
	apiCounterIncrementFlags.Usage = apiCounterIncrementUsage

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
			c := apic.NewClient(cc, opts...)
			switch epn {
			case "auth-token":
				endpoint = c.AuthToken()
				data, err = apic.BuildAuthTokenPayload(*apiAuthTokenMessageFlag)
			case "counter-get":
				endpoint = c.CounterGet()
				data, err = apic.BuildCounterGetPayload(*apiCounterGetTokenFlag)
			case "counter-increment":
				endpoint = c.CounterIncrement()
				data, err = apic.BuildCounterIncrementPayload(*apiCounterIncrementMessageFlag, *apiCounterIncrementTokenFlag)
			}
		case "teapot":
			c := teapotc.NewClient(cc, opts...)
			switch epn {
			case "echo":
				endpoint = c.Echo()
				data, err = teapotc.BuildEchoPayload(*teapotEchoMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
} // apiUsage displays the usage of the api command and its subcommands.
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api auth-token -message JSON

AuthToken implements AuthToken.
    -message JSON: 

Example:
    %[1]s api auth-token --message '{
      "access_token": "Et eum.",
      "provider": "google"
   }'
`, os.Args[0])
}

func apiCounterGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api counter-get -token STRING

CounterGet implements CounterGet.
    -token STRING: 

Example:
    %[1]s api counter-get --token "Qui voluptatum omnis."
`, os.Args[0])
}

func apiCounterIncrementUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] api counter-increment -message JSON -token STRING

CounterIncrement implements CounterIncrement.
    -message JSON: 
    -token STRING: 

Example:
    %[1]s api counter-increment --message '{
      "user": "Voluptates voluptas eius cumque maxime dolore."
   }' --token "Minima dolorem id."
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] teapot echo -message JSON

Echo implements Echo.
    -message JSON: 

Example:
    %[1]s teapot echo --message '{
      "text": "Quidem inventore."
   }'
`, os.Args[0])
}
