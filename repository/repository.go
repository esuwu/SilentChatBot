package repository

import (
	"github.com/esuwu/SilentChatBot/models"
	"github.com/jackc/pgx"
)

type Repository interface {
	GetFreeGroups() (models.Groups, error)
	GetAnotherUser(ExistingUserID int, groupID int) (*models.User, error)
	MakeUserBusy(userID int) error
	AddNextUserToGroup(userID int, group *models.Group) error
	CreateGroupAndAddUser(userID int) (*models.Group, error)
	IsUserExist(userID int) bool
	CreateUser(user *models.User) error
	IsUserBusy(userID int) bool
	GetGroupID(userID int) int
	IsUserInBusyChat(userID int) bool
	LeaveChat(userID int) error
}

type DB struct {
	DBConnPool *pgx.ConnPool
}

func NewDBStore(db *pgx.ConnPool) Repository {
	return &DB{
		DBConnPool: db,
	}
}