www-redirect
============

Redirect utility for k8s clusters

Usage:
```
PORT=80 FORMAT="https://www.%s%s" ./redirect
```
First `%s` is Host, the other one is `URI`.
Scheme should be set explicitly.

Example:
```
curl -v localhost:7811/foobar
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 7811 (#0)
> GET /foobar HTTP/1.1
> Host: localhost:7811
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Location: https://www.localhost:7811/foobar
< Date: Wed, 19 Jul 2017 11:44:53 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
```
