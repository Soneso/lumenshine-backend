package db

import (
	"errors"
	"strings"
	"time"

	"github.com/Soneso/lumenshine-backend/admin/models"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

//UserDetails holds the user info
type UserDetails struct {
	ID        int      `json:"id"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastName"`
	Phone     string   `json:"phone"`
	Active    bool     `json:"active"`
	Groups    []string `json:"groups"`
}

//GetUserByEmail returns the user details for the specified email
func GetUserByEmail(email string) (*UserDetails, error) {
	user, err := models.AdminUsersG(
		qm.Load("UserAdminUsergroups.Group"),
		qm.Where("email=?", email)).One()

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("No user found")
	}

	var groups []string
	for _, g := range user.R.UserAdminUsergroups {
		groups = append(groups, g.R.Group.Name)
	}

	return &UserDetails{
		ID:        user.ID,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.Forename,
		LastName:  user.Lastname,
		Phone:     user.Phone,
		Active:    user.Active,
		Groups:    groups}, nil
}

//GetUserByID returns the user details for the specified ID
func GetUserByID(userID int) (*UserDetails, error) {
	user, err := models.AdminUsersG(
		qm.Load("UserAdminUsergroups.Group"),
		qm.Where("id=?", userID)).One()

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("No user found")
	}

	var groups []string
	for _, g := range user.R.UserAdminUsergroups {
		groups = append(groups, g.R.Group.Name)
	}

	return &UserDetails{
		ID:        user.ID,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.Forename,
		LastName:  user.Lastname,
		Phone:     user.Phone,
		Active:    user.Active,
		Groups:    groups}, nil
}

//UpdateLastLogin updates in the db the last login date
func UpdateLastLogin(user UserDetails) error {
	dbUser, err := models.FindAdminUserG(user.ID)
	if err != nil {
		return err
	}
	dbUser.LastLogin = time.Now()
	return dbUser.UpdateG(models.AdminUserColumns.LastLogin)
}

//RegisterUser creates a new user
func RegisterUser(user *models.AdminUser, groups []string) error {
	groupArgs := make([]interface{}, len(groups))
	for i, v := range groups {
		groupArgs[i] = strings.ToLower(v)
	}

	q := []qm.QueryMod{}
	q = append(q, qm.Select(models.AdminGroupColumns.ID, models.AdminGroupColumns.Name))
	q = append(q, qm.WhereIn("lower("+models.AdminGroupColumns.Name+") IN ?", groupArgs...))
	dbGroups, err := models.AdminGroupsG(q...).All()
	if err != nil {
		return err
	}

	if dbGroups == nil || len(dbGroups) == 0 {
		return errors.New("User must be in a valid group")
	}

	err = user.InsertG()

	if err != nil {
		return err
	}

	for _, group := range dbGroups {
		usergroup := models.AdminUsergroup{UserID: user.ID, GroupID: group.ID}
		err = usergroup.InsertG()
		if err != nil {
			return err
		}
	}

	return nil
}

//ActivateUser activates or deactivates the state of a user
func ActivateUser(userID int, active bool, updatedBy string) error {
	user, err := models.FindAdminUserG(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("User not found")
	}

	user.Active = active
	user.UpdatedAt = time.Now()
	user.UpdatedBy = updatedBy

	err = user.UpdateG()
	if err != nil {
		return err
	}

	return nil
}

//AllUsers - returns all users
func AllUsers() (models.AdminUserSlice, error) {
	users, err := models.AdminUsersG(
		qm.Load("UserAdminUsergroups.Group")).All()

	if err != nil {
		return nil, err
	}

	return users, nil
}

//UpdateUser - edits the user properties
func UpdateUser(user models.AdminUser, updatedBy string) error {
	dbUser, err := models.FindAdminUserG(user.ID)

	if err != nil {
		return err
	}

	if dbUser == nil {
		return errors.New("User not found")
	}

	dbUser.Password = user.Password
	dbUser.Forename = user.Forename
	dbUser.Lastname = user.Lastname
	dbUser.Phone = user.Phone
	dbUser.Active = user.Active
	dbUser.UpdatedBy = updatedBy
	dbUser.UpdatedAt = time.Now()

	err = dbUser.UpdateG()
	if err != nil {
		return err
	}

	return nil
}

//SetGroups - sets the groups for the specified user
func SetGroups(userID int, groups []string) error {
	dbUser, err := models.FindAdminUserG(userID)
	if err != nil {
		return err
	}
	if dbUser == nil {
		return errors.New("User not found")
	}

	groupArgs := make([]interface{}, len(groups))
	for i, v := range groups {
		groupArgs[i] = strings.ToLower(v)
	}
	q := []qm.QueryMod{}
	q = append(q, qm.Select(models.AdminGroupColumns.ID, models.AdminGroupColumns.Name))
	q = append(q, qm.WhereIn("lower("+models.AdminGroupColumns.Name+") IN ?", groupArgs...))
	dbGroups, err := models.AdminGroupsG(q...).All()
	if err != nil {
		return err
	}
	if dbGroups == nil || len(dbGroups) == 0 {
		return errors.New("Invalid groups")
	}

	usergroupsToDelete, err := models.AdminUsergroupsG(qm.Where("user_id=?", userID)).All()
	if err != nil {
		return err
	}

	if len(usergroupsToDelete) > 0 {
		err = usergroupsToDelete.DeleteAllG()
		if err != nil {
			return err
		}
	}

	for _, group := range dbGroups {
		usergroup := models.AdminUsergroup{UserID: userID, GroupID: group.ID}
		err = usergroup.InsertG()
		if err != nil {
			return err
		}
	}

	return nil
}

//AllActiveAdministrators - returns all users
func AllActiveAdministrators() (models.AdminUserSlice, error) {
	users, err := models.AdminUsersG(
		qm.Load("UserAdminUsergroups.Group"),
		qm.Where("active=?", true)).All()

	var admins models.AdminUserSlice
	for _, user := range users {
		for _, userGroup := range user.R.UserAdminUsergroups {
			if strings.EqualFold(userGroup.R.Group.Name, "administrators") {
				admins = append(admins, user)
				break
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return admins, nil
}