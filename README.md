# Web Log

Logging system designed for backend web application.

## Clean

This is very clean logging system and easy to use.

## Process

It is still in development period.

1> The `Panic` and `Panicf` haven't implemented.

2> The basic logger is from standard library, which will be replaced in the future.

3> The flag, option and level is still wait to design.

## Example

```go
wlog.Mix.Status(mix.Devp).Print("status information")
wlog.Mix.Service(mix.Devp).Print("/api", "service information")
wlog.Mix.Database(mix.Devp).Print("query", "table", "database information")
```
