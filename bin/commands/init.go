package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/st-osi/krow/bin/svc"
	"github.com/st-osi/krow/core/env"
	"github.com/st-osi/krow/core/logger"
)

// init krow project
// init command should be able to set the KROW_PATH in .env file which will be loaded by krow.
// it should be creating default files and folders in the path.
type textInputModel struct {
	textInput textinput.Model
	err       error
}

type (
	errMsg error
)

func initialTextInputModel() textInputModel {
	ti := textinput.New()
	ti.Placeholder = "krow"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return textInputModel{
		textInput: ti,
		err:       nil,
	}
}

func (m textInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			executeInitCommand(m.textInput.Value())
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textInputModel) View() string {
	return fmt.Sprintf(
		"Where do you want to initialize your krow project? \n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func initKrowProject() error {
	p := tea.NewProgram(initialTextInputModel())
	if _, err := p.Run(); err != nil {
		logger.Debug("Error occurred while initializing krow project: ", "err", err)
		return err
	}
	return nil
}

// TODO: Will download the sample folder from github repo instead of creating each one one by one
func executeInitCommand(path string) error {
	// if path exists, thats it, otherwise ask if user wants to create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Debug("[log]: Path doesn't exist, creating krow project in: ", "path", path)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			logger.Debug("[error]: Error occurred while creating krow project: ", "err", err)
			return err
		}
		err = svc.CreateDefaultFile(path)
		if err != nil {
			logger.Debug("[error]: Error occurred while creating default files: ", "err", err)
			return err
		}
		env.UpdateEnv(path)
		return nil
	} else {
		env.UpdateEnv(path)
		return nil
	}

}
