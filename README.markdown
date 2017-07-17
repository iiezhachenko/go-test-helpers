# go_test_helpers
--
    import "github.com/iiezhachenko/go-test-helpers"


## Usage

#### func  BuildAppForTest

```go
func BuildAppForTest(cwd string)
```
Runs "go build" and "go install" in specified folder.

If the command fails to run or doesn't complete successfully, error, STDOUT and
STDERR are printed.

#### func  CmdCombOut

```go
func CmdCombOut(cwd string, command string, args ...string) (string, error)
```
Runs command, captures it's STDOUR and STDERR and returns them as combined
String.

The returned error is nil if the command runs, has no problems copying stdout,
stderr and exits with a zero exit status.

If the command fails to run or doesn't complete successfully, the error is of
type *ExitError. Other error types may be returned for I/O problems.

#### func  CmdStderr

```go
func CmdStderr(command string) (string, error)
```
Runs command, captures it's STDERR and returns it as String.

The returned error is nil if the command runs, has no problems copying stdout,
and exits with a zero exit status.

If the command fails to run or doesn't complete successfully, the error is of
type *ExitError. Other error types may be returned for I/O problems.

#### func  CmdStdout

```go
func CmdStdout(command string) (string, error)
```
Runs command, captures it's STDOUT and returns it as String.

The returned error is nil if the command runs, has no problems copying stdout,
and exits with a zero exit status.

If the command fails to run or doesn't complete successfully, the error is of
type *ExitError. Other error types may be returned for I/O problems.
