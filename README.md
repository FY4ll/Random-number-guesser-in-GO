Random Number Guesser
=====================

This code in Go is a simple guessing game where the user has to guess a random number between 1 and 10 in a limited number of attempts.

The code starts by importing the necessary packages: "fmt" for printing to the console, "math/rand" for generating random numbers and "time" for seeding the random number generator with the current time.

The main function initializes a random number using rand.Intn() and sets the number of attempts to 3. It then enters a loop where it prompts the user to enter a number between 1 and 10 using fmt.Scanln(). If the user input is invalid, i.e., 0 or greater than 10, the loop breaks.

If the user input is the same as the random number, the game ends and the user is congratulated for winning with the number of remaining attempts displayed. The user is then prompted to restart the game or exit. If the user chooses to restart, the main function is called again.

If the user input is less than the random number, a message is displayed saying the user's number is too small and the number of remaining attempts is decremented. If the user input is greater than the random number, a message is displayed saying the user's number is too large and the number of remaining attempts is decremented.

If the user runs out of attempts, a message is displayed saying the game is lost and the number of remaining attempts is 0. The user is then prompted to restart the game or exit.

Overall, this code is a simple example of a console-based game that demonstrates the use of basic programming constructs such as loops, conditional statements, and functions.

