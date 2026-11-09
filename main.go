package main

import "github.com/gdamore/tcell/v2"

func main() {

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	screen.Clear()
	blueStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)

	text := [][]rune{{}}
	row:= 0
	col:= 0
	for {
		screen.Clear()
		for y, line := range text {
			for x, char := range line {
				screen.PutStrStyled(x, y, string(char), blueStyle)
			}
		}
		//screen.PutStrStyled(col, row, string(text[row][col]), blueStyle)
		screen.ShowCursor(col, row)
		screen.Show()
		ev := screen.PollEvent()
		switch event := ev.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape {
				return
			} else if event.Key() == tcell.KeyRune {
				insertChar(&text, row, col, event.Rune())
				col++	
			} else if event.Key() == tcell.KeyEnter {
				// Split the current line at the cursor position
				currentLine := text[row]
				newLine := currentLine[col:]
				text[row] = currentLine[:col]
				text = append(text[:row+1], append([][]rune{newLine}, text[row+1:]...)...)
				row++
				col = 0
			} else if event.Key() == tcell.KeyRight {
				col++
			} else if event.Key() == tcell.KeyLeft {
				if col > 0 {
				col--
				}
			} else if event.Key() == tcell.KeyUp {
				if row > 0 {
					row--
				}
			} else if event.Key() == tcell.KeyDown {
				if row < len(text)-1 {
					row++
				}
			} else if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
				if len(text) > 0 {
					text[row] = text[row][:len(text[row])-1] 
				}
			}
		case *tcell.EventResize:
			screen.Sync()

		}

	}
}
func insertChar(text *[][]rune, x, y int, char rune) {
    line := (*text)[x]

    // Make room for the new character
    line = append(line, 0)

    // Shift characters right
    copy(line[y+1:], line[y:])

    // Insert character at cursor position
    line[y] = char

    // Save updated line
    (*text)[x] = line
}
