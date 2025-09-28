package db

import (
	"github.com/zibbadies/homies/internal/homies/db/execers"
	"github.com/zibbadies/homies/internal/homies/models"
)

func NewList(houseId string, name string) error {
	return execers.NewListEx(db, houseId, name)
}

func GetLists(houseId string) ([]models.List, error) {
	return execers.GetListsEx(db, houseId)
}

func GetListHID(listId string) (string, error) {
	return execers.GetListHIDEx(db, listId)
}

func GetItems(listId string) ([]models.Item, error) {
	return execers.GetItemsEx(db, listId)
}

func NewItem(text string, listId string, authorId string) error {
	return execers.NewItemEx(db, text, listId, authorId)
}

func UpdateItem(listId string, itemId string, text string, authorId string) error {
	return execers.UpdateItemEx(db, listId, itemId, text, authorId)
}
