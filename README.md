# Simple Static File Server with Basic Auth and Source IP Whitelisting

Why use this?
1. Basic authentication for accessing a remote filesystem over the web
2. Source IP Whitelisting

> Wait... Don't popular web servers already support this?

Yes. And much more fully featured.

>So why would I use it?

Single purpose binary and extremely simple. 

Motivation: It's raining outside, so why not.

# Options

TODO

# htpasswd file

Generate a new htpasswd file with the first user for authenticating users with the server:
`htpasswd -cB htpasswd daniel.cole`

Changes to the htpasswd file are automatically detected and reloaded. See https://github.com/abbot/go-http-auth for more detail.

htpasswd references: 
1. https://httpd.apache.org/docs/2.4/misc/password_encryptions.html
2. https://linux.die.net/man/1/htpasswd

# TODO
1. TLS 
2. Support proxy mode
3. Build + version
4. Finish options docs
5. Lots more testing + http server mocks
