package assets

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	Normal      = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	LightOrange = lipgloss.AdaptiveColor{Light: "#F2A900", Dark: "#F2A900"}
	DarkOrange  = lipgloss.AdaptiveColor{Light: "#F7931A", Dark: "#F7931A"}
	Black       = lipgloss.AdaptiveColor{Light: "#000", Dark: "#000"}

	Gray  = lipgloss.AdaptiveColor{Light: "#A9A9A9", Dark: "#A9A9A9"}
	White = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#FFF"}
	Red   = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}

	EmojiSparkles = "\U00002728" // ✨
	EmojiError    = "\U0000274C" // ❌
	EmojiTick     = "\U00002714" // ✔
	EmojiThumbsUp = "\U0001F44D" // 👍
	EmojiConfused = "\U0001F615" // 😕

	Logo = lipgloss.NewStyle().
		Foreground(LightOrange).
		PaddingLeft(1).
		Bold(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(DarkOrange).
		Render(`
				\  / _  _
				\/ (_](_)
					._|
				`)

	Text = lipgloss.NewStyle().
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(DarkOrange).
		Foreground(LightOrange)

	FormTheme = func() *huh.Theme {
		t := huh.ThemeBase()

		t.Focused.Base = t.Focused.Base.BorderForeground(DarkOrange)
		t.Focused.Title = t.Focused.Title.Foreground(LightOrange).Bold(true)
		t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(LightOrange)
		t.Focused.Directory = t.Focused.Directory.Foreground(Gray)
		t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
		t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(LightOrange)
		t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(LightOrange)
		t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(LightOrange)
		t.Focused.Option = t.Focused.Option.Foreground(Normal)
		t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(LightOrange)
		t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(LightOrange)
		t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
		t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
		t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(Normal)
		t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(White).Background(LightOrange)
		t.Focused.Next = t.Focused.FocusedButton
		t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(Normal).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

		t.Focused.TextInput.Text = t.Focused.TextInput.Text.Foreground(Black)
		t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(DarkOrange)
		t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
		t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(LightOrange)

		t.Blurred = t.Focused
		t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
		t.Blurred.NextIndicator = lipgloss.NewStyle()
		t.Blurred.PrevIndicator = lipgloss.NewStyle()

		return t
	}()
)
