package main

import (
	"bufio"
	"fmt"
	bb "main/branchbound"
	gd "main/greedy"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	algoUsed = "gd"
	qtd_exec = 1000
)

func main() {
	defer fmt.Println("Fim do Programa")

	var deadlines []int
	var penalties []int
	inputSize, err := strconv.Atoi(os.Args[1])
	algoUsed := os.Args[2]

	if err != nil {
		fmt.Println("Erro ao converter o tamanho da entrada")
		os.Exit(1)
	}
	// Abre os arquivos e Lê

	file, err := os.Open(fmt.Sprintf("inputs/input%d.txt", inputSize))
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		penalty, _ := strconv.Atoi(nums[0])
		deadline, _ := strconv.Atoi(nums[1])
		deadlines = append(deadlines, deadline)
		penalties = append(penalties, penalty)
	}

	if inputSize < 50 {
		fmt.Println("Penalties:", penalties)
		fmt.Println("Deadlines:", deadlines)
	}
	penalty := 0

	fmt.Print("Usando o algoritmo: ")

	if algoUsed == "bb" {
		fmt.Println("Branch and Bound")
		penalty = execBranchBound(deadlines, penalties)
	} else if algoUsed == "gd" {
		fmt.Println("Greedy")
		penalty = execGreedy(deadlines, penalties)
	}
	fmt.Println("Penalidade:", penalty)

}

func execGreedy(deadlines []int, penalties []int) int {
	intiTime := time.Now()
	penalty := 0

	for i := 0; i < qtd_exec; i++ {
		penalty = gd.GetJobs(deadlines, penalties)
	}
	finalTime := time.Now()
	fmt.Println("Tempo de execução:", finalTime.Sub(intiTime))
	return penalty
}

func execBranchBound(deadlines []int, penalties []int) int {
	intiTime := time.Now()
	penalty := 0

	for i := 0; i < qtd_exec; i++ {
		penalty = bb.GetJobs(deadlines, penalties)
	}
	finalTime := time.Now()
	fmt.Println("Tempo de execução:", finalTime.Sub(intiTime))
	return penalty
}
