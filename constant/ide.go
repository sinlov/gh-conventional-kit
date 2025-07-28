package constant

const (
	IDEVscode    = "vscode"
	IDEJetbrains = "jetbrains"
)

var supportIDE = []string{
	IDEVscode,
	IDEJetbrains,
}

func SupportIDE() []string {
	return supportIDE
}
