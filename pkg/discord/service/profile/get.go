package profile

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter/fortress"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type EmployeeSearch struct {
	DiscordID string
	Email     string
	Key       string
}

func (e *Profile) GetEmployeeList(in EmployeeSearch) ([]model.Employee, error) {
	rs, err := e.adapter.Fortress().GetEmployees(fortress.EmployeeSearch{
		DiscordID: in.DiscordID,
		Email:     in.Email,
		Key:       in.Key,
	})
	if err != nil {
		return nil, err
	}

	return rs, nil
}
