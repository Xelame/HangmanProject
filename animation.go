package main

import (
	"fmt"
	"time"
)

const WIN_TEXT = "WWWWWWWW                           WWWWWWWW     EEEEEEEEEEEEEEEEEEEEEE     LLLLLLLLLLL                  LLLLLLLLLLL                            PPPPPPPPPPPPPPPPP        LLLLLLLLLLL                                 AAA                    YYYYYYY       YYYYYYY     EEEEEEEEEEEEEEEEEEEEEE     DDDDDDDDDDDDD                                                                                                                       \nW::::::W                           W::::::W     E::::::::::::::::::::E     L:::::::::L                  L:::::::::L                            P::::::::::::::::P       L:::::::::L                                A:::A                   Y:::::Y       Y:::::Y     E::::::::::::::::::::E     D::::::::::::DDD                                                                                                                    \nW::::::W                           W::::::W     E::::::::::::::::::::E     L:::::::::L                  L:::::::::L                            P::::::PPPPPP:::::P      L:::::::::L                               A:::::A                  Y:::::Y       Y:::::Y     E::::::::::::::::::::E     D:::::::::::::::DD                                                                                                                  \nW::::::W                           W::::::W     EE::::::EEEEEEEEE::::E     LL:::::::LL                  LL:::::::LL                            PP:::::P     P:::::P     LL:::::::LL                              A:::::::A                 Y::::::Y     Y::::::Y     EE::::::EEEEEEEEE::::E     DDD:::::DDDDD:::::D                                                                                                                 \n W:::::W           WWWWW           W:::::W        E:::::E       EEEEEE       L:::::L                      L:::::L                                P::::P     P:::::P       L:::::L                               A:::::::::A                YYY:::::Y   Y:::::YYY       E:::::E       EEEEEE       D:::::D    D:::::D                                                                                                                \n  W:::::W         W:::::W         W:::::W         E:::::E                    L:::::L                      L:::::L                                P::::P     P:::::P       L:::::L                              A:::::A:::::A                  Y:::::Y Y:::::Y          E:::::E                    D:::::D     D:::::D                                                                                                               \n   W:::::W       W:::::::W       W:::::W          E::::::EEEEEEEEEE          L:::::L                      L:::::L                                P::::PPPPPP:::::P        L:::::L                             A:::::A A:::::A                  Y:::::Y:::::Y           E::::::EEEEEEEEEE          D:::::D     D:::::D                                                                                                               \n    W:::::W     W:::::::::W     W:::::W           E:::::::::::::::E          L:::::L                      L:::::L                                P:::::::::::::PP         L:::::L                            A:::::A   A:::::A                  Y:::::::::Y            E:::::::::::::::E          D:::::D     D:::::D                                                                                                               \n     W:::::W   W:::::W:::::W   W:::::W            E:::::::::::::::E          L:::::L                      L:::::L                                P::::PPPPPPPPP           L:::::L                           A:::::A     A:::::A                  Y:::::::Y             E:::::::::::::::E          D:::::D     D:::::D                                                                                                               \n      W:::::W W:::::W W:::::W W:::::W             E::::::EEEEEEEEEE          L:::::L                      L:::::L                                P::::P                   L:::::L                          A:::::AAAAAAAAA:::::A                  Y:::::Y              E::::::EEEEEEEEEE          D:::::D     D:::::D                                                                                                               \n       W:::::W:::::W   W:::::W:::::W              E:::::E                    L:::::L                      L:::::L                                P::::P                   L:::::L                         A:::::::::::::::::::::A                 Y:::::Y              E:::::E                    D:::::D     D:::::D                                                                                                               \n        W:::::::::W     W:::::::::W               E:::::E       EEEEEE       L:::::L         LLLLLL       L:::::L         LLLLLL                 P::::P                   L:::::L         LLLLLL         A:::::AAAAAAAAAAAAA:::::A                Y:::::Y              E:::::E       EEEEEE       D:::::D    D:::::D                                                                                                                \n         W:::::::W       W:::::::W              EE::::::EEEEEEEE:::::E     LL:::::::LLLLLLLLL:::::L     LL:::::::LLLLLLLLL:::::L               PP::::::PP               LL:::::::LLLLLLLLL:::::L        A:::::A             A:::::A               Y:::::Y            EE::::::EEEEEEEE:::::E     DDD:::::DDDDD:::::D                                                                                                                 \n          W:::::W         W:::::W               E::::::::::::::::::::E     L::::::::::::::::::::::L     L::::::::::::::::::::::L               P::::::::P               L::::::::::::::::::::::L       A:::::A               A:::::A           YYYY:::::YYYY         E::::::::::::::::::::E     D:::::::::::::::DD                                                                                                                  \n           W:::W           W:::W                E::::::::::::::::::::E     L::::::::::::::::::::::L     L::::::::::::::::::::::L               P::::::::P               L::::::::::::::::::::::L      A:::::A                 A:::::A          Y:::::::::::Y         E::::::::::::::::::::E     D::::::::::::DDD                                                                                                                    \n            WWW             WWW                 EEEEEEEEEEEEEEEEEEEEEE     LLLLLLLLLLLLLLLLLLLLLLLL     LLLLLLLLLLLLLLLLLLLLLLLL               PPPPPPPPPP               LLLLLLLLLLLLLLLLLLLLLLLL     AAAAAAA                   AAAAAAA         YYYYYYYYYYYYY         EEEEEEEEEEEEEEEEEEEEEE     DDDDDDDDDDDDD                                                                                                                       \n"

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
	for indexVisible := 0; indexVisible <= len(listxt[0]); indexVisible++ {
		time.Sleep(50 * time.Millisecond)
		Clear()
		if numberOfSpace <= 0 {
			indexWrap++
		}
		for _, line := range listxt[:(len(listxt) - 1)] {
			numberOfSpace = len(line) - 307 - indexVisible // -307 pour la valeur pile ou l'animation disparait
			for i := 0; i <= numberOfSpace; i++ {
				fmt.Print(" ")
			}
			if numberOfSpace >= 0 {
				fmt.Print(line[:indexVisible])
			} else {
				fmt.Print(line[indexWrap:indexVisible])
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("\033[0m")
	Clear()
}
