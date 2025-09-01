package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	Attack  string = "attack"
	Defence string = "defence"
	Special string = "special"
	Warrior string = "warrior"
	Mage    string = "mage"
	Healer  string = "healer"
)

func action(charName, charClass, actType string) string {
	switch actType {
	case Attack:
		switch charClass {
		case Warrior:
			return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
		case Mage:
			return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
		case Healer:
			return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
		}
	case Defence:
		switch charClass {
		case Warrior:
			return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(5, 10))
		case Mage:
			return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(-2, 2))
		case Healer:
			return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(2, 5))
		}
	case Special:
		switch charClass {
		case Warrior:
			return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
		case Mage:
			return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
		case Healer:
			return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
		}
	}
	return "неизвестный класс персонажа или действие"
}

func describeClass(charName, charClass string) {
	switch charClass {
	case Warrior:
		fmt.Printf("%s, ты Воитель — отличный боец ближнего боя.\n", charName)
	case Mage:
		fmt.Printf("%s, ты Маг — превосходный укротитель стихий.\n", charName)
	case Healer:
		fmt.Printf("%s, ты Лекарь — чародей, способный исцелять раны.\n", charName)
	default:
		fmt.Println("Неизвестный класс.")
	}
}

// здесь обратите внимание на имена параметров
func training(charName, charClass string, r *bufio.Reader) string {
	describeClass(charName, charClass)
	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: ")
	fmt.Println("attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	for {
		fmt.Print("Введи команду: ")
		cmdRaw, _ := r.ReadString('\n')
		cmd := strings.ToLower(strings.TrimSpace(cmdRaw))
		switch cmd {
		case Attack, Defence, Special:
			fmt.Println(action(charName, charClass, cmd))
		case "skip":
			return "тренировка окончена"
		default:
			fmt.Println("Неизвестная команда. Доступно: attack, defence, special, skip.")
		}
	}
}

// обратите внимание на имя функции и имена переменных
func chooseClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Printf("Введи название персонажа, за которого хочешь играть: Воитель — %s, Маг — %s, Лекарь — %s: ", Warrior, Mage, Healer)
		fmt.Scanf("%s\n", &charClass)
		switch charClass {
		case Warrior:
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case Mage:
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case Healer:
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			continue
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	reader := bufio.NewReader(os.Stdin)
	var charName string
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := chooseClass()

	fmt.Println(training(charName, charClass, reader))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
