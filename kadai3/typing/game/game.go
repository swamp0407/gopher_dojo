package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Game struct {
	name  string
	score int
}

func start(r io.Reader) <-chan struct{} {
	ch := make(chan struct{})
	fmt.Println("Press any key to start")
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- struct{}{}
			fmt.Println("Start!")
			break
		}
		close(ch)
	}()
	return ch
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func (g *Game) Run() {
	sch := start(os.Stdin)
	<-sch

	ch := input(os.Stdin)
	words, err := getWords()
	if err != nil {
		panic(err)
	}
	totalAnsCount := 0
	tch := time.NewTimer(time.Second * 60)
Outer:
	for _, word := range words {
		fmt.Printf("type: %s\n", word)
	Inner:
		for {
			select {
			case <-tch.C:
				fmt.Println("Time up!")
				break Outer
			case s := <-ch:
				totalAnsCount++
				if s == word {
					g.score++
					break Inner
				}
			}
		}
	}
	fmt.Println("your total challenge count is", totalAnsCount)
	fmt.Println("Your score is", g.score)
	return
}

func getWords() ([]string, error) {

	f, err := os.Open("words.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make(map[string]struct{})

	words := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines[s.Text()] = struct{}{}
	}
	count := 0
	for k := range lines {
		words = append(words, k)
		count++
		if count == 30 {
			break
		}
	}
	return words, nil
}
