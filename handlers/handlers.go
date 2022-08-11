package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	if year == "" {
		winners, err := data.ListAllJSON()
		//fmt.Print(winners)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(winners)
	} else {
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
		data.AddNewWinner(req.Body)
		//Inside this block, call res.WriteHeader() passing it
		//the argument http.StatusCreated to set the response
		//status code to 201.
		res.WriteHeader(http.StatusCreated)
	}
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
