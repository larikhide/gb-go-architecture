package repository

import (
	"errors"
	"shop/models"
)

//плюс если нужно проверить что вернулась/не вернулась конкретная ошибка
//то эту ошибку надо в пакете тогда объявить глобальной переменой и возвращать ее, чтобы потом можно было пришедший безликий интерфейс error с ней сравнить
var ErrNotFound = errors.New("not found")

type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItem(ID int32) (*models.Item, error)
	DeleteItem(ID int32) error
	UpdateItem(item *models.Item, newName string, newPrice int32) (*models.Item, error)
}

type mapDB struct {
	db    map[int32]*models.Item
	maxID int32
}

func NewMapDB() Repository {
	return &mapDB{
		db:    make(map[int32]*models.Item),
		maxID: 0,
	}
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	m.maxID++

	newItem := &models.Item{
		ID:    m.maxID,
		Price: item.Price,
		Name:  item.Name,
	}

	m.db[newItem.ID] = newItem

	return &models.Item{
		ID:    newItem.ID,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	item, ok := m.db[ID]
	if !ok {
		return nil, ErrNotFound
	}

	return &models.Item{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	delete(m.db, ID)
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item, newName string, newPrice int32) (*models.Item, error) {
	updateItem, ok := m.db[item.ID]
	if !ok {
		return nil, ErrNotFound
	}
	updateItem.Name = newName
	updateItem.Price = newPrice

	return &models.Item{
		ID:    updateItem.ID,
		Name:  updateItem.Name,
		Price: updateItem.Price,
	}, nil
}
