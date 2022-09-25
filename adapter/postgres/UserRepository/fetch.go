package productrepository

import (
	"context"

	"github.com/boooscaaa/clean-go/core/domain"
	"github.com/boooscaaa/clean-go/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	users := []domain.User{}
	total := int32(0)

	query, queryCount, _ := paginate.Paginate("SELECT * FROM user").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "phone").
		Query()

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			user := domain.User{}

			rows.Scan(
				&user.ID,
				&user.Phone,
				&user.Image,
				&user.Password,
			)

			users = append(users, user)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination{
		Items: users,
		Total: total,
	}, nil
}
