# conditionally-execute-go
conditionally-execute-go is Go port of [conditionally-execute](https://github.com/bopke/conditionally-execute) by [@bopke](https://github.com/bopke).
It lets you abandon `if` keyword
### Usage:
```go
condition := false
ce := condexec.New(condition).OnFalse(func() {
    log.Println("Hey, condition is false!")
}).OnTrue(func() {
    log.Println("Hey, condition is true!")
}).Execute()
```
Code above will obviously print "Hey, condition is false!".
