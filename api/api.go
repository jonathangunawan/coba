package api

import (
	"boyzgenk/coba/user"
	"boyzgenk/coba/user/crud"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	user user.UserItf
	crud crud.UserCRUDItf
}

func NewAPI(u user.UserItf, cu crud.UserCRUDItf) API {
	return API{
		user: u,
		crud: cu,
	}
}

func (a API) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := a.user.Login(vars["email"], vars["password"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success Login!, with Key & Token stored in Redis")
}

func (a API) Logout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := a.user.Logout(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success Logout!")
}

func (a API) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := a.crud.GetUser(vars["email"], vars["password"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Data: %v", res)
}

func (a API) InsertUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	param := crud.UserRow{
		Email:    vars["email"],
		Address:  vars["address"],
		Password: vars["password"],
	}
	err := a.crud.InsertUser(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Insert Success! User Data: %v", param)
}

func (a API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	param := crud.UserRow{
		Email:    vars["email"],
		Address:  vars["address"],
		Password: vars["password"],
		UserID:   id,
	}
	err = a.crud.UpdateUser(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Update Success! User Data: %v", param)
}

func (a API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = a.crud.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Delete Success! User Data with ID: %v", id)
}
