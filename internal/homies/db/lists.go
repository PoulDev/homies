package db

import (
	"github.com/PoulDev/homies/internal/homies/logger"
	"github.com/PoulDev/homies/internal/homies/models"

	"strconv"
	"fmt"
)

func NewListEx(exec Execer, houseId string, name string) error {
	houseIdInt, err := strconv.Atoi(houseId)
	if (err != nil) {
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "houseId", houseId)
		return fmt.Errorf("There's a problem with your house, please try again later")
	}

	_, err = exec.Exec(`
		INSERT INTO lists (house_id, name)
		VALUES ($1, $2)`,
		houseIdInt, name,
	)
	if err != nil {return err;}

	return nil;
}

func NewList(houseId string, name string) error {
	return NewListEx(db, houseId, name)
}

func GetListsEx(exec Execer, houseId string) ([]models.List, error) {
	houseIdInt, err := strconv.Atoi(houseId)
	if (err != nil) {
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "houseId", houseId)
		return nil, fmt.Errorf("There's a problem with your house, please try again later")
	}

	rows, err := exec.Query(`SELECT id, name FROM lists WHERE house_id = $1`, houseIdInt);
	defer rows.Close()

	if (err != nil) {
		logger.Logger.Error("list get error", "err", err.Error())
		return nil, fmt.Errorf("Internal error, please try again later")
	}

	var lists []models.List;
	for rows.Next() {
		var list models.List;
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

func GetLists(houseId string) ([]models.List, error) {
	return GetListsEx(db, houseId)
}

func GetListHID(listId string) (string, error) {
	var houseId int64;
	
	err := db.QueryRow(`SELECT house_id FROM lists WHERE id = $1`, listId).Scan(&houseId);
	if (err != nil) {
		logger.Logger.Error("user house ID retrival error", "err", err.Error())
		return "", fmt.Errorf("There's a problem with this list, please try again later")
	}

	return strconv.FormatInt(houseId, 10), nil;
}

func GetItemsEx(exec Execer, listId string) ([]models.Item, error) {
	b_id, err := strconv.Atoi(listId)
	if (err != nil) { 
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "listId", listId)
		return nil, fmt.Errorf("There's a problem with this list, please try again later")
	}

	rows, err := exec.Query(`SELECT id, text, completed, author FROM todos WHERE list_id = $1`, b_id);
	defer rows.Close()

	if err != nil {
		logger.Logger.Error("list DB select error", "err", err.Error(), "listId", listId)
		return nil, fmt.Errorf("Internal error, please try again later")
	}

	items := make([]models.Item, 0);
	for rows.Next() {
		var item models.Item;
		var iid int64;

		if err := rows.Scan(&iid, &item.Text, &item.Completed, &item.Author); err != nil {
			logger.Logger.Error("list row scan error", "err", err.Error(), "listId", listId)
			return nil, fmt.Errorf("There's a problem with your list, please try again later")
		}
		item.Id = strconv.FormatInt(iid, 10);

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

func GetItems(listId string) ([]models.Item, error) {
	return GetItemsEx(db, listId)
}


func NewItemEx(exec Execer, text string, listId string, authorId string) error {
	l_id, err := strconv.Atoi(listId)
	if err != nil {
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with your list, please try again later")
	}

	_, err = exec.Exec(`UPDATE lists SET items = items + 1 WHERE id = $1`, l_id)
	if err != nil {
		logger.Logger.Error("list update error", "err", err.Error(), "authorId", authorId)
		return fmt.Errorf("There's a problem with updating your list, please try again later")
	}

	_, err = exec.Exec(`
		INSERT INTO todos (text, list_id, author)
		VALUES ($1, $2, $3)`, text, l_id, authorId)
	if err != nil {
		logger.Logger.Error("list insert error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with updating your list, please try again later")
	}

	return nil
}

func NewItem(text string, listId string, authorId string) error {
	return NewItemEx(db, text, listId, authorId)
}

func UpdateItemEx(exec Execer, listId string, itemId string, text string, authorId string) error {
	i_id, err := strconv.Atoi(itemId)
	if err != nil {
		logger.Logger.Error("list item ID atoi error", "err", err.Error(), "listId", itemId)
		return fmt.Errorf("There's a problem with your list, please try again later")
	}

	l_id, err := strconv.Atoi(listId)
	if err != nil {
		logger.Logger.Error("list ID atoi error", "err", err.Error(), "listId", listId)
		return fmt.Errorf("There's a problem with your list, please try again later")
	}

	_, err = exec.Exec(`UPDATE todos SET text = $1, author = $2 WHERE (id = $3 AND list_id = $4)`, text, authorId, i_id, l_id)
	if err != nil {
		logger.Logger.Error("list item update error", "err", err.Error(), "itemId", itemId)
		return fmt.Errorf("There's a problem with updating your list, please try again later")
	}

	return nil
}

func UpdateItem(listId string, itemId string, text string, authorId string) error {
	return UpdateItemEx(db, listId, itemId, text, authorId)
}
