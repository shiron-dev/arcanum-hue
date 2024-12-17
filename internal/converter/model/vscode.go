package model

type VSCodeUITheme string

const (
	VSCodeUIThemeDark      VSCodeUITheme = "vs-dark"
	VSCodeUIThemeLight     VSCodeUITheme = "vs"
	VSCodeUIThemeHighDark  VSCodeUITheme = "hc-black"
	VSCodeUIThemeHighLight VSCodeUITheme = "hc-light"
)

type VSCodeThemeExtension struct {
	Name        string
	Description string
	Version     string
	Themes      []VSCodeThemeModel
}

type VSCodeThemeModel struct {
	Name    string                `json:"name"`
	UITheme VSCodeUITheme         `json:"-"`
	Type    ThemeType             `json:"type"`
	Colors  VSCodeColorsJSONModel `json:"colors"`
}

//nolint:tagliatelle
type VSCodeColorsJSONModel struct {
	ActivityBarActiveBorder                string `json:"activityBar.activeBorder"`
	ActivityBarBackground                  string `json:"activityBar.background"`
	ActivityBarBorder                      string `json:"activityBar.border"`
	ActivityBarForeground                  string `json:"activityBar.foreground"`
	ActivityBarInactiveForeground          string `json:"activityBar.inactiveForeground"`
	ActivityBarBadgeBackground             string `json:"activityBarBadge.background"`
	ActivityBarBadgeForeground             string `json:"activityBarBadge.foreground"`
	BadgeBackground                        string `json:"badge.background"`
	BadgeForeground                        string `json:"badge.foreground"`
	ButtonBackground                       string `json:"button.background"`
	ButtonBorder                           string `json:"button.border"`
	ButtonForeground                       string `json:"button.foreground"`
	ButtonHoverBackground                  string `json:"button.hoverBackground"`
	ButtonSecondaryBackground              string `json:"button.secondaryBackground"`
	ButtonSecondaryForeground              string `json:"button.secondaryForeground"`
	ButtonSecondaryHoverBackground         string `json:"button.secondaryHoverBackground"`
	ChatSlashCommandBackground             string `json:"chat.slashCommandBackground"`
	ChatSlashCommandForeground             string `json:"chat.slashCommandForeground"`
	ChatEditedFileForeground               string `json:"chat.editedFileForeground"`
	CheckboxBackground                     string `json:"checkbox.background"`
	CheckboxBorder                         string `json:"checkbox.border"`
	DebugToolBarBackground                 string `json:"debugToolBar.background"`
	DescriptionForeground                  string `json:"descriptionForeground"`
	DropdownBackground                     string `json:"dropdown.background"`
	DropdownBorder                         string `json:"dropdown.border"`
	DropdownForeground                     string `json:"dropdown.foreground"`
	DropdownListBackground                 string `json:"dropdown.listBackground"`
	EditorBackground                       string `json:"editor.background"`
	EditorFindMatchBackground              string `json:"editor.findMatchBackground"`
	EditorForeground                       string `json:"editor.foreground"`
	EditorGroupBorder                      string `json:"editorGroup.border"`
	EditorGroupHeaderTabsBackground        string `json:"editorGroupHeader.tabsBackground"`
	EditorGroupHeaderTabsBorder            string `json:"editorGroupHeader.tabsBorder"`
	EditorGutterAddedBackground            string `json:"editorGutter.addedBackground"`
	EditorGutterDeletedBackground          string `json:"editorGutter.deletedBackground"`
	EditorGutterModifiedBackground         string `json:"editorGutter.modifiedBackground"`
	EditorLineNumberActiveForeground       string `json:"editorLineNumber.activeForeground"`
	EditorLineNumberForeground             string `json:"editorLineNumber.foreground"`
	EditorOverviewRulerBorder              string `json:"editorOverviewRuler.border"`
	EditorWidgetBackground                 string `json:"editorWidget.background"`
	ErrorForeground                        string `json:"errorForeground"`
	FocusBorder                            string `json:"focusBorder"`
	Foreground                             string `json:"foreground"`
	IconForeground                         string `json:"icon.foreground"`
	InputBackground                        string `json:"input.background"`
	InputBorder                            string `json:"input.border"`
	InputForeground                        string `json:"input.foreground"`
	InputPlaceholderForeground             string `json:"input.placeholderForeground"`
	InputOptionActiveBackground            string `json:"inputOption.activeBackground"`
	InputOptionActiveBorder                string `json:"inputOption.activeBorder"`
	KeybindingLabelForeground              string `json:"keybindingLabel.foreground"`
	MenuBackground                         string `json:"menu.background"`
	MenuSelectionBackground                string `json:"menu.selectionBackground"`
	NotificationCenterHeaderBackground     string `json:"notificationCenterHeader.background"`
	NotificationCenterHeaderForeground     string `json:"notificationCenterHeader.foreground"`
	NotificationsBackground                string `json:"notifications.background"`
	NotificationsBorder                    string `json:"notifications.border"`
	NotificationsForeground                string `json:"notifications.foreground"`
	PanelBackground                        string `json:"panel.background"`
	PanelBorder                            string `json:"panel.border"`
	PanelInputBorder                       string `json:"panelInput.border"`
	PanelTitleActiveBorder                 string `json:"panelTitle.activeBorder"`
	PanelTitleActiveForeground             string `json:"panelTitle.activeForeground"`
	PanelTitleInactiveForeground           string `json:"panelTitle.inactiveForeground"`
	PeekViewEditorBackground               string `json:"peekViewEditor.background"`
	PeekViewEditorMatchHighlightBackground string `json:"peekViewEditor.matchHighlightBackground"`
	PeekViewResultBackground               string `json:"peekViewResult.background"`
	PeekViewResultMatchHighlightBackground string `json:"peekViewResult.matchHighlightBackground"`
	PickerGroupBorder                      string `json:"pickerGroup.border"`
	ProgressBarBackground                  string `json:"progressBar.background"`
	QuickInputBackground                   string `json:"quickInput.background"`
	QuickInputForeground                   string `json:"quickInput.foreground"`
	SettingsDropdownBackground             string `json:"settings.dropdownBackground"`
	SettingsDropdownBorder                 string `json:"settings.dropdownBorder"`
	SettingsHeaderForeground               string `json:"settings.headerForeground"`
	SettingsModifiedItemIndicator          string `json:"settings.modifiedItemIndicator"`
	SideBarBackground                      string `json:"sideBar.background"`
	SideBarBorder                          string `json:"sideBar.border"`
	SideBarForeground                      string `json:"sideBar.foreground"`
	SideBarSectionHeaderBackground         string `json:"sideBarSectionHeader.background"`
	SideBarSectionHeaderBorder             string `json:"sideBarSectionHeader.border"`
	SideBarSectionHeaderForeground         string `json:"sideBarSectionHeader.foreground"`
	SideBarTitleForeground                 string `json:"sideBarTitle.foreground"`
	StatusBarBackground                    string `json:"statusBar.background"`
	StatusBarBorder                        string `json:"statusBar.border"`
	StatusBarDebuggingBackground           string `json:"statusBar.debuggingBackground"`
	StatusBarDebuggingForeground           string `json:"statusBar.debuggingForeground"`
	StatusBarFocusBorder                   string `json:"statusBar.focusBorder"`
	StatusBarForeground                    string `json:"statusBar.foreground"`
	StatusBarNoFolderBackground            string `json:"statusBar.noFolderBackground"`
	StatusBarItemFocusBorder               string `json:"statusBarItem.focusBorder"`
	StatusBarItemProminentBackground       string `json:"statusBarItem.prominentBackground"`
	StatusBarItemRemoteBackground          string `json:"statusBarItem.remoteBackground"`
	StatusBarItemRemoteForeground          string `json:"statusBarItem.remoteForeground"`
	TabActiveBackground                    string `json:"tab.activeBackground"`
	TabActiveBorder                        string `json:"tab.activeBorder"`
	TabActiveBorderTop                     string `json:"tab.activeBorderTop"`
	TabActiveForeground                    string `json:"tab.activeForeground"`
	TabSelectedBorderTop                   string `json:"tab.selectedBorderTop"`
	TabBorder                              string `json:"tab.border"`
	TabHoverBackground                     string `json:"tab.hoverBackground"`
	TabInactiveBackground                  string `json:"tab.inactiveBackground"`
	TabInactiveForeground                  string `json:"tab.inactiveForeground"`
	TabUnfocusedActiveBorder               string `json:"tab.unfocusedActiveBorder"`
	TabUnfocusedActiveBorderTop            string `json:"tab.unfocusedActiveBorderTop"`
	TabUnfocusedHoverBackground            string `json:"tab.unfocusedHoverBackground"`
	TerminalForeground                     string `json:"terminal.foreground"`
	TerminalTabActiveBorder                string `json:"terminal.tab.activeBorder"`
	TextBlockQuoteBackground               string `json:"textBlockQuote.background"`
	TextBlockQuoteBorder                   string `json:"textBlockQuote.border"`
	TextCodeBlockBackground                string `json:"textCodeBlock.background"`
	TextLinkActiveForeground               string `json:"textLink.activeForeground"`
	TextLinkForeground                     string `json:"textLink.foreground"`
	TextPreformatForeground                string `json:"textPreformat.foreground"`
	TextPreformatBackground                string `json:"textPreformat.background"`
	TextSeparatorForeground                string `json:"textSeparator.foreground"`
	TitleBarActiveBackground               string `json:"titleBar.activeBackground"`
	TitleBarActiveForeground               string `json:"titleBar.activeForeground"`
	TitleBarBorder                         string `json:"titleBar.border"`
	TitleBarInactiveBackground             string `json:"titleBar.inactiveBackground"`
	TitleBarInactiveForeground             string `json:"titleBar.inactiveForeground"`
	WelcomePageTileBackground              string `json:"welcomePage.tileBackground"`
	WelcomePageProgressForeground          string `json:"welcomePage.progress.foreground"`
	WidgetBorder                           string `json:"widget.border"`
	EditorInactiveSelectionBackground      string `json:"editor.inactiveSelectionBackground"`
	EditorIndentGuideBackground1           string `json:"editorIndentGuide.background1"`
	EditorIndentGuideActiveBackground1     string `json:"editorIndentGuide.activeBackground1"`
	EditorSelectionHighlightBackground     string `json:"editor.selectionHighlightBackground"`
	ListDropBackground                     string `json:"list.dropBackground"`
	MenuForeground                         string `json:"menu.foreground"`
	MenuSeparatorBackground                string `json:"menu.separatorBackground"`
	MenuBorder                             string `json:"menu.border"`
	TabSelectedBackground                  string `json:"tab.selectedBackground"`
	TabSelectedForeground                  string `json:"tab.selectedForeground"`
	TabLastPinnedBorder                    string `json:"tab.lastPinnedBorder"`
	ListActiveSelectionIconForeground      string `json:"list.activeSelectionIconForeground"`
	TerminalInactiveSelectionBackground    string `json:"terminal.inactiveSelectionBackground"`
	ActionBarToggledBackground             string `json:"actionBar.toggledBackground"`

	TerminalAnsiBlack         string `json:"terminal.ansiBlack"`
	TerminalAnsiBrightBlack   string `json:"terminal.ansiBrightBlack"`
	TerminalAnsiBlue          string `json:"terminal.ansiBlue"`
	TerminalAnsiCyan          string `json:"terminal.ansiCyan"`
	TerminalAnsiGreen         string `json:"terminal.ansiGreen"`
	TerminalAnsiBrightBlue    string `json:"terminal.ansiBrightBlue"`
	TerminalAnsiBrightCyan    string `json:"terminal.ansiBrightCyan"`
	TerminalAnsiBrightGreen   string `json:"terminal.ansiBrightGreen"`
	TerminalAnsiWhite         string `json:"terminal.ansiWhite"`
	TerminalAnsiBrightWhite   string `json:"terminal.ansiBrightWhite"`
	TerminalAnsiMagenta       string `json:"terminal.ansiMagenta"`
	TerminalAnsiYellow        string `json:"terminal.ansiYellow"`
	TerminalAnsiRed           string `json:"terminal.ansiRed"`
	TerminalAnsiBrightYellow  string `json:"terminal.ansiBrightYellow"`
	TerminalAnsiBrightMagenta string `json:"terminal.ansiBrightMagenta"`
	TerminalAnsiBrightRed     string `json:"terminal.ansiBrightRed"`
}

type VSCodeColorsModel struct {
	Foreground          string
	EditorBackground    string
	Background          string
	BackgroundAccent    string
	WarningForeground   string
	HiForeground        string
	Border              string
	ForegroundAccent    string
	SecondaryForeground string
	SecondaryBackground string
	HiBackgroundAccent  string
	MatchBackground     string

	TerminalAnsi TerminalAnsiModel
}

//nolint:funlen
func VSCodeColorsModelToJSONModel(model *VSCodeColorsModel) *VSCodeColorsJSONModel {
	return &VSCodeColorsJSONModel{
		ActivityBarActiveBorder:        model.BackgroundAccent,
		ActivityBarBadgeBackground:     model.BackgroundAccent,
		ButtonBackground:               model.BackgroundAccent,
		EditorGutterModifiedBackground: model.BackgroundAccent,
		FocusBorder:                    model.BackgroundAccent,
		PanelTitleActiveBorder:         model.BackgroundAccent,
		ProgressBarBackground:          model.BackgroundAccent,
		StatusBarDebuggingBackground:   model.BackgroundAccent,
		StatusBarFocusBorder:           model.BackgroundAccent,
		StatusBarItemFocusBorder:       model.BackgroundAccent,
		StatusBarItemRemoteBackground:  model.BackgroundAccent,
		TabActiveBorderTop:             model.BackgroundAccent,
		TerminalTabActiveBorder:        model.BackgroundAccent,
		WelcomePageProgressForeground:  model.BackgroundAccent,
		MenuSelectionBackground:        model.BackgroundAccent,

		ActivityBarBackground:           model.Background,
		DebugToolBarBackground:          model.Background,
		EditorGroupHeaderTabsBackground: model.Background,
		PanelBackground:                 model.Background,
		SideBarBackground:               model.Background,
		SideBarSectionHeaderBackground:  model.Background,
		StatusBarBackground:             model.Background,
		TabInactiveBackground:           model.Background,
		TitleBarActiveBackground:        model.Background,

		DropdownListBackground:             model.EditorBackground,
		EditorBackground:                   model.EditorBackground,
		MenuBackground:                     model.EditorBackground,
		NotificationCenterHeaderBackground: model.EditorBackground,
		NotificationsBackground:            model.EditorBackground,
		PeekViewEditorBackground:           model.EditorBackground,
		PeekViewResultBackground:           model.EditorBackground,
		StatusBarNoFolderBackground:        model.EditorBackground,
		TabActiveBackground:                model.EditorBackground,
		TabActiveBorder:                    model.EditorBackground,
		TabHoverBackground:                 model.EditorBackground,
		TabUnfocusedActiveBorder:           model.EditorBackground,
		TabUnfocusedHoverBackground:        model.EditorBackground,
		TitleBarInactiveBackground:         model.EditorBackground,

		ButtonSecondaryForeground:          model.Foreground,
		DropdownForeground:                 model.Foreground,
		EditorForeground:                   model.Foreground,
		EditorLineNumberActiveForeground:   model.Foreground,
		Foreground:                         model.Foreground,
		IconForeground:                     model.Foreground,
		InputForeground:                    model.Foreground,
		KeybindingLabelForeground:          model.Foreground,
		NotificationCenterHeaderForeground: model.Foreground,
		NotificationsForeground:            model.Foreground,
		PanelTitleActiveForeground:         model.Foreground,
		QuickInputForeground:               model.Foreground,
		SideBarForeground:                  model.Foreground,
		SideBarSectionHeaderForeground:     model.Foreground,
		SideBarTitleForeground:             model.Foreground,
		StatusBarForeground:                model.Foreground,
		TerminalForeground:                 model.Foreground,
		TitleBarActiveForeground:           model.Foreground,
		MenuForeground:                     model.Foreground,

		EditorGutterDeletedBackground: model.WarningForeground,
		ErrorForeground:               model.WarningForeground,

		BadgeForeground:                   model.HiForeground,
		ListActiveSelectionIconForeground: model.HiForeground,
		ActivityBarBadgeForeground:        model.HiForeground,
		ButtonBorder:                      model.HiForeground + "12",
		ButtonForeground:                  model.HiForeground,
		EditorGroupBorder:                 model.HiForeground + "17",
		SettingsHeaderForeground:          model.HiForeground,
		StatusBarDebuggingForeground:      model.HiForeground,
		StatusBarItemRemoteForeground:     model.HiForeground,
		TabActiveForeground:               model.HiForeground,
		TabSelectedForeground:             model.HiForeground + "a0",

		ActivityBarBorder:           model.Border,
		EditorGroupHeaderTabsBorder: model.Border,
		NotificationsBorder:         model.Border,
		PanelBorder:                 model.Border,
		PanelInputBorder:            model.Border,
		SideBarBorder:               model.Border,
		SideBarSectionHeaderBorder:  model.Border,
		StatusBarBorder:             model.Border,
		TabBorder:                   model.Border,
		TabUnfocusedActiveBorderTop: model.Border,
		TextBlockQuoteBackground:    model.Border,
		TextCodeBlockBackground:     model.Border,
		TitleBarBorder:              model.Border,
		WelcomePageTileBackground:   model.Border,

		ButtonSecondaryBackground:  model.Border,
		CheckboxBackground:         model.Border,
		DropdownBackground:         model.Border,
		InputBackground:            model.Border,
		SettingsDropdownBackground: model.Border,
		WidgetBorder:               model.Border,

		ButtonSecondaryHoverBackground: model.Border,
		CheckboxBorder:                 model.Border,
		DropdownBorder:                 model.Border,
		InputBorder:                    model.Border,
		PickerGroupBorder:              model.Border,
		SettingsDropdownBorder:         model.Border,
		TextPreformatBackground:        model.Border,

		EditorOverviewRulerBorder: model.Border,

		EditorWidgetBackground: model.Background,

		TextSeparatorForeground: model.Border,

		QuickInputBackground:  model.Background,
		TabSelectedBackground: model.Background,

		ChatSlashCommandBackground: model.Background,

		ListDropBackground: model.Background,

		ActionBarToggledBackground: model.Background,

		EditorInactiveSelectionBackground:   model.Background,
		TerminalInactiveSelectionBackground: model.Background,

		EditorIndentGuideBackground1: model.Background,

		MenuSeparatorBackground: model.SecondaryBackground,
		MenuBorder:              model.SecondaryBackground,

		BadgeBackground:      model.SecondaryBackground,
		TextBlockQuoteBorder: model.SecondaryBackground,

		EditorLineNumberForeground:       model.SecondaryForeground,
		StatusBarItemProminentBackground: model.SecondaryForeground + "66",

		EditorIndentGuideActiveBackground1: model.SecondaryBackground,

		InputPlaceholderForeground: model.SecondaryForeground,

		DescriptionForeground:        model.SecondaryForeground,
		PanelTitleInactiveForeground: model.SecondaryForeground,
		TabInactiveForeground:        model.SecondaryForeground,
		TitleBarInactiveForeground:   model.SecondaryForeground,

		ActivityBarInactiveForeground: model.SecondaryForeground,

		TabLastPinnedBorder: model.Foreground + "33",

		TextPreformatForeground: model.Foreground,

		ActivityBarForeground: model.Foreground,

		ButtonHoverBackground: model.BackgroundAccent,

		InputOptionActiveBorder: model.BackgroundAccent,

		InputOptionActiveBackground: model.BackgroundAccent + "82",

		EditorGutterAddedBackground: model.BackgroundAccent,

		ChatSlashCommandForeground: model.ForegroundAccent,

		TextLinkActiveForeground: model.ForegroundAccent,
		TextLinkForeground:       model.ForegroundAccent,

		TabSelectedBorderTop: model.ForegroundAccent,

		EditorSelectionHighlightBackground: model.HiBackgroundAccent,

		EditorFindMatchBackground: model.MatchBackground,

		PeekViewEditorMatchHighlightBackground: model.MatchBackground + "aa",
		PeekViewResultMatchHighlightBackground: model.MatchBackground + "aa",
		SettingsModifiedItemIndicator:          model.MatchBackground + "aa",

		ChatEditedFileForeground: model.ForegroundAccent,

		TerminalAnsiBlack:         model.TerminalAnsi.TerminalAnsiBlack,
		TerminalAnsiBrightBlack:   model.TerminalAnsi.TerminalAnsiBrightBlack,
		TerminalAnsiBlue:          model.TerminalAnsi.TerminalAnsiBlue,
		TerminalAnsiCyan:          model.TerminalAnsi.TerminalAnsiCyan,
		TerminalAnsiGreen:         model.TerminalAnsi.TerminalAnsiGreen,
		TerminalAnsiBrightBlue:    model.TerminalAnsi.TerminalAnsiBrightBlue,
		TerminalAnsiBrightCyan:    model.TerminalAnsi.TerminalAnsiBrightCyan,
		TerminalAnsiBrightGreen:   model.TerminalAnsi.TerminalAnsiBrightGreen,
		TerminalAnsiWhite:         model.TerminalAnsi.TerminalAnsiWhite,
		TerminalAnsiBrightWhite:   model.TerminalAnsi.TerminalAnsiBrightWhite,
		TerminalAnsiMagenta:       model.TerminalAnsi.TerminalAnsiMagenta,
		TerminalAnsiYellow:        model.TerminalAnsi.TerminalAnsiYellow,
		TerminalAnsiRed:           model.TerminalAnsi.TerminalAnsiRed,
		TerminalAnsiBrightYellow:  model.TerminalAnsi.TerminalAnsiBrightYellow,
		TerminalAnsiBrightMagenta: model.TerminalAnsi.TerminalAnsiBrightMagenta,
		TerminalAnsiBrightRed:     model.TerminalAnsi.TerminalAnsiBrightRed,
	}
}
