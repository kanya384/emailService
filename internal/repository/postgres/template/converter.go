package template

import (
	"emailservice/internal/domain/template"
	"emailservice/internal/domain/template/path"
	"emailservice/internal/repository/postgres/template/dao"
)

func (r Repository) toDomainTemplate(dao *dao.Template) (result *template.Template, err error) {
	path, err := path.NewPath(dao.Path)
	if err != nil {
		return
	}
	result, err = template.NewWithID(
		dao.ID,
		*path,
		dao.CreatedAt,
		dao.ModifiedAt,
	)
	return
}
