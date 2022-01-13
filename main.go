package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings =  make([]UserData, 0)

type UserData struct {
  firstName string
  lastName string
  email string
  numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){

  greetUsers()
  
    firstName, lastName, email, userTickets := getUserInput()

    isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName,userTickets, email, remainingTickets)

    if isValidName && isValidEmail && isValidTicketNumber {
      
      bookTicket(userTickets, firstName,lastName, email)
      wg.Add(1)
      go sendTicket(userTickets, firstName,lastName, email)
      firstNames := getFirstNames()
      fmt.Printf("The first names of bookings are: %v\n", firstNames)
      
      if remainingTickets == 0 {
        //End the program
        fmt.Println("Our conference is booked out. Come back next year.")
      }
    } else {
      if !isValidName{
        fmt.Println("First name or last name you entered is too short.")
      }

      if !isValidEmail {
        fmt.Println("Email address you entered doesn't contain @ sign.")
      }

      if !isValidTicketNumber {
        fmt.Println("Number of tickets you entered is invalid.")
      }
    }
    wg.Wait()
}

func greetUsers(){
  fmt.Printf("Welcome to %v booking application.\n", conferenceName)
  fmt.Printf("We have total of %v tickets and %v are still remaining.\n",conferenceTickets, remainingTickets)
  fmt.Printf("Get your tickets here to attend.\n")
}

func getUserInput()(string, string, string, uint){
    var firstName string
    var lastName string
    var email string
    var userTickets  uint

    // Ask user for their name;
    fmt.Printf("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Printf("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Printf("Enter your email address: ")
    fmt.Scan(&email)

    fmt.Printf("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}


func bookTicket(userTickets uint,firstName string, lastName string, email string){

      remainingTickets = remainingTickets - userTickets

      var userData = UserData {
        firstName : firstName,
        lastName : lastName,
        email : email,
        numberOfTickets: userTickets,
      }
      
      bookings = append(bookings, userData)

      fmt.Printf("List of bookings is %v\n", bookings)
      fmt.Printf("Thank you, %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n",firstName,lastName, userTickets, email)
      fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func getFirstNames()[]string{
  firstNames := []string{}
  for _, booking := range bookings {
    firstNames = append(firstNames, booking.firstName)
  }
  return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
  time.Sleep(10 * time.Second)
  var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
  fmt.Println("#################")
  fmt.Printf("Sending tickets: %v to email address %v.\n", ticket, email)
  fmt.Println("#################")
  wg.Done()
}
