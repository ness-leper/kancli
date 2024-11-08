package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

/*
  CUSTOM ITEM
*/

type Task struct {
	status      status
	title       string
	description string
}

// implement the list.item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */

type Model struct {
	list list.Model
	err  error
}

func New() *Model {
	return &Model{}
}

// TODO: call this on tea.WindowSizeMsg
func (m *Model) InitList(width, height int) {
	// width and height are the terminal values
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "To Do"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "strawberry"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Udpate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.InitList(msg.Width, msg.Height)
	}
	// var cmd tea.Cmd
	// m.list, cmd = m.list.Update(msg)
	return m, nil
}

func (m Model) View() string {
	return m.list.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
