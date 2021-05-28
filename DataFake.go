package main

import "strconv"

type Coffee struct {
	Id          string
	Name        string
	Amount      float64
	Image       string
	Description string
}

var data = []map[string]string {
	{"Id":"1", "Name":"Chanh đá1", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"2", "Name":"Chanh đá2", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"3", "Name":"Chanh đá3", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"4", "Name":"Chanh đá4", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"5", "Name":"Chanh đá5", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"6", "Name":"Chanh đá6", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"7", "Name":"Chanh đá7", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"8", "Name":"Chanh đá8", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"9", "Name":"Chanh đá9", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"10", "Name":"Chanh đá10", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"11", "Name":"Chanh đá11", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
	{"Id":"12", "Name":"Chanh đá12", "Amount":"12000", "Image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR2HeUbYitn9tncUeOtOgjqV7be8IJj3tc7pg&usqp=CAU", "Description":"Ngon lành"},
}

func DataFake() (l []Coffee) {
	for _, v := range data {
		amount, _ := strconv.ParseFloat(v["Amount"], 64)
		l = append(l, Coffee{
			Id: v["Id"],
			Name: v["Name"],
			Amount: amount,
			Image: v["Image"],
			Description: v["Description"],
		})
	}

	return l
}