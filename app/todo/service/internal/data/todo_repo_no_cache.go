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

func (t TodoRepoNoCache) InsertItem(ctx context.Context, id uint64, content string) (uint64, error) {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return 0, err
	}

	item := &Item{
		Content: content,
	}

	err = t.db.Model(&data).Association("Items").Append(item)
	return item.Id, err
}

func (t TodoRepoNoCache) DeleteItem(ctx context.Context, id uint64, itemId uint64) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var item []Item
	if err = t.db.Model(&data).Association("Items").Find(&item, "id = ?", itemId); err != nil {
		return err
	}

	if len(item) == 0 {
		return nil
	}

	err = t.db.Model(&data).Association("Items").Delete(item[0])
	return err
}

func (t TodoRepoNoCache) FinishItem(ctx context.Context, id uint64, itemId uint64, isFinished bool, finishedUserId uint64) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var item []Item
	if err = t.db.Model(&data).Association("Items").Find(&item, "id = ?", itemId); err != nil {
		return err
	}

	if len(item) == 0 {
		return nil
	}

	item[0].IsFinished = isFinished
	item[0].FinishedUserId = finishedUserId

	res := t.db.Save(item[0])
	return res.Error

}

func (t TodoRepoNoCache) UpdateContent(ctx context.Context, id uint64, itemId uint64, content string) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var item []Item
	if err = t.db.Model(&data).Association("Items").Find(&item, "id = ?", itemId); err != nil {
		return err
	}

	if len(item) == 0 {
		return nil
	}

	item[0].Content = content

	res := t.db.Save(item[0])
	return res.Error
}

func (t TodoRepoNoCache) RestartItem(ctx context.Context, id uint64, itemId uint64) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var item []Item
	if err = t.db.Model(&data).Association("Items").Find(&item, "id = ?", itemId); err != nil {
		return err
	}

	if len(item) == 0 {
		return nil
	}

	item[0].IsFinished = false
	item[0].FinishedUserId = 0

	res := t.db.Save(item[0])
	return res.Error
}

func NewTodoRepoNoCache(db *gorm.DB) TodoRepo {
	return TodoRepoNoCache{db: db}
}
