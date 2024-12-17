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
	VSCodeUITheme VSCodeUITheme `yaml:"vscodeUiTheme"`
	Type          ThemeType     `yaml:"type"`
	Model         ColorsModel   `yaml:"colors"`
}

type ColorsModel struct {
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
