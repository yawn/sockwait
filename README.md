# sockwait (`sw`)

`sw` will wait for a given number of hosts to be available for opening TCP connections and `exec`s a passed binary (with optional arguments) if successful.

It waits a configurable amount of time (`-sleep`, defaults to `100ms`) between attempts and will give up after a given amount of time (`-timeout`, defaults to `5s`).

## Example

`sw www.google.de:80 172.16.0.1:1234 -- /my/fancy/app --port 23 --verbose`

This will either

* return with `1` when neither `www.google.de` (on port `80`) nor `172.16.0.1` (on port `1234`) respond within `5s`
* `exec` my fancy `app` with the flags `--port 23` and `--verbose`.
