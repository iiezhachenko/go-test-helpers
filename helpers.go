package go_test_helpers

import (
	"os/exec"
	"log"
	"io/ioutil"
)

// Runs command, captures it's STDOUT and returns it as String.
//
// The returned error is nil if the command runs, has no problems copying stdout,
// and exits with a zero exit status.
//
// If the command fails to run or doesn't complete successfully, the error is of type *ExitError.
// Other error types may be returned for I/O problems.
func CmdStdout(cwd string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}
	output, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return string(output), nil
}

// Runs command, captures it's STDERR and returns it as String.
//
// The returned error is nil if the command runs, has no problems copying stdout,
// and exits with a zero exit status.
//
// If the command fails to run or doesn't complete successfully, the error is of type *ExitError.
// Other error types may be returned for I/O problems.
func CmdStderr(cwd string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	output, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return string(output), nil
}

// Runs command, captures it's STDOUR and STDERR and returns them as combined String.
//
// The returned error is nil if the command runs, has no problems copying stdout, stderr
// and exits with a zero exit status.
//
// If the command fails to run or doesn't complete successfully, the error is of type *ExitError.
// Other error types may be returned for I/O problems.
func CmdCombOut(cwd string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd

	combout, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(combout), nil
}

// Runs "go build" and "go install" in specified folder.
//
// If the command fails to run or doesn't complete successfully, error, STDOUT and STDERR are printed.
func BuildAppForTest(cwd string) {
	out, err := CmdCombOut(cwd, "go", "build")
	if err != nil {
		log.Printf("Command \"go build\" failed with error: \"%s\"\n", err)
		log.Fatal(out)
	}

	out = ""
	err = nil

	out, err = CmdCombOut(cwd, "go", "install")
	if err != nil {
		log.Printf("Command \"go install\" failed with error: \"%s\"\n", err)
		log.Fatal(out)
	}
}
