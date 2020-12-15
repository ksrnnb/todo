package models

// Todo model
type Todo struct {
	ID    int    `gorm:"primaryKey"`
	UUID  string `gorm:"type:varchar(36);unique;not null"`
	Items []Item
}

func (todo Todo) Create() {
	Db.Create(&todo)
}

func (todo Todo) Save() {
	Db.Save(&todo)
}

func (todo Todo) Delete() {
	Db.Delete(&todo)
}

// HasItem check whether Todo struct has item or not by item id
func (todo Todo) GetItem(id int) (item Item, found bool) {
	found = false
	for _, itemInTodo := range todo.Items {
		if itemInTodo.ID == id {
			item = itemInTodo
			found = true
			break
		}
	}

	return item, found
}

// FindTodo finds todo without items
func FindTodo(uuid string) (todo Todo) {
	Db.Where("uuid=?", uuid).First(&todo)
	return todo
}

// FindTodoWithItems finds todo with related items
func FindTodoWithItems(uuid string) (todo Todo) {
	Db.Where("uuid=?", uuid).Preload("Items").Find(&todo)
	return todo
}
