package data

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VoteRepoNoCache struct {
	db *gorm.DB
}

func (u VoteRepoNoCache) Insert(ctx context.Context, data *Vote) (uint64, error) {
	res := u.db.Create(data)
	return data.Id, res.Error
}

func (u VoteRepoNoCache) FindOne(ctx context.Context, id uint64) (*Vote, error) {
	var data Vote
	res := u.db.Model(&data).Preload("Items").Where("id = ?", id).First(&data)
	return &data, res.Error
}

//
//// FindAll 批量查询卡片信息
//func (u CardRepoNoCache) FindAll(ctx context.Context, limit uint64, offset uint64, filters []Filter) ([]Card, error) {
//	var datas []Card
//	db := u.db.Model(&Card{})
//
//	for _, f := range filters {
//		db = f.Filter(db)
//	}
//
//	res := db.Preload("Members").Preload("Tags").Find(&datas)
//	return datas, res.Error
//}

func (u VoteRepoNoCache) Update(ctx context.Context, newData *Vote) error {
	res := u.db.Omit(clause.Associations).Save(newData)
	return res.Error
}

func (u VoteRepoNoCache) Delete(ctx context.Context, id uint64) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	res := u.db.Select(clause.Associations).Delete(&data)
	return res.Error
}

func (u VoteRepoNoCache) InsertItem(ctx context.Context, id uint64, content string) (uint64, error) {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return 0, err
	}

	item := &Item{
		Content: content,
	}
	err = u.db.Model(&data).Association("Items").Append(item)
	return item.Id, err
}

func (u VoteRepoNoCache) DeleteItem(ctx context.Context, id uint64, itemId uint64) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var items []Item
	if err = u.db.Model(&data).Association("Items").Find(&items, "item_id = ?", itemId); err != nil {
		return err
	}

	if len(items) == 0 {
		return nil
	}

	err = u.db.Model(&data).Association("Items").Delete(items[0])
	return err
}
func (u VoteRepoNoCache) UpdateItem(ctx context.Context, id uint64, itemId uint64, context string) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}
	var items []Item
	if err = u.db.Model(&data).Association("Items").Find(&items, "item_id = ?", itemId); err != nil {
		return err
	}

	if len(items) == 0 {
		return nil
	}

	items[0].Content = context
	//这里这里，，是不是这样要测试
	res := u.db.Save(items[0])
	return res.Error
}

func (u VoteRepoNoCache) InsertCommit(ctx context.Context, id uint64, itemId uint64, userId uint64) error {
	return nil
}
func (u VoteRepoNoCache) DeleteCommit(ctx context.Context, id uint64, itemId uint64, userId uint64) error {
	return nil
}

func NewVoteRepoNoCache(db *gorm.DB) VoteRepo {
	return VoteRepoNoCache{db: db}
}
