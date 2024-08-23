package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderArena(m *Model) {
	m.arena = append(
		m.arena,
		strings.Split(
			m.verticalEdgeSymbol+strings.Repeat(m.horizontalEdgeSymbol, m.width-2)+m.verticalEdgeSymbol,
			"",
		),
	)

	for i := 0; i < m.height-1; i++ {
		m.arena = append(
			m.arena,
			strings.Split(
				m.verticalEdgeSymbol+strings.Repeat(m.emptySymbol, m.width-2)+m.verticalEdgeSymbol,
				"",
			),
		)
	}

	m.arena = append(
		m.arena,
		strings.Split(
			m.verticalEdgeSymbol+strings.Repeat(m.horizontalEdgeSymbol, m.width-2)+m.verticalEdgeSymbol,
			"",
		),
	)
}

func RenderRabbit(m *Model) {
	m.arena[m.rabbit.position.x][m.rabbit.position.y] = m.rabbitSymbol
}

func RenderFood(m *Model) {
	m.arena[m.food.x][m.food.y] = m.foodSymbol
}

func RenderTitle() string {
	ts := lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("202")).
		Background(lipgloss.Color("#3C3C3C")).
		Width(60).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Underline(true)
	return ts.Render("go rabbits!")
}

func RenderScore(score int) string {
	scoreStr := fmt.Sprintf("Score: %d ", score)
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))

	return ts.Render(scoreStr)
}

func RenderHelp(h string) string {
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3C3C3C"))

	return ts.Render(h)
}

func RenderQuitcommand() string {
	qc := "Press ctrl+c to quit"
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3C3C3C"))
	return ts.Render((qc))
}

func RenderGameOver() string {
	return lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#FF0000")).
		Width(40).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Render("Game Over!")
}
