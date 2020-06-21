# go-commander

## Types of flags

### Switch

A switch is a flag that can only be a boolean.
If the switch is defined, the boolean will be set to true.

Example:<br>
`go run main.go --logging`

In this example 3 forms of a switch are shown.
The short version, the long version and the value version.

### Value flag

Example:<br>
`go run main.go -c dev --debugmessage=hello`

In this example the short and long verson are shown of the flag.
