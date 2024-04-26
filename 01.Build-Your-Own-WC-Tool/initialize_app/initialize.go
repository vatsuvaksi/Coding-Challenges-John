package initializeapp

import (
	"fmt"
	"wc-tool/service"
)

func InitialzeApplication() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                  WELCOME TO WC TOOLS                           ║")
	fmt.Println("║    The Application gives you complete knowledge of files       ║")
	fmt.Println("║    Currently, the application supports the following commands: ║")
	fmt.Println("║    [-c -l -w  -m (default)]                                    ║")
	fmt.Println("║    Please enter the command in the following format:           ║")
	fmt.Println("║    wc -[flag] txt_file_path                                    ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║                            Enjoy!                              ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║                                      - Vatsal       			  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
}

func Start() {
	// Parse Input
	input, err := service.FetchInput()
	if err != nil {
		fmt.Print("err-> ", err)
	}
	fmt.Println(service.CommandFactory(input)(input))
}
