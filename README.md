# go-commander

## Types of flags

### Switch

A switch is a flag that can only be a boolean.
If the switch is defined, the boolean will be set to true.

Example:<br>
`go run main.go -f --logging --verbose=true`

In this example 3 forms of a switch are shown.
The short version, the long version and the value version.

### Value flag

A switch is a flag that can only be a boolean.
If the switch is defined, the boolean will be set to true.

Example:<br>
`go run main.go -c dev --configuration=dev`

In this example the short and long verson are shown of the flag.
