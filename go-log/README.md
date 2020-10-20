zap相对logrus等组件的性能较高

go-log 对 zap进行了一次简单的封装，使用环境变量

**GOLOG_LOG_LEVEL**  对log level 进行设置，**error** 为默认的log level，

**GOLOG_LOG_FMT** 对日志格式(`json`, `nocolor`)进行指定

**GOLOG_OUTPUT** 对日志的输出源进行指定，支持多个源用 `+` 拼接，可选项 `stdout|stderr|file`

**GOLOG_FILE** 对日志文件进行指定，需要 **GOLOG_OUTPUT**  带有 `file` 

项目代码 https://github.com/ipfs/go-log

