package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	name string
	number int
}

func parseInstructions(input []string) map[int]instruction {
	out := make(map[int]instruction)
	for line, instr := range input {
		if instr == "" {
			break
		}
		instr := strings.Split(instr, " ")
		name := instr[0]
		number, _ := strconv.Atoi(instr[1])
		out[line] = instruction {
			name: name,
			number: number,
		}
	}
	return out
}

func execute(i map[int]instruction) (int, error) {
	sum := 0
	seen := make(map[int]struct{})
	for line := 0; line<len(i); line++ {
		instr := i[line]
		_, ok := seen[line]; if ok {
			return sum, errors.New("error: Infinite Loop")
		}
		switch instr.name {
		case "nop":
			continue
		case "jmp":
			line += instr.number - 1
		case "acc":
			sum += instr.number
		}
		seen[line] = struct{}{}
	}
	return sum, nil
}

func bugfix(in instruction) instruction {
	out := instruction{}
	if in.name == "jmp" {
		out.name = "nop"
	} else {
		out.name = "jmp"
	}
	out.number = in.number
	return out
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	instructions := strings.Split(sc, "\n")
	i := parseInstructions(instructions)

	// part 1 - accumulator value before any repeated instructions
	out, _ := execute(i)
	fmt.Println("Accumulator's value before any repeated instructions:", out)

	// part 2 - accumulator value after fixing the one bugged jmp to nop
	tmp := make(map[int]instruction, len(i))
	var a int
	for k,v := range i {
		tmp[k] = v
	}
	for line := 0; line<len(i); line++ {
		instr := tmp[line]
		if instr.name == "acc" {
			continue
		}
		tmp[line] = bugfix(instr)
		out, err := execute(tmp); if err == nil {
			fmt.Println(line, instr, tmp[line])
			a = out
			break
		}
		tmp[line] = bugfix(tmp[line])
	}
	fmt.Println("Accumulator's value after bugfix:", a)
}