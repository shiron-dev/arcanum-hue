package converter

type VSCodeThemeModel struct {
	Name   string                `json:"name"`
	Type   ThemeType             `json:"type"`
	Colors VSCodeColorsJSONModel `json:"colors"`
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
}

type VSCodeColorsModel struct {
	Foreground              string // #D7D7D7
	SecondaryForeground     string // #9D9D9D
	Background              string // #181818
	SecondaryBackground     string // #313131
	BackgroundHighlight     string // #0078D4
	PeekHighlightBackground string // #BB800966
	BorderForeground        string // #2B2B2B
	BorderBackground        string // #3C3C3C
	EditorForeground        string // #CCCCCC
	EditorBackground        string // #1F1F1F
	DeleteForeground        string // #F85149
	Border2                 string // #616161
	TextLinkForeground      string // #4DA3FF
	// Uncategorized
	ActivityBarInactiveForeground    string // #868686
	BadgeForeground                  string // #F8F8F8
	ButtonHoverBackground            string // #026EC1
	ChatSlashCommandForeground       string // #40A6FF
	ChatEditedFileForeground         string // #E2C08D
	EditorFindMatchBackground        string // #9E6A03
	EditorGutterAddedBackground      string // #2EA043
	EditorLineNumberForeground       string // #6E7681
	EditorOverviewRulerBorder        string // #010409
	EditorWidgetBackground           string // #202020
	InputPlaceholderForeground       string // #989898
	InputOptionActiveBackground      string // #2489DB82
	InputOptionActiveBorder          string // #2488DB
	QuickInputBackground             string // #222222
	StatusBarItemProminentBackground string // #6E768166
	TabSelectedBorderTop             string // #6caddf
	TextPreformatForeground          string // #D0D0D0
	TextSeparatorForeground          string // #21262D
}

//nolint:funlen
func ColorsModelToJSONModel(model ColorsModel) ColorsJSONModel {
	foregroundBorder := model.Foreground + "12"

	return &VSCodeColorsJSONModel{
		ButtonForeground:                       model.Foreground,
		SettingsHeaderForeground:               model.Foreground,
		StatusBarDebuggingForeground:           model.Foreground,
		StatusBarItemRemoteForeground:          model.Foreground,
		TabActiveForeground:                    model.Foreground,
		ActivityBarBadgeForeground:             model.Foreground,
		ActivityBarForeground:                  model.Foreground,
		ButtonBorder:                           foregroundBorder,
		EditorGroupBorder:                      foregroundBorder,
		ActivityBarBackground:                  model.Background,
		DebugToolBarBackground:                 model.Background,
		EditorGroupHeaderTabsBackground:        model.Background,
		PanelBackground:                        model.Background,
		SideBarBackground:                      model.Background,
		SideBarSectionHeaderBackground:         model.Background,
		StatusBarBackground:                    model.Background,
		TabInactiveBackground:                  model.Background,
		TitleBarActiveBackground:               model.Background,
		DropdownListBackground:                 model.EditorBackground,
		EditorBackground:                       model.EditorBackground,
		MenuBackground:                         model.EditorBackground,
		NotificationCenterHeaderBackground:     model.EditorBackground,
		NotificationsBackground:                model.EditorBackground,
		PeekViewEditorBackground:               model.EditorBackground,
		PeekViewResultBackground:               model.EditorBackground,
		StatusBarNoFolderBackground:            model.EditorBackground,
		TabActiveBackground:                    model.EditorBackground,
		TabActiveBorder:                        model.EditorBackground,
		TabHoverBackground:                     model.EditorBackground,
		TabUnfocusedActiveBorder:               model.EditorBackground,
		TabUnfocusedHoverBackground:            model.EditorBackground,
		TitleBarInactiveBackground:             model.EditorBackground,
		ButtonSecondaryForeground:              model.EditorForeground,
		DropdownForeground:                     model.EditorForeground,
		EditorForeground:                       model.EditorForeground,
		EditorLineNumberActiveForeground:       model.EditorForeground,
		Foreground:                             model.EditorForeground,
		IconForeground:                         model.EditorForeground,
		InputForeground:                        model.EditorForeground,
		KeybindingLabelForeground:              model.EditorForeground,
		NotificationCenterHeaderForeground:     model.EditorForeground,
		NotificationsForeground:                model.EditorForeground,
		PanelTitleActiveForeground:             model.EditorForeground,
		QuickInputForeground:                   model.EditorForeground,
		SideBarForeground:                      model.EditorForeground,
		SideBarSectionHeaderForeground:         model.EditorForeground,
		StatusBarForeground:                    model.EditorForeground,
		TerminalForeground:                     model.EditorForeground,
		TitleBarActiveForeground:               model.EditorForeground,
		SideBarTitleForeground:                 model.EditorForeground,
		ButtonSecondaryBackground:              model.SecondaryBackground,
		CheckboxBackground:                     model.SecondaryBackground,
		DropdownBackground:                     model.SecondaryBackground,
		InputBackground:                        model.SecondaryBackground,
		SettingsDropdownBackground:             model.SecondaryBackground,
		WidgetBorder:                           model.SecondaryBackground,
		ActivityBarBorder:                      model.BorderForeground,
		EditorGroupHeaderTabsBorder:            model.BorderForeground,
		NotificationsBorder:                    model.BorderForeground,
		PanelBorder:                            model.BorderForeground,
		PanelInputBorder:                       model.BorderForeground,
		SideBarBorder:                          model.BorderForeground,
		SideBarSectionHeaderBorder:             model.BorderForeground,
		StatusBarBorder:                        model.BorderForeground,
		TabBorder:                              model.BorderForeground,
		TabUnfocusedActiveBorderTop:            model.BorderForeground,
		TextBlockQuoteBackground:               model.BorderForeground,
		TextCodeBlockBackground:                model.BorderForeground,
		TitleBarBorder:                         model.BorderForeground,
		WelcomePageTileBackground:              model.BorderForeground,
		ActivityBarActiveBorder:                model.BackgroundHighlight,
		ActivityBarBadgeBackground:             model.BackgroundHighlight,
		ButtonBackground:                       model.BackgroundHighlight,
		EditorGutterModifiedBackground:         model.BackgroundHighlight,
		FocusBorder:                            model.BackgroundHighlight,
		MenuSelectionBackground:                model.BackgroundHighlight,
		PanelTitleActiveBorder:                 model.BackgroundHighlight,
		ProgressBarBackground:                  model.BackgroundHighlight,
		StatusBarDebuggingBackground:           model.BackgroundHighlight,
		StatusBarFocusBorder:                   model.BackgroundHighlight,
		StatusBarItemFocusBorder:               model.BackgroundHighlight,
		StatusBarItemRemoteBackground:          model.BackgroundHighlight,
		TabActiveBorderTop:                     model.BackgroundHighlight,
		TerminalTabActiveBorder:                model.BackgroundHighlight,
		WelcomePageProgressForeground:          model.BackgroundHighlight,
		ButtonSecondaryHoverBackground:         model.Background,
		CheckboxBorder:                         model.BorderBackground,
		DropdownBorder:                         model.BorderBackground,
		InputBorder:                            model.BorderBackground,
		PickerGroupBorder:                      model.BorderBackground,
		SettingsDropdownBorder:                 model.BorderBackground,
		TextPreformatBackground:                model.BorderBackground,
		DescriptionForeground:                  model.SecondaryForeground,
		PanelTitleInactiveForeground:           model.SecondaryForeground,
		TabInactiveForeground:                  model.SecondaryForeground,
		TitleBarInactiveForeground:             model.SecondaryForeground,
		EditorGutterDeletedBackground:          model.DeleteForeground,
		ErrorForeground:                        model.DeleteForeground,
		PeekViewEditorMatchHighlightBackground: model.PeekHighlightBackground,
		PeekViewResultMatchHighlightBackground: model.PeekHighlightBackground,
		SettingsModifiedItemIndicator:          model.PeekHighlightBackground,
		BadgeBackground:                        model.Border2,
		TextBlockQuoteBorder:                   model.Border2,
		TextLinkActiveForeground:               model.TextLinkForeground,
		TextLinkForeground:                     model.TextLinkForeground,
		ActivityBarInactiveForeground:          model.ActivityBarInactiveForeground,
		BadgeForeground:                        model.BadgeForeground,
		ButtonHoverBackground:                  model.ButtonHoverBackground,
		ChatSlashCommandForeground:             model.ChatSlashCommandForeground,
		ChatEditedFileForeground:               model.ChatEditedFileForeground,
		EditorFindMatchBackground:              model.EditorFindMatchBackground,
		EditorGutterAddedBackground:            model.EditorGutterAddedBackground,
		EditorLineNumberForeground:             model.EditorLineNumberForeground,
		EditorOverviewRulerBorder:              model.EditorOverviewRulerBorder,
		EditorWidgetBackground:                 model.EditorWidgetBackground,
		InputPlaceholderForeground:             model.InputPlaceholderForeground,
		InputOptionActiveBackground:            model.InputOptionActiveBackground,
		InputOptionActiveBorder:                model.InputOptionActiveBorder,
		QuickInputBackground:                   model.QuickInputBackground,
		StatusBarItemProminentBackground:       model.StatusBarItemProminentBackground,
		TabSelectedBorderTop:                   model.TabSelectedBorderTop,
		TextPreformatForeground:                model.TextPreformatForeground,
		TextSeparatorForeground:                model.TextSeparatorForeground,
	}
}
