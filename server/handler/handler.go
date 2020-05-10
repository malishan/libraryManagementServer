package handler

import (
	"encoding/json"
	"net/http"

	"project/libraryManagementServer/db"
	mongoOp "project/libraryManagementServer/db/mongo"
	"project/libraryManagementServer/server/common"
	"project/libraryManagementServer/util"
)

// Register : Signup for new user
func Register(w http.ResponseWriter, r *http.Request) {
	var newUser db.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		util.Logf("invalid json payload received : %s", err.Error())
		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Invalid json payload received"}
		returnResponse(w, http.StatusBadRequest, msg)
		return
	}

	if newUser.FName == "" || newUser.DOB == "" || newUser.Address == "" || newUser.ContactNo == "" || newUser.Email == "" || newUser.Password == "" {
		util.Logf("Empty fields found")
		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Empty fields found"}
		returnResponse(w, http.StatusBadRequest, msg)
		return
	}

	if isPresent, err := mongoOp.UserExist(newUser); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	} else if isPresent {
		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: User Already Exists"}
		returnResponse(w, http.StatusForbidden, msg)
		return
	}

	if err := mongoOp.RegisterUser(newUser); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	msg := common.MetaStatus{Code: common.Success, Msg: "Registration Successful"}
	returnResponse(w, http.StatusBadRequest, msg)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

// // Home : the homepage of the server
// func Home(w http.ResponseWriter, r *http.Request) {

// 	msg := struct {
// 		Val string
// 	}{
// 		"Welcome to MyShelf Libaray Management Service",
// 	}

// 	err := json.NewEncoder(w).Encode(&msg)
// 	if err != nil {
// 		util.Logf("Error in encoding /home response: %s", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	util.Log(r.Host, " : Homepage Hit")
// }

// // AddBook : inserts a book into the db
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	var reqBook db.NewBook

// 	if err := json.NewDecoder(r.Body).Decode(&reqBook); err != nil {
// 		util.Logf("invalid json payload received : %s", err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Invalid json payload received"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	if len(reqBook.Author) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Author Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Author Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if len(reqBook.BookName) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Book Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if len(reqBook.Category) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Category Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Category Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if reqBook.SerialNos == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Serial Number not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Serial Number not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	_, err := reqBook.FindOneBook()
// 	if err != nil {
// 		switch err.(type) {
// 		case *db.ErrNotFound:
// 			break
// 		default:
// 			util.Log(reqBook.SerialNos, reqBook.BookName, "Error : "+err.Error())
// 			msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 			returnResponse(w, http.StatusInternalServerError, msg)
// 			return
// 		}
// 	} else if err == nil {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Duplicate Book Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Duplicate Book Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	rslt, err := reqBook.InsertBook()
// 	if err != nil {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Insertion Failed : "+err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Not Inserted"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	objectID := rslt.(primitive.ObjectID)

// 	util.Log(r.Host, " : Insertion Done for Serail No:", reqBook.SerialNos)
// 	returnResponse(w, http.StatusOK, struct {
// 		Code     int                `json:"statusCode"`
// 		ObjectID primitive.ObjectID `json:"ObjectID"`
// 	}{common.Success, objectID})
// }

// // GetBookBySerialNo : returns only one book with the particular serial ID
// func GetBookBySerialNo(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	serailNo := vars["id"]
// 	bName := vars["name"]

// 	if len(serailNo) == 0 || len(bName) == 0 {
// 		util.Log(serailNo, bName, "Error: SerailNo OR BookName Missing")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: SerailNo OR BookName Missing"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	id, err := strconv.Atoi(serailNo)
// 	if err != nil {
// 		util.Log(serailNo, bName, "Error : "+err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	respBook := &db.NewBook{
// 		SerialNos: id,
// 		BookName:  bName,
// 	}

// 	respBook, err = respBook.FindOneBook()
// 	if err != nil {
// 		switch err.(type) {
// 		case *db.ErrNotFound:
// 			util.Log(serailNo, bName, "Book Not Found : "+err.Error())
// 			msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Not Found"}
// 			returnResponse(w, http.StatusNotFound, msg)
// 			return
// 		default:
// 			util.Log(serailNo, bName, "Error : "+err.Error())
// 			msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 			returnResponse(w, http.StatusInternalServerError, msg)
// 			return
// 		}
// 	}

// 	util.Log(r.Host, " : Book Fetched for Serial Number:", id)
// 	returnResponse(w, http.StatusOK, respBook)
// }

// // GetBooksByName : returns all the books with a particular name
// func GetBooksByName(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	bName := vars["name"]
// 	if len(bName) == 0 {
// 		util.Log(bName, "Error : BookName Missing")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: BookName Missing"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	allBooks, err := db.FindAllBooks(bName)
// 	if err != nil {
// 		util.Log(bName, "Error : "+err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	if allBooks == nil {
// 		util.Log(bName, "Error : Book Not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: '" + bName + "' Not Found"}
// 		returnResponse(w, http.StatusNotFound, msg)
// 		return
// 	}

// 	util.Log(r.Host, " : Fetched '"+bName+"' books")
// 	returnResponse(w, http.StatusOK, allBooks)
// }

// // GetAllBooks : returns all books from db
// func GetAllBooks(w http.ResponseWriter, r *http.Request) {

// 	allBooks, err := db.FindAllBooks("")
// 	if err != nil {
// 		util.Log("Erro : " + err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	if allBooks == nil {
// 		util.Log("Error : No Books Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed:  No Books Found"}
// 		returnResponse(w, http.StatusNotFound, msg)
// 		return
// 	}

// 	util.Log(r.Host, " : Fetched All Books")
// 	returnResponse(w, http.StatusOK, allBooks)
// }

// // UpdateBook : updates the info of given book
// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var reqBook db.NewBook

// 	if err := json.NewDecoder(r.Body).Decode(&reqBook); err != nil {
// 		util.Log("invalid json payload received :", err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Invalid json payload received"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	if len(reqBook.Author) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Author Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Author Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if len(reqBook.BookName) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Book Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if len(reqBook.Category) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Category Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Category Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if reqBook.SerialNos == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Serial Number not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Serial Number not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	update, err := reqBook.UpdateBook()
// 	if err != nil {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Error : "+err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	updatedRslt := update.(*mongo.UpdateResult)

// 	if updatedRslt.MatchedCount <= 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Error : Book Not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Not Found"}
// 		returnResponse(w, http.StatusNotFound, msg)
// 		return
// 	} else if updatedRslt.ModifiedCount <= 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Error : Book Not Updated")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Not Updated"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	util.Log(r.Host, " : Updated SerialNo - ", reqBook.SerialNos)
// 	returnResponse(w, http.StatusOK, common.MetaStatus{Code: common.Success, Msg: reqBook.BookName + " updated successfully"})
// }

// // DeleteBook : deletes the given book from db
// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	var reqBook db.NewBook

// 	if err := json.NewDecoder(r.Body).Decode(&reqBook); err != nil {
// 		util.Log("invalid json payload received :", err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Invalid json payload received"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	if len(reqBook.BookName) == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Book Name not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Name not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	} else if reqBook.SerialNos == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Request API: Serial Number not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Serial Number not Found"}
// 		returnResponse(w, http.StatusBadRequest, msg)
// 		return
// 	}

// 	deleteCount, err := reqBook.DeleteBook()
// 	if err != nil {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Error : "+err.Error())
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Internal Server Error"}
// 		returnResponse(w, http.StatusInternalServerError, msg)
// 		return
// 	}

// 	if deleteCount == 0 {
// 		util.Log(reqBook.SerialNos, reqBook.BookName, "Error : Book Not Found")
// 		msg := common.MetaStatus{Code: common.Fail, Msg: "Failed: Book Not Found"}
// 		returnResponse(w, http.StatusNotFound, msg)
// 		return
// 	}

// 	util.Log(r.Host, " : Deleted SerialNo - ", reqBook.SerialNos)
// 	returnResponse(w, http.StatusOK, common.MetaStatus{Code: common.Success, Msg: reqBook.BookName + " deleted successfully"})
// }

// returnResponse : forms the general response for API calls
func returnResponse(w http.ResponseWriter, statusCode int, status interface{}) {
	respb, _ := json.Marshal(status)

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respb)
}
