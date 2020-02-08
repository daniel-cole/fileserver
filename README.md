# Simple Static File Server with Basic Auth and Source IP Whitelisting

Why use this?
1. Basic authentication for accessing a remote filesystem over the web
2. Source IP Whitelisting

> Wait... Don't popular web servers already support this?

Yes. And much more fully featured. Caddy, Nginx or httpd are all suitable.

>So why would I use it?

Single purpose binary and extremely simple. 

Motivation: It's raining outside, so why not.


# Command Line Arguments
| Argument        | Flag            | Description  | Default Value
| -------------|:---------------------|:-----|:-----------|
| port         | -port          | Port to listen on | 9000
| address      | -address       | Address to bind to | "0.0.0.0"
| sourceRanges | -sourceRanges  | CSV string of source ranges to allow access to the file server | "0.0.0.0/0,::/0"
| directory    | -directory     | Root directory to serve files from | "."
| htpasswdFile | -htpasswdFile  | htpasswd file to use for authenticating users | "htpasswd"
| logLevel     | -logLevel      | Set the log level INFO,WARN,ERROR,DEBUG | INFO
| tlsCertFile  | -tlsCertFile   | TLS certificate file to use. Must be specified with tlsKeyFile | ""
| tlsKeyFile   | -tlsKeyFile    | TLS key to use. Must be specified with tlsCertFile | ""

# htpasswd file

Generate a new htpasswd file with the first user for authenticating users with the server:
`htpasswd -cB htpasswd daniel.cole`

Changes to the htpasswd file are automatically detected and reloaded. See https://github.com/abbot/go-http-auth for more detail.

htpasswd references: 
1. https://httpd.apache.org/docs/2.4/misc/password_encryptions.html
2. https://linux.die.net/man/1/htpasswd

# TODO
1. Support proxy mode
2. Lots more testing + http server mocks
