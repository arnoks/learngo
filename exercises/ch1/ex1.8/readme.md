# Excercise 1.8

Modify fetch to add the prefix http:// to each argument URL if it is
missing. You might want to use strings.HasPrefix

Implementation usinig switch statment without operand aka tagless switch:
```

switch {
	case strings.HasPrefix(url, "http://"):
		break
	case strings.HasPrefix(url, "https://"):
		break
	default:
		url = "http://" + url
	}
```


`$ go run main.go gopl.io  `

is equal to:

`$ go run main.go http://gopl.io ` 
