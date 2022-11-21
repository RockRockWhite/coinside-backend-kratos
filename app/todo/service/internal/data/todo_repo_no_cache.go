package data

import (
	"gorm.io/gorm"
)

type TodoRepoNoCache struct {
	db *gorm.DB
}

//func (u TeamRepoNoCache) Insert(ctx context.Context, data *Team) (uint64, error) {
//	res := u.db.Create(data)
//	return data.Id, res.Error
//}
//
//func (u TeamRepoNoCache) FindOne(ctx context.Context, id uint64) (*Team, error) {
//	var team Team
//	res := u.db.Model(&team).Preload("Members").Where("id = ?", id).First(&team)
//
//	return &team, res.Error
//}
//
//func (u TeamRepoNoCache) Update(ctx context.Context, newData *Team) error {
//	res := u.db.Omit(clause.Associations).Save(newData)
//	return res.Error
//}
//
//func (u TeamRepoNoCache) Delete(ctx context.Context, id uint64) error {
//	data, err := u.FindOne(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	res := u.db.Select(clause.Associations).Delete(&data)
//	return res.Error
//}
//
//func (u TeamRepoNoCache) SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error {
//	data, err := u.FindOne(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	var members []TeamMember
//	if err = u.db.Model(&data).Association("Members").Find(&members, "user_id = ?", userId); err != nil {
//		return err
//	}
//
//	// add a member
//	if len(members) == 0 {
//		err = u.db.Model(&data).Association("Members").Append(
//			&TeamMember{
//				UserId:  userId,
//				IsAdmin: admin,
//			})
//		return err
//	}
//
//	// update a member
//	members[0].IsAdmin = admin
//	res := u.db.Save(members[0])
//	return res.Error
//}
//
//func (u TeamRepoNoCache) DeleteMember(ctx context.Context, id uint64, userId uint64) error {
//	data, err := u.FindOne(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	var members []TeamMember
//	if err = u.db.Model(&data).Association("Members").Find(&members, "user_id = ?", userId); err != nil {
//		return err
//	}
//
//	if len(members) == 0 {
//		return nil
//	}
//
//	err = u.db.Model(&data).Association("Members").Delete(members[0])
//	return err
//}

func NewTodoRepoNoCache(db *gorm.DB) TodoRepo {
	return TodoRepoNoCache{db: db}
}
