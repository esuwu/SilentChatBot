package repository

import (
	"github.com/esuwu/SilentChatBot/models"
)

func (db *DB) GetFreeGroups() (models.Groups, error){
	groups := models.Groups{}
	rows, err := db.DBConnPool.Query("select group_id from groups where is_group_busy = false")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		group := models.Group{}
		err = rows.Scan(&group.GroupID)
		if err != nil {
			rows.Close()
			return nil, err
		}
		groups = append(groups, &group)
	}
	rows.Close()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (db *DB) IsUserExist(userID int) bool {
	rows, _ := db.DBConnPool.Exec("select id from users where id = $1", userID)
	if rows.RowsAffected() == 0 {
		return false
	}
	return true
}

func (db *DB)  IsUserBusy(userID int) bool {
	var isBusy bool
	db.DBConnPool.QueryRow("select is_user_busy from users where id = $1", userID).Scan(&isBusy)
	return isBusy
}

func (db *DB) CreateUser(user *models.User) error{
	_, err := db.DBConnPool.Exec("insert into users (id, nickname, is_user_busy, user_chat_id) values($1, $2, $3, $4)", user.UserID, user.NickName, false, user.UserChatID)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB)  GetGroupID(userID int) int {
	var groupID int
	db.DBConnPool.QueryRow("select group_id from groups where first_user_id = $1 OR second_user_id = $2", userID, userID).Scan(&groupID)

	return groupID
}

func (db *DB) IsUserInBusyChat(userID int) bool {
	var isBusy bool
	db.DBConnPool.QueryRow("select is_group_busy from groups where first_user_id = $1 OR second_user_id = $2", userID, userID).Scan(&isBusy)

	return isBusy
}

func (db *DB) GetAnotherUser(ExistingUserID int, groupID int) (*models.User, error){
	var userID int
	user := models.User{}
	db.DBConnPool.QueryRow("select first_user_id from groups where group_id = $1", groupID).
		Scan(&userID)
	if userID != ExistingUserID{
		db.DBConnPool.QueryRow("select id, nickname, user_chat_id, is_user_busy from users where id = $1", userID).
			Scan(&user.UserID, &user.NickName, &user.UserChatID, &user.IsUserBusy)
		return &user, nil
	}
	db.DBConnPool.QueryRow("select second_user_id from groups where group_id = $1", groupID).
			Scan(&userID)

	db.DBConnPool.QueryRow("select id, nickname, user_chat_id, is_user_busy from users where id = $1", userID).
		Scan(&user.UserID, &user.NickName, &user.UserChatID, &user.IsUserBusy)
	return &user, nil
}

func (db *DB) MakeUserBusy(userID int) error {
	_, err := db.DBConnPool.Exec("UPDATE users SET is_user_busy = true WHERE id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) AddNextUserToGroup(userID int, group *models.Group) error {
	_, err := db.DBConnPool.Exec("update groups set second_user_id = $1, is_group_busy = true WHERE group_id = $2", userID, group.GroupID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateGroupAndAddUser(userID int) (*models.Group, error) {
	group := models.Group{}
	db.DBConnPool.QueryRow("INSERT INTO groups (is_group_busy, first_user_id) values(false, $1) RETURNING groups.group_id, is_group_busy", userID).
		Scan(&group.GroupID, &group.IsGroupBusy)

	return &group, nil
}

func (db *DB) LeaveChat(userID int) error {
	var firstID int
	var secondID int
	db.DBConnPool.QueryRow("delete from groups where first_user_id = $1 OR second_user_id = $2 RETURNING first_user_id, second_user_id", userID, userID).
		Scan(&firstID, &secondID)
	_, err := db.DBConnPool.Exec("update users set is_user_busy = false where id = $1 OR id = $2;", firstID, secondID)
	if err != nil {
		return err
	}
	return nil
}