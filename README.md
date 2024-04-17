# HasteCat
A termbin style connector to post to haste-server

## Building
`go build -o hastecat .`

## Running
`./hastcat --port 99 --hasteurl https://haste.egglabs.net`

## Environment Variables
These variables can be used to set the configuration instead of using command line flags

`LISTEN_IP` Sets the listening IP address. Defaults to `0.0.0.0`\
`LISTEN_PORT` Sets the listening port. Defaults to `99`\
`HASTEBIN_URL` Sets the URL of the hastebin to post to. Defaults to `https://haste.egglabs.net`





