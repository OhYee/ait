package server

// LogCallbackType callback when output log
type LogCallbackType func(format string, args ...interface{})

// DebugCallbackType callback when output debug data
type DebugCallbackType func(format string, args ...interface{})

// ErrCallbackType callback when error occurs
type ErrCallbackType func(err error)

// Singleton Pattern
var (
	logCallback   = func(format string, args ...interface{}) {}
	debugCallback = func(format string, args ...interface{}) {}
	errCallback   = func(err error) {}
)

// SetLogCallback set callback when output log
func SetLogCallback(callback LogCallbackType) {
	logCallback = callback
}

// SetDebugCallback set callback when output log
func SetDebugCallback(callback DebugCallbackType) {
	debugCallback = callback
}

// SetErrCallback set callback when error occurs
func SetErrCallback(callback ErrCallbackType) {
	errCallback = callback
}

// Log make a log
func Log(format string, args ...interface{}) {
	logCallback(format, args...)
}

// Debug make a debug output
func Debug(format string, args ...interface{}) {
	debugCallback(format, args...)
}

// Err make a error output
func Err(err error) {
	errCallback(err)
}
