package repository

import (
	"database/sql"

	"github.com/suguan/template-api/_internal"
	"github.com/suguan/template-api/db"
)

type MessageRepository struct {
	db *sql.DB
}

type MessageRepositoryI interface {
	GetMessageListByTalkerID(talkerID string, pagination _internal.PaginationQuery) []message
}

type message struct {
	MsgID string `json:"msgId"`
	// enum IsSend 1: send, 0: receive
	IsSend   uint8  `json:"isSend"`
	Content  string `json:"content"`
	TalkerID string `json:"talkerId"`
}

func NewMessageRepository(source *db.SqliteDB) MessageRepositoryI {
	return &MessageRepository{
		db: source.DB,
	}
}

func (repo *MessageRepository) GetMessageListByTalkerID(talkerID string, pagination _internal.PaginationQuery) []message {
	rows, err := repo.db.Query(`
		SELECT
			msgId,
			isSend,
			content
		FROM
			message
		WHERE
			talkerId = ?
		ORDER BY
			createTime DESC
		LIMIT
			?, ?
	`, talkerID, pagination.Page*pagination.PageSize, pagination.PageSize)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var messages []message
	for rows.Next() {
		var msg message
		err = rows.Scan(&msg.MsgID, &msg.IsSend, &msg.Content)
		if err != nil {
			panic(err)
		}
		messages = append(messages, msg)
	}

	return messages
}
