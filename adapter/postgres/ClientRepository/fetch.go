package ClientRepository

import (
	"context"

	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
	//"github.com/yeganebagheri/Smart-RealEstate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	users := []domain.User{}
	total := int32(0)

	// query, queryCount, _ := paginate.Paginate("SELECT * FROM user").
	// 	Page(pagination.Page).
	// 	Desc(pagination.Descending).
	// 	Sort(pagination.Sort).
	// 	RowsPerPage(pagination.ItemsPerPage).
	// 	SearchBy(pagination.Search, "phone").
	// 	Query()
	//query := "SELECT * FROM movies"
	//{
	// rows, err := repository.db.Query(
	// 	ctx,
	// 	query,
	// )

	rows, err := repository.db.Query(ctx, "SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := domain.User{}

		rows.Scan(
			&user.Id,
			&user.Username,
			&user.PhonNo,
			&user.Password,
		)

		users = append(users, user)
	}
	//}

	// {
	// 	err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &domain.Pagination{
		Items: users,
		Total: total,
	}, nil
}
