package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cwiggins0217/go-todo/internal"
)

//TODO investigate go's fucked up way to make enums

type CmdFlags struct {
	AddFlag    string
	DelFlag    int
	EditFlag   string
	ToggleFlag int
	ListFlag   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.AddFlag, "add", "", "Add a new todo item")
	flag.IntVar(&cf.DelFlag, "delete", -1, "delete a specified todo item")
	flag.StringVar(&cf.EditFlag, "edit", "", "edit a specified todo item")
	flag.IntVar(&cf.ToggleFlag, "toggle", -1, "toggle completion status of a specified todo item")
	flag.BoolVar(&cf.ListFlag, "list", false, "list all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *internal.Todos) {
	switch {

	case cf.ListFlag:
		todos.Print()

	case cf.AddFlag != "":
		todos.Add(cf.AddFlag)

	case cf.EditFlag != "":
		parts := strings.SplitN(cf.EditFlag, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Use ID:Title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("error: invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])

	case cf.ToggleFlag != -1:
		todos.Toggle(cf.ToggleFlag)

	case cf.DelFlag != -1:
		todos.Delete(cf.DelFlag)
	}
}
