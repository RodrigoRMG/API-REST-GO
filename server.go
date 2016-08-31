package main

import(
"log"
"net/http"
"encoding/json"
"github.com/gorilla/mux"
"./db"
)




func main(){
	db.InitializeDatabase()
	defer db.CloseConnection()
	r :=mux.NewRouter()
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
	log.Println("El servidor se encuentra en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	user_id := vars["id"]
	user:=db.Getuser(user_id)
	json.NewEncoder(w).Encode(user)
 }

