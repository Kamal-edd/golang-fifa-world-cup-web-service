package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
	"strconv"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	if year == "banana" { //this just a joke
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if year == "" {
		winners, err := data.ListAllJSON()
		//fmt.Print(winners)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(winners)
	} else {
		year_nbr, err := strconv.Atoi(year)
		if err != nil && (year_nbr > 1930 || year_nbr < 2022) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		filteredWinners, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(filteredWinners)

	}

}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	//Find the AddNewWinner function. Inside this function,
	//call req.Header.Get() passing it the argument "X-ACCESS-TOKEN"
	//to read from this request header. Assign the result to
	//a new variable called accessToken.
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	//Next, call the data.IsAccessTokenValid() function,
	//passing it accessToken as argument and store the result
	//in a new variable called isTokenValid.
	isTokenValid := data.IsAccessTokenValid(accessToken)
	//Using this variable, create an if statement which uses
	//!isTokenValid (notice the negation character !) as the condition.
	if !isTokenValid {
		//On the if block, call res.WriteHeader() passing it
		//the argument http.StatusUnauthorized to set the response
		//status code to 401.
		res.WriteHeader(http.StatusUnauthorized)
	} else {
		//Create an else block for this same conditional.
		//Now let's read from the payload of the POST request
		//and add a new winner. Inside the else block, before
		//writing the status code, call the data.AddNewWinner()
		//method, passing it req.Body as argument
		var winners data.Winners
		if len(winners.Winners) < 21 {
			err := data.AddNewWinner(req.Body)
			//Assign the return value of calling the data.AddNewWinner()
			//method to a new variable called err. Then, create an
			//if statement which checks if err != nil.
			if err != nil { //On the if block of this statement,
				//call res.WriteHeader() passing it the argument
				//http.StatusUnprocessableEntity to respond
				//with a status code of 422,
				res.WriteHeader(http.StatusUnprocessableEntity)
				//and on the very next line
				return
				//add an empty return statement.
			}
			//Inside this block, call res.WriteHeader() passing it
			//the argument http.StatusCreated to set the response
			//status code to 201.
			res.WriteHeader(http.StatusCreated)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {
	//Find the WinnersHandler function. Inside this function,
	//create a switch/case statement that switches on req.Method.
	switch req.Method {
	//Add a single case for http.MethodGet.
	case http.MethodGet:
		//This case should call the function ListWinners passing
		//it the arguments res and req in this order.
		ListWinners(res, req)
	//Add a second case statement. This time, for http.MethodPost.
	case http.MethodPost:
		//This case should call the function AddNewWinner passing
		//it the arguments res and req in this order.
		AddNewWinner(res, req)
	//Lastly, add a default case to switch/case
	default:
		//which calls res.WriteHeader passing it the argument
		//http.StatusMethodNotAllowed. This will create a response with the status code of 405.
		res.WriteHeader(http.StatusMethodNotAllowed)
	}

}
