# Logging System Desgned For Web

## Clean

This is very clean logging system and easy to use.

## Example

```go
if logger := wlog.Mix.GetStatusLogger(Development); logger != nil {
    logger.Print("status information")
}
if logger := wlog.Mix.GetServiceLogger(Development); logger != nil {
    logger.Print("/api", "service information")
}
if logger := Mix.GetDatabaseLogger(Development); logger != nil {
    logger.Print("query", "table", "database information")
}
```

