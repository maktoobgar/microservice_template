package load

import (
	"fmt"
	"log"
	g "service/auth/global"
	"service/pkg/colors"
)

func Info() {
	fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sSystem Info%s==%s\n", colors.Yellow, colors.Cyan, colors.Reset))
	fmt.Printf("Name:\t\t%s%s%s\n", colors.Blue, g.Name, colors.Reset)
	fmt.Printf("Version:\t%s%s%s\n", colors.Blue, g.Version, colors.Reset)
	mainOrTest := "test"
	mainOrTestColor := colors.Red + mainOrTest + colors.Reset
	if !g.CFG.Debug {
		mainOrTest = "main"
		mainOrTestColor = colors.Green + mainOrTest + colors.Reset
	}
	for name, database := range g.CFG.CurrentMicroservice.Databases {
		if name == mainOrTest {
			if database.Type == "sqlite3" {
				fmt.Printf("Main Database:\t%v, %v (%v)\n", database.Type, database.DbName, mainOrTestColor)
			} else {
				fmt.Printf("Main Database:\t%v, %v, %v:%v (%v)\n", database.Type, database.DbName, database.Host, database.Port, mainOrTestColor)
			}
			if g.DB == nil {
				log.Fatal("expected database connection is not assigned as main database")
			}
			break
		}
	}
	if g.CFG.Debug {
		fmt.Printf("Debug:\t\t%s%v%s\n", colors.Red, g.CFG.Debug, colors.Reset)
	} else {
		fmt.Printf("Debug:\t\t%s%v%s\n", colors.Green, g.CFG.Debug, colors.Reset)
	}
	fmt.Printf("Address:\tgrpc://%s:%s\n", g.CFG.CurrentMicroservice.IP, g.CFG.CurrentMicroservice.Port)
	fmt.Print(colors.Cyan, "===============\n\n", colors.Reset)
}
