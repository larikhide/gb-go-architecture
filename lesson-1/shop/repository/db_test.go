package repository

import (
	"fmt"
	"shop/models"
	"testing"
)

var existingItem = &models.Item{
	ID:    int32(1),
	Name:  "TestName_1",
	Price: 10.0,
}

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
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	exampleItem, err := mDB.CreateItem(existingItem)
	if err != nil {
		t.Error("unexpected create error") // неожидаемая, т.к. проверяем другой метод в этом тесте, верна логика?
	}

	gottenItem, err := mDB.GetItem(exampleItem.ID)
	if gottenItem == nil {
		if err == fmt.Errorf("Item with ID: %d is not found", gottenItem.ID) {
			t.Error("expected ID error")
			return
		}
		t.Error("unexpected get error")
	}

	if gottenItem.Name != exampleItem.Name {
		t.Errorf("expected name == %s, have %s", gottenItem.Name, exampleItem.Name)
	}

	if gottenItem.Price != exampleItem.Price {
		t.Errorf("expected name == %d, have %d", gottenItem.Price, exampleItem.Price)
	}
}

func TestMapDBDeleteItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	deletedItem, err := mDB.CreateItem(existingItem)
	if err != nil {
		t.Error("unexpected create error")
	}

	err = mDB.DeleteItem(deletedItem.ID)
	if err != nil {
		t.Error("some delete error")
	}

	_, err = mDB.GetItem(deletedItem.ID)
	if err == nil {
		t.Error("item not deleted")
	}
}

func TestMapDBUpdateItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	exampleItem, err := mDB.CreateItem(existingItem)
	if err != nil {
		t.Error("unexpected create error") // неожидаемая, т.к. проверяем другой метод в этом тесте, верна логика?
	}

	updName := "UpdatedName"
	updPrice := int32(322)

	updatedItem, err := mDB.UpdateItem(exampleItem, updName, updPrice)
	if updatedItem == nil {
		if err == fmt.Errorf("Item with ID: %d is not found", exampleItem.ID) {
			t.Error("expected ID error")
			return
		}
		t.Error("unexpected update error")
	}

	if updatedItem.Name != updName {
		t.Errorf("expected name == %s, have %s", updatedItem.Name, updName)
	}

	if updatedItem.Price != updPrice {
		t.Errorf("expected name == %d, have %d", updatedItem.Price, updPrice)
	}
}
