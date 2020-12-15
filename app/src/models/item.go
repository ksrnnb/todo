package models

// Item model
type Item struct {
	ID     int    `gorm:"primaryKey"`
	TodoID int    `gorm:"not null"`
	Todo   Todo
	Name   string `gorm:"not null"`
	Done   bool   `gorm:"default:false;not null"`
}

func (item Item) Create() {
	Db.Create(&item)
}

func (item Item) Save() {
	Db.Save(&item)
}

func (item Item) Delete() {
	Db.Delete(&item)
}

func CreateNewItem(data map[string]string) {
	todo := FindTodo(data["uuid"])

	item := Item{Todo: todo, Name: data["name"]}
	item.Create()
}

func UpdateItemName(id int, uuid string, name string) {
	todo := FindTodoWithItems(uuid)
	item, found := todo.GetItem(id)

	if found {
		item.Name = name
		item.Save()
	}
}

func UpdateItemDone(id int, uuid string) {
	todo := FindTodoWithItems(uuid)
	item, found := todo.GetItem(id)

	if found {
		item.Done = !item.Done
		item.Save()
	}
}