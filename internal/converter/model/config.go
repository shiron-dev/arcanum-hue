package model

type ThemeType string

const (
	ThemeTypeLight ThemeType = "light"
	ThemeTypeDark  ThemeType = "dark"
)

type Config struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Colors      []ColorTheme `yaml:"colorThemes"`
}

type ColorTheme struct {
	Name          string        `yaml:"name"`
	VSCodeUITheme VSCodeUITheme `jsonschema:"enum=vs-dark,enum=vs,enum=hc-black,enum=hc-light" yaml:"vscodeUiTheme"`
	Type          ThemeType     `jsonschema:"enum=light,enum=dark"                             yaml:"type"`
	Model         ColorsModel   `yaml:"colors"`
}

type ColorsModel struct {
	Foreground          string `yaml:"foreground"`
	HiForeground        string `yaml:"hiForeground"`
	SecondaryForeground string `yaml:"secondaryForeground"`
	ForegroundAccent    string `yaml:"foregroundAccent"`
	WarningForeground   string `yaml:"warningForeground"`
	LinkForeground      string `yaml:"linkForeground"`

	Background          string `yaml:"background"`
	EditorBackground    string `yaml:"editorBackground"`
	SecondaryBackground string `yaml:"secondaryBackground"`
	BackgroundAccent    string `yaml:"backgroundAccent"`
	HiBackgroundAccent  string `yaml:"hiBackgroundAccent"`
	MatchBackground     string `yaml:"matchBackground"`

	Border string `yaml:"border"`

	TerminalAnsi TerminalAnsiModel `yaml:"terminalAnsi"`
}

type TerminalAnsiModel struct {
	TerminalAnsiBlack         string `yaml:"black"`
	TerminalAnsiBrightBlack   string `yaml:"brightBlack"`
	TerminalAnsiBlue          string `yaml:"blue"`
	TerminalAnsiCyan          string `yaml:"cyan"`
	TerminalAnsiGreen         string `yaml:"green"`
	TerminalAnsiBrightBlue    string `yaml:"brightBlue"`
	TerminalAnsiBrightCyan    string `yaml:"brightCyan"`
	TerminalAnsiBrightGreen   string `yaml:"brightGreen"`
	TerminalAnsiWhite         string `yaml:"white"`
	TerminalAnsiBrightWhite   string `yaml:"brightWhite"`
	TerminalAnsiMagenta       string `yaml:"magenta"`
	TerminalAnsiYellow        string `yaml:"yellow"`
	TerminalAnsiRed           string `yaml:"red"`
	TerminalAnsiBrightYellow  string `yaml:"brightYellow"`
	TerminalAnsiBrightMagenta string `yaml:"brightMagenta"`
	TerminalAnsiBrightRed     string `yaml:"brightRed"`
}
