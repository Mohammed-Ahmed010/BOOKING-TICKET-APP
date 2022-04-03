package main
import 
(
"fmt"
"strings"
"booking-app/helper"
"time"
"sync"
)
 

var conferenceName = "Go conference1"
const conferenceTickets int =50
var remainingTickets uint =50
var bookings =make([]UserData,0)
var wg =sync.WaitGroup{}
type UserData struct{
	firstName string
	lastName string
	email string
	userTicket uint
	
}
func main()  {
    

	helper.GreetUser()
	firstName,lastName,email,userTicket:=getUserInput()
   
	isValidName, isValidEmail,isValidTicketNUmber:=  validateUserInput(firstName,lastName,email,userTicket)
	if isValidEmail==true && isValidTicketNUmber==true && isValidName==true{
		
	if remainingTickets==0{
		fmt.Println("sorry ticket are over")
		
	}
	remainingTickets=remainingTickets-userTicket
	bookingTickets(userTicket,firstName,lastName,email)
	wg.Add(1)
	go sendTicket(userTicket,firstName,lastName,email)
	fmt.Printf("%v tickets are remaing in %v conference \n",remainingTickets,conferenceName)
	printFirstNames()
	}else{
		fmt.Println("enter details correctly")
	}

 
wg.Wait()
}




func printFirstNames(){
firstName:=[]string{}
for _,booking :=range bookings{
	//var names = strings.Fields(booking)

	firstName =append(firstName,booking.firstName)

}
fmt.Printf("THE first names of booking are %v\n",firstName)
}
func validateUserInput(firstName string,lastName string,email string,userTicket uint)(bool,bool,bool){
	isValidName := len(firstName)>=2 && len(lastName)>=2
    isValidEmail := strings.Contains(email,"@")
	isValidTicketNUmber := userTicket<=remainingTickets && userTicket>0


	return isValidName==true,  isValidEmail==true ,isValidTicketNUmber==true 
}
func getUserInput()(string,string,string,uint){
	var firstName string 
    var lastName  string 
	var email  string 
	//var booking = [50] string{"AHMED","MOHAMMED","JAMAL"}
	var userTicket uint
	fmt.Println("Please enter first name ")
	fmt.Scan(&firstName)
	
	fmt.Println("Please enter last name ")
	fmt.Scan(&lastName)
	//booking[0]=firstName +" "+lastName
	fmt.Println("Please email ")
	fmt.Scan(&email)
	fmt.Println("enter number of tickets ")
    fmt.Scan(&userTicket)
	return firstName,lastName,email,userTicket
}

func bookingTickets(userTicket uint,firstName string,lastName string,email string){
	
	//creating map
	var userData =UserData{
		firstName:firstName,
		lastName :lastName,
		email:email,
		userTicket:userTicket,
	}
	
	fmt.Printf("Details of user %v",userData)
	bookings=append(bookings,userData)
	fmt.Printf(" thank you User %v %v has booked %d tickets ,you will recieve a conformation email at %v\n",firstName,lastName,userTicket,email)
	
}
func sendTicket(userTicket uint,firstName string,lastName string,email string) {
time.Sleep(10 * time.Second)
 var ticket =fmt.Sprintf("%v tickets for %v %v ",userTicket,firstName,lastName)
 fmt.Println("************************************************************")
 fmt.Printf("Sending ticket : %v \n to email :%v \n",ticket,email)
 fmt.Println("************************************************************")
 wg.Done()
}