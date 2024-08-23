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
	rabbit rabbit
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

func InitialModel() Model {
	return Model{
		rabbit: rabbit{
			position: coord{x: 2, y: 2},
			direction: DOWN,
		},
		width: 60,
		height: 20,
		arena: [][]string{},
		food: coord{x: 10, y: 10},
		score: 0,
		impassableSymbol: "#",
		foodSymbol: "*",
		rabbitSymbol: "@",
		verticalEdgeSymbol: "|",
		horizontalEdgeSymbol: "-",
		emptySymbol: " ",
	}
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Duration(INTERVAL)*time.Millisecond, func(t time.Time) tea.Msg{
		return TickMsg(t)
	})
}

func (m Model) changeDirection(direction int) (tea.Model, tea.Cmd) {

	m.rabbit.direction = direction

	return m, nil
}

func (m Model) Init() tea.Cmd {
	var x, y int

	x = rand.Intn(m.height - 1)
	y = rand.Intn(m.width - 1)

	m.food = coord{x: x, y: y}
	return m.tick()
}

func (m Model) moveRabbit() (tea.Model, tea.Cmd) {
	pos := coord{x: m.rabbit.position.x, y: m.rabbit.position.y}

	switch m.rabbit.direction {
	case UP:
		pos.x--
	case DOWN:
		pos.x++
	case LEFT:
		pos.y--
	case RIGHT:
		pos.y++
	}

	if pos.x == m.food.x && pos.y == m.food.y {
		m.food.x = rand.Intn(m.height-2) + 1
		m.food.y = rand.Intn(m.width-2) + 1
	}

	m.rabbit.position = pos
	return m, m.tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
		}
	
	case TickMsg:
		return m.moveRabbit()
	}

	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder
	sb.WriteString(RenderTitle())
	sb.WriteString("\n")

	var stringArena strings.Builder
	RenderArena(&m)
	RenderRabbit(&m)
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
