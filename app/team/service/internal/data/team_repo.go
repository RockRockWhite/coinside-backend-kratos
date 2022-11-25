package data

import (
	"context"
	"gorm.io/gorm"
)

type TeamRepo interface {
	Insert(ctx context.Context, data *Team) (uint64, error)

	FindOne(ctx context.Context, id uint64) (*Team, error)
	FindAll(ctx context.Context, limit uint64, offset uint64, filters []Filter) ([]Team, uint64, error)

	Update(ctx context.Context, newData *Team) error

	Delete(ctx context.Context, id uint64) error

	SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error
	DeleteMember(ctx context.Context, id uint64, userId uint64) error
}

type Filter interface {
	Filter(db *gorm.DB) *gorm.DB
}

type UserFilter struct {
	userId   uint64
	adminOpt *FilterAdminOption
}

type FilterAdminOption struct {
	IsAdmin bool
}

func (u UserFilter) Filter(db *gorm.DB) *gorm.DB {
	if u.adminOpt != nil {
		return db.
			Joins("JOIN `t_member` ON `t_member`.`team_id` = `t_team`.`id` AND `t_member`.`user_id` = ? AND `t_member`.`is_admin` = ?", u.userId, u.adminOpt.IsAdmin)

	}

	return db.
		Joins("JOIN `t_member` ON `t_member`.`team_id` = `t_team`.`id` AND `t_member`.`user_id` = ?", u.userId)
}

func NewUserFilter(userId uint64, adminOpt *FilterAdminOption) Filter {
	return &UserFilter{userId: userId, adminOpt: adminOpt}
}
