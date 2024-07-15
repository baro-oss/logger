# Go Logger
## The library used for wrap multiple golang logger libraries such as [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), [logrus](https://github.com/sirupsen/logrus).
### 1. Conceptional
Logger interface base on common log instance that used in many real golang projects. <br>
Logger provides multiple level of logging:
- *__Info__*
- *__Warn__*
- *__Err__*
- *__Debug__*
- *__Fatal__*
- *__Trace__*

If some projects that need handling process by context, the logger also provides methods for handling context, tracing request.
- *__InfoWithCtx__*
- *__WarnWithCtx__*
- *__ErrWithCtx__*
- *__DebugWithCtx__*
- *__FatalWithCtx__*
- *__TraceWithCtx__*

Writing to io.Stdout, io.Stderr, and system log files decrease performance of program so some logger libraries use buffering mechanism to asynchronous writing logs to io.Out. Use *__logger.Sync()__* to retrieving logs from buffer and writing to io.Out.

### 2. Example
```go
package main

func main()  {
	logger := NewLogger(LogDriverZap, true)
	defer logger.Sync()

	logger.Info("This is an example of using logger for show log info", WithField("repo", "logger"))
}
```
Result:
```json
{"level":"info","@timestamp":1721031147.7998312,"caller":"go-logger/zap.go:47","msg":"This is an example of using logger for show log info","repo":"logger","log-driver":"zap"}
```