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

func (t TodoRepoNoCache) SetItem(ctx context.Context, id uint64, content string, isFinished bool, finishedUsedId uint64) (uint64, error) {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	//
	var items []Item
	if err = t.db.Model(&data).Association("Items").Find(&items, "todo_id = ?", userId); err != nil {
		return err
	}

	// add a item
	if len(items) == 0 {
		err = t.db.Model(&data).Association("Items").Append(
			&Item{
				Content:        content,
				IsFinished:     isFinished,
				FinishedUsedId: finishedUsedId,
			})
		return err
	}

	// update a item
	items[0].Content = content
	items[0].IsFinished = isFinished
	items[0].FinishedUsedId = finishedUsedId
	res := t.db.Save(items[0])
	return res.Error
}

func (t TodoRepoNoCache) DeleteItem(ctx context.Context, id uint64) error {
	data, err := t.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var members []TeamMember
	if err = u.db.Model(&data).Association("Members").Find(&members, "user_id = ?", userId); err != nil {
		return err
	}

	if len(members) == 0 {
		return nil
	}

	err = u.db.Model(&data).Association("Members").Delete(members[0])
	return err
}

func NewTodoRepoNoCache(db *gorm.DB) TodoRepo {
	return TodoRepoNoCache{db: db}
}
