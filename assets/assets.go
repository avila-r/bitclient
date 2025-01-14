package assets

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors used for styling
	Normal      = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}         // Default color
	LightOrange = lipgloss.AdaptiveColor{Light: "#F2A900", Dark: "#F2A900"} // Light orange color
	DarkOrange  = lipgloss.AdaptiveColor{Light: "#F7931A", Dark: "#F7931A"} // Dark orange color
	Black       = lipgloss.AdaptiveColor{Light: "#000", Dark: "#000"}       // Black color
	Gray        = lipgloss.AdaptiveColor{Light: "#A9A9A9", Dark: "#A9A9A9"} // Gray color
	White       = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#FFF"}       // White color
	Red         = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"} // Red color

	// Emoji symbols used in output
	EmojiSparkles = "\U00002728" // ✨ Sparkles
	EmojiError    = "\U0000274C" // ❌ Error
	EmojiTick     = "\U00002714" // ✔ Tick
	EmojiThumbsUp = "\U0001F44D" // 👍 Thumbs Up
	EmojiConfused = "\U0001F615" // 😕 Confused

	// Logo style with color and border properties
	Logo = lipgloss.NewStyle().
		Foreground(LightOrange).             // Light orange text color
		PaddingLeft(1).                      // Padding on the left
		Bold(true).                          // Bold text style
		BorderStyle(lipgloss.ThickBorder()). // Thick border around the logo
		BorderLeft(true).                    // Border on the left side of the logo
		BorderForeground(DarkOrange).        // Dark orange border color
		Render(`
				\  / _  _
				\/ (_](_)
					._|
				`)

	// Text style with specific padding, border, and color settings
	Text = lipgloss.NewStyle().
		PaddingLeft(1).                      // Padding on the left
		BorderStyle(lipgloss.ThickBorder()). // Thick border around the text
		BorderLeft(true).                    // Border on the left side of the text
		BorderForeground(DarkOrange).        // Dark orange border color
		Foreground(LightOrange)              // Light orange text color

	// Custom theme for the "huh" package (CLI input handling)
	FormTheme = func() *huh.Theme {
		t := huh.ThemeBase() // Base theme from huh package

		// Customizations for focused input fields
		t.Focused.Base = t.Focused.Base.BorderForeground(DarkOrange)                                                                         // Border color for focused elements
		t.Focused.Title = t.Focused.Title.Foreground(LightOrange).Bold(true)                                                                 // Title style with light orange color and bold text
		t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(LightOrange)                                                                    // Note title with light orange color
		t.Focused.Directory = t.Focused.Directory.Foreground(Gray)                                                                           // Directory text color set to gray
		t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})                             // Description color for dark mode
		t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(LightOrange)                                                          // Color for selection
		t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(LightOrange)                                                            // Color for next indicator
		t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(LightOrange)                                                            // Color for previous indicator
		t.Focused.Option = t.Focused.Option.Foreground(Normal)                                                                               // Option color set to normal
		t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(LightOrange)                                                // Multi select option color
		t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(LightOrange)                                                          // Selected option color
		t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ") // Checkmark for selected options
		t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")          // Bullet for unselected options
		t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(Normal)                                                           // Unselected option color
		t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(White).Background(LightOrange)                                          // Focused button with white text and light orange background
		t.Focused.Next = t.Focused.FocusedButton                                                                                             // Use the focused button as the next button style
		t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(Normal).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})   // Blurred button style with normal foreground and light gray background

		// Customizations for text input fields
		t.Focused.TextInput.Text = t.Focused.TextInput.Text.Foreground(Black)                                                           // Text input text color set to black
		t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(DarkOrange)                                                  // Cursor color set to dark orange
		t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"}) // Placeholder text color
		t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(LightOrange)                                                 // Prompt color set to light orange

		// Blurred state inherits focused state
		t.Blurred = t.Focused
		t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder()) // Hide border in blurred state
		t.Blurred.NextIndicator = lipgloss.NewStyle()                        // Reset next indicator style
		t.Blurred.PrevIndicator = lipgloss.NewStyle()                        // Reset previous indicator style

		return t // Return the customized theme
	}()
)
