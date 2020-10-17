package banner

import "fmt"

type Banner struct {
	msg string
}

func GetBanner() *Banner {
	return &Banner{msg: " __                           __                       \n/\\ \\        __               /\\ \\                      \n\\ \\ \\____  /\\_\\      __      \\_\\ \\    __  __   __  __  \n \\ \\ '__`\\ \\/\\ \\   /'_ `\\    /'_` \\  /\\ \\/\\ \\ /\\ \\/\\ \\ \n  \\ \\ \\L\\ \\ \\ \\ \\ /\\ \\L\\ \\  /\\ \\L\\ \\ \\ \\ \\_\\ \\\\ \\ \\_\\ \\\n   \\ \\_,__/  \\ \\_\\\\ \\____ \\ \\ \\___,_\\ \\ \\____/ \\ \\____/\n    \\/___/    \\/_/ \\/___L\\ \\ \\/__,_ /  \\/___/   \\/___/ \n                     /\\____/                           \n                     \\_/__/                            "}
}

func (receiver Banner) PrintBanner() {
	fmt.Println(receiver.msg)
}
