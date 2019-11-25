package errors

// Error uses to wrap the error happened in soupbintcpclient.
type Error struct {
	appName string
	oriErr  error

	//interface
	error
}

// New creates soupbintcpClient error.
func New(err error) error {
	if err == nil {
		return nil
	}
	return &Error{appName: "SoupBinTcpClient", oriErr: err}
}

// Error construct the error message with original error and app name.
func (err *Error) Error() string {
	return "Error of " + err.appName + ": " + err.oriErr.Error()
}
