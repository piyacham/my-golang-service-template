package users

import (
	"encoding/json"

	user_mgr "github.com/tinnagorn/my-golang-service-template/users"
	status_code "github.com/tinnagorn/my-golang-service-template/statuscode"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) getUser(requestID string, req *Request) ([]byte, error) {

	user, err := user_mgr.GetUserByID(req.UserID)

	if err != nill {
		result := Response{
			Code:    status_code.ValidationError,
			Message: "Inquiry Failed !"
		}
		return nil, &result
	}

	result := Response{
		Code:    status_code.Success,
		Message: "Get User Success !",
		UserID   : user.user_id,
		MobileNo : user.mobile_no,
		FirstName : user.first_name
		LastName  : user.last_name,
		CreatedAt : user.createdAt,
		UpdatedAt : user.updatedAt,
		DeletedAt : user.deletedAt
	}

	return result, err
}


func (s *Service) createUser(requestID string, req *Request) ([]byte, error) {

	user, err := user_mgr.CreateUser(req)

	if err != nill {
		result := Response{
			Code:    status_code.CreateUserError,
			Message: "Create User Failed !"
		}
		return nil, &result
	}

	result := Response{
		Code:    status_code.Success,
		Message: "Create User Success !",
		UserID   : user.user_id,
		MobileNo : user.mobile_no,
		FirstName : user.first_name
		LastName  : user.last_name,
		CreatedAt : user.createdAt,
		UpdatedAt : user.updatedAt,
		DeletedAt : user.deletedAt
	}

	return result, err

}

func (s *Service) updateUser(requestID string, req *Request) ([]byte, error) {

	user, err := user_mgr.updateUser(req)

	if err != nill {
		result := Response{
			Code:    status_code.ErrorDefault,
			Message: "Create User Failed !"
		}
		return nil, &result
	}

	user, err = user_mgr.GetListByUserID(req.UserID)

	result := Response{
		Code:    status_code.Success,
		Message: "Update User Success !",
		UserID   : user.user_id,
		MobileNo : user.mobile_no,
		FirstName : user.first_name
		LastName  : user.last_name,
		CreatedAt : user.createdAt,
		UpdatedAt : user.updatedAt,
		DeletedAt : user.deletedAt
	}

	return result, err
}
func (s *Service) deleteUser(requestID string, req *Request) ([]byte, error) {
	
	user, err := user_mgr.deleteUser(req)

	if err != nil { 
		result := Response{
			Code:    status_code.DeleteUserError,
			Message: "Delete User Failed !"
			UserID   : user.user_id,
			FirstName : user.first_name
			LastName  : user.last_name
		}
		return nil, &result
	}

	result, err := json.Marshal(Response{
		Code:    status_code.Success,
		Message: "Delete User Success !",
	})
	return result, err
}
