package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	type placeholder [5]string
	zero := placeholder{
		"▮▮▮",
		"▮ ▮",
		"▮ ▮",
		"▮ ▮",
		"▮▮▮",
	}

	one := placeholder{
		"▮▮ ",
		" ▮ ",
		" ▮ ",
		" ▮ ",
		"▮▮▮",
	}

	two := placeholder{
		"▮▮▮",
		"  ▮",
		" ▮ ",
		"▮  ",
		"▮▮▮",
	}

	three := placeholder{
		"▮▮▮",
		"  ▮",
		"▮▮▮",
		"  ▮",
		"▮▮▮",
	}

	four := placeholder{
		"▮ ▮",
		"▮ ▮",
		"▮▮▮",
		"  ▮",
		"  ▮",
	}

	five := placeholder{
		"▮▮▮",
		"▮  ",
		"▮▮▮",
		"  ▮",
		"▮▮▮",
	}

	six := placeholder{
		"▮▮▮",
		"▮  ",
		"▮▮▮",
		"▮ ▮",
		"▮▮▮",
	}

	seven := placeholder{
		"▮▮▮",
		"  ▮",
		"  ▮",
		"  ▮",
		"  ▮",
	}

	eight := placeholder{
		"▮▮▮",
		"▮ ▮",
		"▮▮▮",
		"▮ ▮",
		"▮▮▮",
	}

	nine := placeholder{
		"▮▮▮",
		"▮ ▮",
		"▮▮▮",
		"  ▮",
		"▮▮▮",
	}

	calon := placeholder{
		" ",
		"▮",
		" ",
		"▮",
		" ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			calon,
			digits[min/10], digits[min%10],
			calon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == calon && sec%2 == 0 {
					next = " "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()

		}
		time.Sleep(time.Second)
	}

}
