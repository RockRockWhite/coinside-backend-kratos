package data

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoRepoNoCache struct {
	db *gorm.DB
}

func (t TodoRepoNoCache) Insert(ctx context.Context, data *Todo) (uint64, error) {
	res := t.db.Create(data)
	return data.Id, res.Error
}

func (t TodoRepoNoCache) FindOne(ctx context.Context, id uint64) (*Todo, error) {
	var todo Todo
	res := t.db.Model(&todo).Preload("Items").Where("id = ?", id).First(&todo)

	return &todo, res.Error
}

func (t TodoRepoNoCache) Update(ctx context.Context, newData *Todo) error {
	res := t.db.Omit(clause.Associations).Save(newData)
	return res.Error
}

func (t TodoRepoNoCache) Delete(ctx context.Context, id uint64) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	res := t.db.Select(clause.Associations).Delete(&data)
	return res.Error
}

func (t TodoRepoNoCache) InsertItem(ctx context.Context, id uint64, content string, isFinished bool, finishedUsedId uint64) (uint64, error) {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return -1,err
	}
	//
	res := t.db.Model(&data).Association("Items").Append(
		&Item{
			Content:        content,
			IsFinished:     isFinished,
			FinishedUsedId: finishedUsedId,
		})
	return ,err
}

func (t TodoRepoNoCache) DeleteItem(ctx context.Context, item_id uint64) error {
	//data, err := t.FindOne(ctx, id)
	//if err != nil {
	//	return err
	//}
	//
	//var todos []Todo
	//if err = t.db.Model(&data).Association("Todos").Find(&todos, "content = ?", content); err != nil {
	//	return err
	//}
	//
	//if len(todos) == 0 {
	//	return nil
	//}
	//
	//err = t.db.Model(&data).Association("Todos").Delete(todos[0])
	//return err
}
func (t TodoRepoNoCache) UpdateItem(ctx context.Context, data *Item) error {

}

func NewTodoRepoNoCache(db *gorm.DB) TodoRepo {
	return TodoRepoNoCache{db: db}
}
