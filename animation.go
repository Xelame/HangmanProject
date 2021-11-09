package main

import (
	"fmt"
	"time"
)

const WIN_TEXT = `YYYYYYY       YYYYYYY          OOOOOOOOO          UUUUUUUU     UUUUUUUU               WWWWWWWW                           WWWWWWWW     IIIIIIIIII     NNNNNNNN        NNNNNNNN
Y:::::Y       Y:::::Y        OO:::::::::OO        U::::::U     U::::::U               W::::::W                           W::::::W     I::::::::I     N:::::::N       N::::::N
Y:::::Y       Y:::::Y      OO:::::::::::::OO      U::::::U     U::::::U               W::::::W                           W::::::W     I::::::::I     N::::::::N      N::::::N
Y::::::Y     Y::::::Y     O:::::::OOO:::::::O     UU:::::U     U:::::UU               W::::::W                           W::::::W     II::::::II     N:::::::::N     N::::::N
YYY:::::Y   Y:::::YYY     O::::::O   O::::::O      U:::::U     U:::::U                 W:::::W           WWWWW           W:::::W        I::::I       N::::::::::N    N::::::N
   Y:::::Y Y:::::Y        O:::::O     O:::::O      U:::::D     D:::::U                  W:::::W         W:::::W         W:::::W         I::::I       N:::::::::::N   N::::::N
    Y:::::Y:::::Y         O:::::O     O:::::O      U:::::D     D:::::U                   W:::::W       W:::::::W       W:::::W          I::::I       N:::::::N::::N  N::::::N
     Y:::::::::Y          O:::::O     O:::::O      U:::::D     D:::::U                    W:::::W     W:::::::::W     W:::::W           I::::I       N::::::N N::::N N::::::N
      Y:::::::Y           O:::::O     O:::::O      U:::::D     D:::::U                     W:::::W   W:::::W:::::W   W:::::W            I::::I       N::::::N  N::::N:::::::N
       Y:::::Y            O:::::O     O:::::O      U:::::D     D:::::U                      W:::::W W:::::W W:::::W W:::::W             I::::I       N::::::N   N:::::::::::N
       Y:::::Y            O:::::O     O:::::O      U:::::D     D:::::U                       W:::::W:::::W   W:::::W:::::W              I::::I       N::::::N    N::::::::::N
       Y:::::Y            O::::::O   O::::::O      U::::::U   U::::::U                        W:::::::::W     W:::::::::W               I::::I       N::::::N     N:::::::::N
       Y:::::Y            O:::::::OOO:::::::O      U:::::::UUU:::::::U                         W:::::::W       W:::::::W              II::::::II     N::::::N      N::::::::N
    YYYY:::::YYYY          OO:::::::::::::OO        UU:::::::::::::UU                           W:::::W         W:::::W               I::::::::I     N::::::N       N:::::::N
    Y:::::::::::Y            OO:::::::::OO            UU:::::::::UU                              W:::W           W:::W                I::::::::I     N::::::N        N::::::N
    YYYYYYYYYYYYY              OOOOOOOOO                UUUUUUUUU                                 WWW             WWW                 IIIIIIIIII     NNNNNNNN         NNNNNNN`

var winText = Split(WIN_TEXT, "\n")

func Split(s, sep string) []string {
	arg := []string{}
	begin := -1
	for index, value := range s {
		if value == rune(sep[0]) && s[index:index+len(sep)] == sep {
			if s[begin+1:index] != "" {
				if begin != -1 {
					arg = append(arg, s[begin+len(sep):index])
				} else {
					arg = append(arg, s[begin+1:index])
				}
			}
			begin = index
		}
	}
	arg = append(arg, s[begin+len(sep):])
	return arg
}

func Animation(listxt []string) {
	fmt.Print("\033[32m")
	numberOfSpace := 0
	indexWrap := 1
	for indexVisible := 0; indexWrap < len(listxt[0]); indexVisible++ {
		time.Sleep(50 * time.Millisecond)
		Clear()
		if numberOfSpace <= 0 {
			indexWrap++
		}
		for _, line := range listxt[:(len(listxt) - 1)] {
			if indexVisible < len(line) {
				numberOfSpace = len(line)/2 - indexVisible
				for i := 0; i <= numberOfSpace; i++ {
					fmt.Print(" ")
				}
				if numberOfSpace >= 0 {
					fmt.Print(line[:indexVisible])
				} else {
					fmt.Print(line[indexWrap:indexVisible])
				}
			} else {
				fmt.Print(line[indexWrap:])
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("\033[0m")
	Clear()
}
