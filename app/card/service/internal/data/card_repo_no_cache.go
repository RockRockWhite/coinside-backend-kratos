package data

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CardRepoNoCache struct {
	db *gorm.DB
}

func (u CardRepoNoCache) Insert(ctx context.Context, data *Card) (uint64, error) {
	res := u.db.Create(data)
	return data.Id, res.Error
}

func (u CardRepoNoCache) FindOne(ctx context.Context, id uint64) (*Card, error) {
	var data Card
	res := u.db.Model(&data).Preload("Members").Preload("Tags").Where("id = ?", id).First(&data)
	return &data, res.Error
}

// FindAll 批量查询卡片信息
func (u CardRepoNoCache) FindAll(ctx context.Context, limit uint64, offset uint64, filters []Filter) ([]Card, uint64, error) {
	var datas []Card
	db := u.db.Model(&Card{})

	for _, f := range filters {
		db = f.Filter(db)
	}

	db = db.Preload("Members").Preload("Tags")

	var count int64
	db.Count(&count)
	
	res := db.Limit(int(limit)).Offset(int(offset)).Find(&datas)
	return datas, uint64(count), res.Error
}

func (u CardRepoNoCache) Update(ctx context.Context, newData *Card) error {
	res := u.db.Omit(clause.Associations).Save(newData)
	return res.Error
}

func (u CardRepoNoCache) Delete(ctx context.Context, id uint64) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	res := u.db.Select(clause.Associations).Delete(&data)
	return res.Error
}

func (u CardRepoNoCache) InsertTag(ctx context.Context, id uint64, content string) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	err = u.db.Model(&data).Association("Tags").Append(
		&Tag{
			Content: content,
		})
	return err
}

func (u CardRepoNoCache) DeleteTag(ctx context.Context, id uint64, content string) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var tags []Tag
	if err = u.db.Model(&data).Association("Tags").Find(&tags, "content = ?", content); err != nil {
		return err
	}

	if len(tags) == 0 {
		return nil
	}

	err = u.db.Model(&data).Association("Tags").Delete(tags[0])
	return err
}

func (u CardRepoNoCache) SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var members []Member
	if err = u.db.Model(&data).Association("Members").Find(&members, "user_id = ?", userId); err != nil {
		return err
	}

	// add a member
	if len(members) == 0 {
		err = u.db.Model(&data).Association("Members").Append(
			&Member{
				UserId:  userId,
				IsAdmin: admin,
			})
		return err
	}

	// update a member
	members[0].IsAdmin = admin
	res := u.db.Save(members[0])
	return res.Error
}

func (u CardRepoNoCache) DeleteMember(ctx context.Context, id uint64, userId uint64) error {
	data, err := u.FindOne(ctx, id)
	if err != nil {
		return err
	}

	var members []Member
	if err = u.db.Model(&data).Association("Members").Find(&members, "user_id = ?", userId); err != nil {
		return err
	}

	if len(members) == 0 {
		return nil
	}

	err = u.db.Model(&data).Association("Members").Delete(members[0])
	return err
}

func NewCardRepoNoCache(db *gorm.DB) CardRepo {
	return CardRepoNoCache{db: db}
}
