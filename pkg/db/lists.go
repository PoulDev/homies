package db

import (
	"github.com/PoulDev/roommates-api/pkg/logger"

	"strconv"
	"log"
	"fmt"
)

type List struct {
	Name string		`json:"name"`
	Id string		`json:"id"`
}

type Item struct {
	Text string 	`json:"text"`
	Completed bool	`json:"completed"`
	Author string	`json:"author"`
}

func NewListEx(exec Execer, userId string, name string) error {
	b_id, err := UUIDString2Bytes(userId);
	if (err != nil) {return err;}

	_, err = exec.Exec(`
		INSERT INTO lists (user_id, name)
		VALUES (?, ?)`,
		b_id, name,
	)
	if err != nil {return err;}

	return nil;
}

func NewList(userId string, name string) error {
	return NewListEx(db, userId, name)
}

func GetListsEx(exec Execer, userId string) ([]List, error) {
	b_id, err := UUIDString2Bytes(userId)
	if (err != nil) { return nil, err; }

	rows, err := exec.Query(`SELECT id, name FROM lists WHERE user_id = ?`, b_id);
	defer rows.Close()

	if (err != nil) {
		logger.Logger.Error("list get error", "err", err.Error())
		return nil, fmt.Errorf("Internal error, please try again later")
	}

	var lists []List;
	for rows.Next() {
		var list List;
		var id uint;

		if err := rows.Scan(&id, &list.Name); err != nil {
            logger.Logger.Error("list get error", "err", err.Error())
			return nil, fmt.Errorf("Internal error, please try again later")
        }
		list.Id = strconv.FormatUint(uint64(id), 10);
		lists = append(lists, list)
	}

    if err := rows.Err(); err != nil {
		logger.Logger.Error("list get error", "err", err.Error())
		return nil, fmt.Errorf("Internal error, please try again later")
    }

	return lists, nil;
}

func GetLists(userId string) ([]List, error) {
	return GetListsEx(db, userId)
}

func GetItemsEx(exec Execer, listId string) ([]Item, error) {
	b_id, err := strconv.Atoi(listId)
	if (err != nil) { 
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "listId", listId)
		return nil, fmt.Errorf("There's a problem with this list, please try again later")
	}

	rows, err := exec.Query(`SELECT text, completed, author FROM todos WHERE list_id = ?`, b_id);
	defer rows.Close()

	if err != nil {
		logger.Logger.Error("list DB select error", "err", err.Error(), "listId", listId)
		return nil, fmt.Errorf("Internal error, please try again later")
	}

	var items []Item;
	for rows.Next() {
		var item Item;
		var author []byte;

		if err := rows.Scan(&item.Text, &item.Completed, &author); err != nil {
			logger.Logger.Error("list row scan error", "err", err.Error(), "listId", listId)
			return nil, fmt.Errorf("There's a problem with your list, please try again later")
		}
		log.Println(item.Text)
		item.Author, err = UUIDBytes2String(author);

		if (err != nil) {
			logger.Logger.Error("list UUIDBytes2String error", "err", err.Error(), "listId", listId)
			return nil, fmt.Errorf("There's a problem with your list, please try again later")
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		logger.Logger.Error("list rows error", "err", err.Error(), "listId", listId)
		return nil, fmt.Errorf("There's a problem with your list, please try again later")
	}

	return items, nil;
}

func GetItems(listId string) ([]Item, error) {
	return GetItemsEx(db, listId)
}


func NewItemEx(exec Execer, text string, listId string, authorId string) error {
	l_id, err := strconv.Atoi(listId)
	if err != nil {
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with your list, please try again later")
	}

	a_id, err := UUIDString2Bytes(authorId)
	if err != nil {
		logger.Logger.Error("list UUIDString2Bytes error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with your list, please try again later")
	}

	_, err = exec.Exec(`UPDATE lists SET items = items + 1 WHERE id = ?`, l_id)
	if err != nil {
		logger.Logger.Error("list update error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with updating your list, please try again later")
	}

	_, err = exec.Exec(`
		INSERT INTO todos (text, list_id, author)
		VALUES (?, ?, ?)`, text, l_id, a_id)
	if err != nil {
		logger.Logger.Error("list insert error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with updating your list, please try again later")
	}

	return nil
}

func NewItem(text string, listId string, authorId string) error {
	return NewItemEx(db, text, listId, authorId)
}
