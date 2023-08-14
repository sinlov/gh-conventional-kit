package constant

const (
	NameCliDryRun = "dry-run"

	// EnvKeyCliVerbose
	//	Provides the debug flag. This value is true when the command is open debug mode
	EnvKeyCliVerbose = "CLI_VERBOSE"
	// NameCliVerbose
	//	Provides the debug flag. This value is true when the command is open debug mode
	NameCliVerbose = "verbose"

	// EnvKeyCliTimeoutSecond
	//	Provides the timeout second flag
	EnvKeyCliTimeoutSecond = "CLI_CONFIG_TIMEOUT_SECOND"
	// NameCliTimeoutSecond
	//	Provides the timeout second flag
	NameCliTimeoutSecond = "config.timeout_second"

	// EnvLogLevel
	//	env ENV_WEB_LOG_LEVEL default ""
	EnvLogLevel string = "CLI_LOG_LEVEL"
	// NameLogLevel
	//	Provides the log level flag
	NameLogLevel = "config.log_level"

	// NameCliRunPath
	// 	Provides the cwd path flag
	NameCliRunPath = "config.cli_run_root_path"
)
