package models

import (
	"encoding/json"
	"github.com/divoc/portal-api/swagger_gen/models"
	"github.com/pkg/errors"
	"time"
)

//mobile,enrollmentScopeId,nationalId,dob,gender,name,email
type Enrollment struct {
	Phone             string          `json:"phone"`
	EnrollmentScopeId string          `json:"enrollmentScopeId"`
	NationalId        string          `json:"nationalId"`
	Dob               string          `json:"dob"`
	Gender            string          `json:"gender"`
	Name              string          `json:"name"`
	Email             string          `json:"email"`
	Code              string          `json:"code"`
	Certified         bool            `json:"certified"`
	ProgramId         string          `json:"programId"`
	Address           *models.Address `json:"address"`
}

func (e Enrollment) MarshalJSON() ([]byte, error) {
	_, err := time.Parse("2006-01-02", e.Dob)
	if err != nil {
		return nil, errors.New("Detail: DoB is not in required format (Ex: 2006-01-21)")
	}
	attributes := map[string]interface{}{
		"phone":             e.Phone,
		"enrollmentScopeId": e.EnrollmentScopeId,
		"nationalId":        e.NationalId,
		"dob":               e.Dob,
		"gender":            e.Gender,
		"name":              e.Name,
		"email":             e.Email,
		"code":              e.Code,
		"certified":         e.Certified,
		"programId":         e.ProgramId,
		"address":           e.Address,
	}
	bytes, err := json.Marshal(attributes)
	return bytes, err
}
