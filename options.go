package simaqian

type options struct {
	level      level
	skip       int
	logType    logType
	stacktrace *stacktrace
}

func defaultOptions() *options {
	return &options{
		level:   LevelInfo,
		skip:    2,
		logType: TypeZap,
	}
}
