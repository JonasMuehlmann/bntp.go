// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/repository.tpl

package tag_repository

import (
    "context"
	"github.com/JonasMuehlmann/bntp.go/domain"
)

type TagRepository interface {
	New(...any) (TagRepository, error)

	Add(context.Context, []domain.Tag) (numAffectedRecords int, newID int, err error)
	Replace(context.Context, []domain.Tag) error
	UpdateWhere(context.Context, domain.TagFilter, map[domain.TagField]domain.TagUpdateOperation) (numAffectedRecords int, err error)
	Delete(context.Context, []domain.Tag) error
	DeleteWhere(context.Context, domain.TagFilter) (numAffectedRecords int, err error)
	CountWhere(context.Context, domain.TagFilter) int
	CountAll(context.Context) int
	DoesExist(context.Context, domain.Tag) bool
	DoesExistWhere(context.Context, domain.TagFilter) bool
	GetWhere(context.Context, domain.TagFilter) []domain.Tag
	GetFirstWhere(context.Context, domain.TagFilter) domain.Tag
	GetAll(context.Context) []domain.Tag
}