package repository

import (
	"fmt"
	"shop/models"
	"testing"
)

func TestMapDBCreateItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 15.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 20.0,
	}
	// TEST BEGINS HERE

	mDB.maxID = currentID

	newItem := &models.Item{
		Name:  "TestName_4",
		Price: 25.0,
	}

	createdItem, err := mDB.CreateItem(newItem)
	if err != nil {
		t.Error("some expected create error")
	}
	currentID++

	if createdItem.ID != currentID {
		t.Errorf("expected id == %d, have %d", currentID, createdItem.ID)
	}
	if createdItem.Name != newItem.Name {
		t.Errorf("expected name == %s, have %s", newItem.Name, createdItem.Name)
	}
	if createdItem.Price != newItem.Price {
		t.Errorf("expected price == %d, have %d", newItem.Price, createdItem.Price)
	}

	if createdItem == nil {
		t.Error("got nil item")
	}

	// а это вообще зачем здесь было?
	/* existingItem := mDB.db[currentID]
	if existingItem == nil {
		t.Fatal("got nil item")
	}

	if existingItem.ID != currentID {
		t.Fatal("expected id == ")
	}
	if existingItem.Name != newItem.Name {
		t.Fatal("expected name == ")
	}
	if existingItem.Price != newItem.Price {
		t.Fatal("expected name == ")
	} */
}

func TestMapDBGetItem(t *testing.T) {
	// тесты в последствии запускаются сразу все? если да, то логично тестовую БД вынести в глобальую переменную с тестовыми экземплярами
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}

	currentID++

	_, err := mDB.GetItem(currentID)

	if err == fmt.Errorf("Item with ID: %d is not found", currentID) {
		t.Error("expected ID error")
		return
	}
}

/* func TestMapDBDeleteItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)

	err := mDB.DeleteItem(currentID)
	if err != nil {
		t.Error("some expected delete error")
		return
	}

} */
