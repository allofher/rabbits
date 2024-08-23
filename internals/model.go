package model

import (
	"time"
	"strings"
	"math/rand"
	tea "github.com/charmbracelet/bubbletea"
)



const (
	LEFT = 1 + iota
	RIGHT
	UP
	DOWN
)

const (
	HELP     = "\n\nuse arrow keys, frst or hjkl to move the snake.\n"
	GAMEOVER = "\n\nGAME OVER\n"
	QUIT     = "press 'q' or 'ctrl + c' to quit.\n"
)

const INTERVAL = 100

type TickMsg time.Time

type Model struct {
	allRabbits []rabbit
	width int
	height int
	arena [][]string
	food coord
	score int
	impassableSymbol string
	foodSymbol string
	rabbitSymbol string
	verticalEdgeSymbol string
	horizontalEdgeSymbol string
	emptySymbol string
}

type select_box struct {
	startPos coord
	endPos   coord
}

var selectionBox = select_box{
	startPos: coord{x: 0, y: 0},
	endPos: coord{x: 0, y: 0},
}

var selectedRabbits = []rabbit{}

func InitialModel() Model {
	return Model{
		allRabbits: []rabbit{{position: coord{x: 2, y: 2}, direction: DOWN}},
		width: 200,
		height: 35,
		arena: [][]string{},
		food: coord{x: 10, y: 10},
		score: 0,
		impassableSymbol: "#",
		foodSymbol: "*",
		rabbitSymbol: "@",
		verticalEdgeSymbol: ".",
		horizontalEdgeSymbol: ".",
		emptySymbol: " ",
	}
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Duration(INTERVAL)*time.Millisecond, func(t time.Time) tea.Msg{
		return TickMsg(t)
	})
}

func (m Model) changeDirection(direction int) (tea.Model, tea.Cmd) {
	if len(selectedRabbits) < 1 {
		return m, nil
	}

	for i, _ := range selectedRabbits {
		m.allRabbits[i].direction = direction
	}
	
	return m, nil
}

func (m Model) selectRabbits() []rabbit{
	
	var transformedBox = select_box{
		startPos: coord{x: min(selectionBox.startPos.x, selectionBox.endPos.x), y: min(selectionBox.startPos.y, selectionBox.endPos.y)},
		endPos: coord{x: max(selectionBox.startPos.x, selectionBox.endPos.x), y: max(selectionBox.startPos.y, selectionBox.endPos.y)},
	}

	boxedRabbits := []rabbit{}
	for i, val := range m.allRabbits {
		if transformedBox.startPos.x <= val.position.x &&
			val.position.x <= transformedBox.endPos.x &&
			transformedBox.startPos.y <= val.position.y &&
			val.position.y <= transformedBox.endPos.y {
			boxedRabbits = append(boxedRabbits, m.allRabbits[i])
		}
	}

	return boxedRabbits
}

func (m Model) Init() tea.Cmd {
	var x, y int

	x = rand.Intn(m.height - 1)
	y = rand.Intn(m.width - 1)

	m.food = coord{x: x, y: y}
	selectedRabbits = m.allRabbits
	
	return m.tick()
}

func (m Model) moveRabbits() (tea.Model, tea.Cmd) {
	for i, val := range selectedRabbits {
		switch val.direction {
		case UP:
			m.allRabbits[i].position.x--
		case DOWN:
			m.allRabbits[i].position.x++
		case LEFT:
			m.allRabbits[i].position.y--
		case RIGHT:
			m.allRabbits[i].position.y++
		}

		if val.position.x == m.food.x && val.position.y == m.food.y {
			m.food.x = rand.Intn(m.height-2) + 1
			m.food.y = rand.Intn(m.width-2) + 1
		}
	}
	
	return m, m.tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.MouseMsg:
		if msg.Action == 0 && msg.Button == 1 {
			selectionBox.startPos = coord{x: msg.X, y: msg.Y}
		} else if msg.Action == 1 && msg.Button == 1 {
			selectionBox.endPos = coord{x: msg.X, y: msg.Y}
			selectedRabbits = m.selectRabbits()
		}
		
	
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "f", "k":
			return m.changeDirection(UP)
		case "down", "s", "j":
			return m.changeDirection(DOWN)
		case "left", "r", "h":
			return m.changeDirection(LEFT)
		case "right", "t", "l":
			return m.changeDirection(RIGHT)
		case "b":
			selectedRabbits = []rabbit{}
			return m, nil
		}
	
	case TickMsg:
		return m.moveRabbits()
	}

	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder
	sb.WriteString(RenderTitle())
	sb.WriteString("\n")

	var stringArena strings.Builder
	RenderArena(&m)
	RenderRabbits(&m)
	RenderFood(&m)

	for _, row := range m.arena {
		stringArena.WriteString(strings.Join(row, "") + "\n")
	}

	sb.WriteString(stringArena.String())
	sb.WriteString("\n")
	sb.WriteString(RenderScore(m.score))
	sb.WriteString("\n")

	sb.WriteString(RenderHelp(HELP))
	sb.WriteString("\n")
	sb.WriteString(RenderHelp(QUIT))

	return sb.String()
}
