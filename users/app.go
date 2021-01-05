package users

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	status_code "github.com/tinnagorn/my-golang-service-template/statuscode"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "password"
	dbName := "example_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (s *Service) CreateUser(requestID string, req *Request) ([]byte, error) {

	db := dbConn()

	insForm, err := db.Prepare("insert into users(user_id,mobile_no,first_name,last_name,created_at) values(?,?,?,?,CURDATE())")

	if err != nil {
		result, err := json.Marshal(Response{
			Code:    status_code.CreateUserError,
			Message: "Create User fail",
		})
		return result, err
	}

	//userID := uuid.New().String()
	userID := strconv.Itoa(rand.Intn(1000000)) // Mock ID
	mobileNO := req.MobileNo
	firstName := req.FirstName
	lastName := req.LastName

	insForm.Exec(userID, mobileNO, firstName, lastName)

	defer db.Close()

	return nil, nil
}

func (s *Service) GetUser(requestID string, req *Request) *Response {

	db := dbConn()

	UserID := requestID

	selDB, err := db.Query("SELECT * FROM users WHERE user_id=?", UserID)

	if err != nil {
		result := Response{
			Code:    status_code.ValidationError,
			Message: "Inquiry Failed !",
		}
		return &result
	}

	user := User{}

	for selDB.Next() {
		var userID, mobileNo, firstName, lastName string
		err = selDB.Scan(&userID, &mobileNo, &firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}
		user.UserID = userID
		user.MobileNo = mobileNo
		user.FirstName = firstName
		user.LastName = lastName
	}

	defer db.Close()

	result := Response{
		Code:      status_code.Success,
		Message:   "Get User Success !",
		UserID:    user.UserID,
		MobileNo:  user.MobileNo,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return &result

}
