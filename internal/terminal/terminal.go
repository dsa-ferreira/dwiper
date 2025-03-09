package terminal

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
)

func InteractiveSelection(options []string) []string {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}
	defer screen.Fini()

	selected := make(map[int]bool) // Stores selected indices
	currentIndex := 0

	// Function to render the menu
	drawMenu := func() {
		screen.Clear()
		styleNormal := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		styleSelected := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)

		for i, option := range options {
			prefix := "  "
			style := styleNormal

			if i == currentIndex {
				style = styleSelected
				prefix = "> "
			}

			if selected[i] {
				option = "[X] " + option
			} else {
				option = "[ ] " + option
			}

			// Print the option
			for j, char := range prefix + option {
				screen.SetContent(j, i+1, char, nil, style)
			}
		}

		screen.Show()
	}

	// Draw initial menu
	drawMenu()

	// Event loop for user input
	for {
		event := screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				if currentIndex > 0 {
					currentIndex--
				}
			case tcell.KeyDown:
				if currentIndex < len(options)-1 {
					currentIndex++
				}
			case tcell.KeyEnter:
				selected[currentIndex] = !selected[currentIndex] // Toggle selection
			case tcell.KeyEscape:
				// Collect selected items and return
				var result []string
				for i, isSelected := range selected {
					if isSelected {
						result = append(result, options[i])
					}
				}
				return result
			}

			// Redraw the menu after key press
			drawMenu()
		}
	}

}
