# Excercise 1.10

## Description
Find a website with lots of data and compare the results of subsequent runs. For investigation
write the content of the site to a file.

## Implementation 

Change the worker to persist the data in tmp using the url as a prefix
for a tempfile in /tmp. Disk latency is ignored.

## Example
```
$ go run main.go  https://heise.de https://google.de https://spiegel.de https://theregister.co.uk
```

# Exercise 1.11

## Description
Try longer argument list with top 100 web sites (alexa.com).
What happens if a web site does not respond?

## Implementation

Use 1.10 example and use different urls