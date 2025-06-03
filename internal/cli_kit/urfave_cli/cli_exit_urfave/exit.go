package cli_exit_urfave

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// exitCodeDefault SIGUSR1 as 10.
const exitCodeDefault = 10

var exitCode = exitCodeDefault

func ChangeDefaultExitCode(code int) {
	exitCode = code
}

func Format(format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), exitCode)
}

func FormatCode(code int, format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), code)
}

func Err(err error) cli.ExitCoder {
	return cli.Exit(err.Error(), exitCode)
}

func ErrCode(code int, err error) cli.ExitCoder {
	return cli.Exit(err.Error(), code)
}

func ErrMsg(msg string) cli.ExitCoder {
	return Err(fmt.Errorf("err: %s", msg))
}

func ErrMsgf(format string, a ...any) cli.ExitCoder {
	return Err(fmt.Errorf(format, a...))
}

func ErrMsgCode(code int, msg string) cli.ExitCoder {
	return ErrCode(code, fmt.Errorf("err: %s", msg))
}

func ErrMsgCodef(code int, format string, a ...any) cli.ExitCoder {
	return ErrCode(code, fmt.Errorf(format, a...))
}
