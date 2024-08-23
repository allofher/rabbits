package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/allofher/rabbits/internals"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	p := tea.NewProgram(model.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("R.I.P. : %v", err)
		os.Exit(1)
	}
}
