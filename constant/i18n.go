package constant

const (
	LangEnUS = "en-US"
	LangZhCN = "zh-CN"
)

var (
	supportLanguage = []string{
		LangEnUS,
		LangZhCN,
	}
)

func SupportLanguage() []string {
	return supportLanguage
}
