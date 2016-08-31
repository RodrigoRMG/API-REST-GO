package main

import(
"log"
"net/http"
"encoding/json"
"github.com/gorilla/mux"
"./db"
"./structures"
)


func main(){
	db.InitializeDatabase()
	defer db.CloseConnection()
	r :=mux.NewRouter()
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
	r.HandleFunc("/user/new",NewUser).Methods("POST")


	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]
	
	status :="success"
	var message string
	user:=db.Getuser(user_id)

	if(user.Usuario==""){
		status ="error"
		message = "usuario no encontrado"
	}
	response :=structures.Response{status,user,message}
	json.NewEncoder(w).Encode(response)
 }

 func NewUser(w http.ResponseWriter, r* http.Request){
 	user :=getUserRequest(r)
 	//user:=db.CreateUser(user)
 	response:=structures.Response{"success",db.CreateUser(user),""}
 	json.NewEncoder(w).Encode(response)
 }


func getUserRequest(r* http.Request) structures.User{
	var user structures.User

	decoder := json.NewDecoder(r.Body)
	err:=decoder.Decode(&user)
	if err!=nil{
		log.Fatal(err)
	}
	return user
}
