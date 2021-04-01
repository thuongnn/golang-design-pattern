package main

import "fmt"

const (
	SUCCESS = "success"
	FAILURE = "failure"
	INVALID = "invalid"
	EXPIRED = "expired"
)

type Observer interface {
	update(user User)
}

type Subject interface {
	attach(observer Observer)
	detach(observer Observer)
	notifyAllObserver()
}

type User struct {
	email  string
	ip     string
	status string
}

/*------------------ setup account service ------------------*/
type AccountService struct {
	user      User
	observers []Observer // init list observer here for register
}

func NewAccountService(email, ip string) *AccountService {
	return &AccountService{
		user:      User{email: email, ip: ip},
		observers: []Observer{},
	}
}

func (as *AccountService) attach(observer Observer) {
	as.observers = append(as.observers, observer)
}

func (as *AccountService) detach(observer Observer) {
	var newOb []Observer
	for _, ob := range as.observers {
		if ob != observer {
			newOb = append(newOb, ob)
		}
	}
}

func (as *AccountService) notifyAllObserver() {
	for _, ob := range as.observers {
		ob.update(as.user)
	}
}

func (as *AccountService) changeStatus(status string) {
	as.user.status = status
	fmt.Println("Status is changed")
	as.notifyAllObserver()
}

func (as *AccountService) login() {
	if !as.isValidIP() {
		as.user.status = INVALID
	} else if as.isValidEmail() {
		as.user.status = SUCCESS
	} else {
		as.user.status = FAILURE
	}

	fmt.Println("Login is handled")
	as.notifyAllObserver()
}

func (as *AccountService) isValidIP() bool {
	return as.user.ip != "127.0.0.1"
}

func (as *AccountService) isValidEmail() bool {
	return as.user.email == "thuongnn1997@gmail.com"
}

/*------------------ setup observers ------------------*/
type Logger struct{}
func (Logger) update(user User) {
	fmt.Printf("Logger: ")
	fmt.Println(user)
}

type Mailer struct{}
func (Mailer) update(user User) {
	if user.status == EXPIRED {
		fmt.Printf("Mailer: User %s  is expired. An email was sent!\n", user.email)
	}
}

type Protector struct{}
func (Protector) update(user User) {
	if user.status == INVALID {
		fmt.Printf("Protector: User %s  is invalid.\n", user.email)
		fmt.Printf("           IP %s is blocked.\n", user.ip)
	}
}

/*------------------ running ------------------*/
func main() {
	account1 := createAccount("thuongnn1997@gmail.com", "127.0.0.1")
	account1.login()
	account1.changeStatus(EXPIRED)

	fmt.Println("---")
	account2 := createAccount("thuongnn1997@gmail.com", "192.168.191.130")
	account2.login()
}

func createAccount(email, ip string) *AccountService {
	sa := NewAccountService(email, ip)

	// register observers
	sa.attach(Logger{})
	sa.attach(Mailer{})
	sa.attach(Protector{})
	return sa
}
