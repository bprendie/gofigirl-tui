package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type preset struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var videoStyles = []string{"Lofi Video", "HD Video"}

type model struct {
	presets      []preset
	videoStyles  []string
	cursor       int
	vsCursor     int
	textInput    textinput.Model
	ctx          context.Context
	cancel       context.CancelFunc
	err          error
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func initialModel() model {
	// Read presets from file
	file, err := os.ReadFile("presets.json")
	if err != nil {
		fmt.Println("Error reading presets.json:", err)
		os.Exit(1)
	}

	var presets []preset
	err = json.Unmarshal(file, &presets)
	if err != nil {
		fmt.Println("Error unmarshalling presets.json:", err)
		os.Exit(1)
	}

	ti := textinput.New()
	ti.Placeholder = "Paste a YouTube URL here"
	ti.CharLimit = 156
	ti.Width = 50

	ctx, cancel := context.WithCancel(context.Background())

	return model{
		presets:     presets,
		videoStyles: videoStyles,
		vsCursor:    0, // "Lofi Video" is the default
		textInput:   ti,
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func playStream(ctx context.Context, url string, videoStyle string) tea.Cmd {
	return func() tea.Msg {
		logFile, err := os.OpenFile("gofigirl_errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errMsg{fmt.Errorf("error opening log file: %w", err)}
		}
		defer logFile.Close()

		var vo string
		if videoStyle == "Lofi Video" {
			vo = "tct"
		}

		var ytDlpErr bytes.Buffer
		c := exec.CommandContext(ctx, "yt-dlp", "-f", "b", "-g", url)
		c.Stderr = io.MultiWriter(logFile, &ytDlpErr)
		output, err := c.Output()
		if err != nil {
			if strings.Contains(ytDlpErr.String(), "Forbidden") {
				return nil
			}
			return errMsg{fmt.Errorf("error getting stream URL: %w\n%s", err, ytDlpErr.String())}
		}

		streamURL := strings.TrimSpace(string(output))
		var cmd *exec.Cmd
		if vo != "" {
			cmd = exec.CommandContext(ctx, "mpv", fmt.Sprintf("--vo=%s", vo), streamURL)
		} else {
			cmd = exec.CommandContext(ctx, "mpv", streamURL)
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = logFile
		if err := cmd.Run(); err != nil {
			return errMsg{fmt.Errorf("error running mpv: %w", err)}
		}
		return nil
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.cancel()
			return m, tea.Quit
		case "up", "k":
			if m.textInput.Focused() {
				m.textInput.Blur()
				m.cursor = len(m.presets) - 1
			} else if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor == len(m.presets)-1 {
				m.textInput.Focus()
			} else {
				m.cursor++
			}
		case "left", "h":
			if m.vsCursor > 0 {
				m.vsCursor--
			}
		case "right", "l":
			if m.vsCursor < len(m.videoStyles)-1 {
				m.vsCursor++
			}
		case "enter":
			if m.textInput.Focused() {
				return m, playStream(m.ctx, m.textInput.Value(), m.videoStyles[m.vsCursor])
			}
			return m, playStream(m.ctx, m.presets[m.cursor].URL, m.videoStyles[m.vsCursor])
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			PaddingLeft(2).
			PaddingRight(2)
	cursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	vsCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	inactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

func (m model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	s := titleStyle.Render("GoFigirl TUI")
	s += "\n\n"

	for i, preset := range m.presets {
		cursor := " "
		if m.cursor == i && !m.textInput.Focused() {
			cursor = ">"
			s += cursorStyle.Render(fmt.Sprintf("%s %s", cursor, preset.Name))
		} else {
			s += fmt.Sprintf("%s %s", cursor, preset.Name)
		}
		s += "\n"
	}

	s += "\nVideo Style: "
	vs_s := ""
	for i, vs := range m.videoStyles {
		if i == m.vsCursor {
			vs_s += vsCursorStyle.Render(vs)
		} else {
			vs_s += inactiveStyle.Render(vs)
		}
		if i < len(m.videoStyles)-1 {
			vs_s += " | "
		}
	}
	s += vs_s

	s += "\n\n" + m.textInput.View()
	s += "\n\nPress 'q' to quit.\n"

	return s
}

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		if err := playStream(context.Background(), url, "Lofi Video")(); err != nil {
			fmt.Printf("Error playing stream: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	m := initialModel()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
