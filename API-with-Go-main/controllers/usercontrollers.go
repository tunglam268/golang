package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllPerson get all user data
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	database.Connector.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//GetPersonByID returns user with specific ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ID"]

	var user entity.User
	database.Connector.First(&user, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//CreatePerson creates user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)

	database.Connector.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

//UpdatePersonByID updates user with respective ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)
	database.Connector.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//DeletePersonByID delete's user with specific ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user entity.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}

func UserWithdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user entity.User
	database.Connector.First(&user, key)
	fmt.Print(user.Balance)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Unreadable !!!")
	}

	var cb entity.ChangeBalance

	json.Unmarshal(requestBody, &cb)
	Withdraw(&user, cb.Amount)

	// user.modified_time = fmt.Sprintf(time.Now())
	database.Connector.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
func UserDeposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user entity.User
	database.Connector.First(&user, key)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Unreadble ")

	}
	var cb entity.ChangeBalance
	err1 := json.Unmarshal(requestBody, &cb)
	if err1 != nil {
		fmt.Print("error")
	}
	Deposit(&user, cb.Amount)
	database.Connector.Save(&user)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
func Withdraw(user *entity.User, num float64) {
	if user.Balance < 5000 {
		fmt.Println("You dont have enough money !")
	}
	if user.Balance < num {
		fmt.Println("the amount you need to withdraw must be less than your balance !")
	}
	user.Balance = user.Balance - num
}

func Deposit(user *entity.User, num float64) {
	fmt.Printf("Deposit %v", num)
	user.Balance = user.Balance + num

}
