package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 4

type Game struct {
	board [size][size]int
	score int
}

/*
This function does some setup for the game. It creates a seed for the random value to add, then adds it
@param: pointer to the an instance of the game struct
@return: none
*/
func (g *Game) init() {
	rand.Seed(time.Now().UnixNano())
	g.addValue()
}

/*
This function iterates through each of the values and determines whether there exists a valid move
@param: pointer to the an instance of the game struct
@return: bool (true if there are no valid moves and false otherwise)
*/
func(g *Game) isGameOver() bool{
	for i := 0; i < size; i++{
		for j := 0; j < size; j++{
			if(g.board[i][j] == 0) {
				return false
			}
			if i > 0 && g.board[i][j] == g.board[i-1][j] {
				return false
			}
			if j > 0 && g.board[i][j] == g.board[i][j-1] {
				return false
			}
		}
	}
	return true
}
/*
This function uses the random seed to place a new value - either a 2 or 4 in an unoccupied spot
@param: pointer to the an instance of the game struct
@return: none
*/
func(g *Game) addValue() {
	emptyCells := []struct{ x, y int }{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g.board[i][j] == 0 {
				emptyCells = append(emptyCells, struct{ x, y int }{i, j})
			}
		}
	}
	if len(emptyCells) == 0 {
		return
	}
	cell := emptyCells[rand.Intn(len(emptyCells))]
	g.board[cell.x][cell.y] = 2 * (rand.Intn(2) + 1)
}
/*
This function simply displays the board, which is being stored in the g.board variable
@param: pointer to the an instance of the game struct
@return: none
*/
func(g *Game) displayBoard(){
	fmt.Println("Score:", g.score)
	fmt.Println("-----------------")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {	
			if g.board[i][j] == 0 {
				fmt.Print(". \t")
			} else {
				fmt.Printf("%d \t", g.board[i][j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------")	
}
/*
This function does all of the logic for moving the board, combining like values, and calculating the score
@param: pointer to the an instance of the game struct, and the direction the user would like to move
@return: none
*/
func(g *Game) move(direction string) {
	if direction == "left" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {	
				if g.board[i][j] != 0 {  
					k := j
					//move the value as far as it can go without hitting a wall or another number
					for k > 0 && g.board[i][k-1] == 0 {
						g.board[i][k-1] = g.board[i][k]
						g.board[i][k] = 0
						k--  
					}
					
					if k > 0 && g.board[i][k-1] == g.board[i][k] {
						g.board[i][k-1] *= 2 
						g.board[i][k] = 0  
						g.score += g.board[i][k-1] 
					}
				}
			}
		}
	}
	if direction == "right" {
		for i := 0; i < size; i++ {
			for j := size - 2; j >= 0; j-- {
				if g.board[i][j] != 0 {
					k := j
					for k < size-1 && g.board[i][k+1] == 0 {
						g.board[i][k+1] = g.board[i][k]
						g.board[i][k] = 0
						k++ 
					}
					if k < size-1 && g.board[i][k+1] == g.board[i][k] {
						g.board[i][k+1] *= 2
						g.board[i][k] = 0
						g.score += g.board[i][k+1]
					}
				}
			}
		}
	}
	
	if direction == "up" {
		for j := 0; j < size; j++ {
			for i := 1; i < size; i++ {
				if g.board[i][j] != 0 {
					k := i
					for k > 0 && g.board[k-1][j] == 0 {
						g.board[k-1][j] = g.board[k][j]
						g.board[k][j] = 0
						k--
					}
					if k > 0 && g.board[k-1][j] == g.board[k][j] {
						g.board[k-1][j] *= 2 
						g.board[k][j] = 0
						g.score += g.board[k-1][j]
					}
				}
			}
		}
	}
	
	if direction == "down" {
		for j := 0; j < size; j++ {
			for i := size - 2; i >= 0; i-- {
				if g.board[i][j] != 0 {
					k := i
					for k < size-1 && g.board[k+1][j] == 0 {
						g.board[k+1][j] = g.board[k][j]
						g.board[k][j] = 0
						k++
					}
					if k < size-1 && g.board[k+1][j] == g.board[k][j] {
						g.board[k+1][j] *= 2 
						g.board[k][j] = 0 
						g.score += g.board[k+1][j] 
					}
				}
			}
		}
	}
}
/*
Calls the init function, calls the functions, and recieves input from user
@param: none
@return: none
*/
func main() {
	game := Game{}
	game.init()
	for {
		//check if game is over
		if game.isGameOver() {
			fmt.Println("Game Over. Final Score:", game.score)
			break
		}
		//if it isnt over, add random value - 2 or 4
		game.addValue();
		//then display the board and score
		game.displayBoard();
		//get input from user and move the tiles
		fmt.Print("Enter move (w=up, s=down, a=left, d=right): ")
		var move string
		fmt.Scanln(&move)

		switch move {
		case "w":
		game.move("up")
		case "s":
		game.move("down")
		case "a":
		game.move("left")
		case "d":
		game.move("right")
		default:
			fmt.Println("Invalid move. Please enter w, a, s, or d.")
			continue
		}
	}
}