1.9 Modify fetch to also print the HTTP status code, found in resp.Status

In order to not add the status code to the output is is sent to Stderr.

```
// Always print the status code to Stderr
fmt.Fprintf(os.Stderr, "http status code: %d\n", resp.StatusCode)
```

Test:

'$ go run main.go http://gopl.io >> /dev/null'
