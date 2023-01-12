package PostRepository

import (
	"context"
	"github.com/yeganebagheri/Smart-RealEstate/core/domain"
	"github.com/yeganebagheri/Smart-RealEstate/core/dto"
	
)
// type FilterTransaction struct {
// 	//price           string `json:"Location"`
// 	location         string `json:"Price"`
// 	//financialAccountID string `json:"Title"`
// }

func (repository repository) Get(
	getPostRequest *dto.GetPostRequest,
) (*domain.Post, error) {
	ctx := context.Background()
	post := domain.Post{}

	//list := []Transaction{}
	//filter := FilterTransaction{}

	// if getPostRequest.Price != "" {
	// 	filter.price = getPostRequest.Price
	// }

	// if getPostRequest.Location != "" {
	// 	filter.location = getPostRequest.Location
	// }

	// if getPostRequest.Title != "" {
	// 	filter.financialAccountID = getPostRequest.Title
	// }

	// cursor, err := db.Find(ctx, filter)
	// if err = cursor.All(ctx, &list); err != nil {
	// 	return nil, errors.Wrapf(err, "filtering transactions")
	// }
	if getPostRequest.Location != "" {
	 err := repository.db.QueryRow(
	 	ctx,
	 	// "SELECT * FROM public.post where (price = $1 OR $1 IS NULL) and (location like %$2% OR $2 IS NULL) "+
	 	// 	"and (title = $3 OR $3 IS NULL)",
		 "SELECT * FROM public.post where location = $1 " ,
	 	//getPostRequest.Price,
	 	getPostRequest.Location,
	 	//getPostRequest.Title,
	 ).Scan(
	 	&post.Id,
	 	&post.Date,
	 	&post.Price,
	 	&post.Location,
	 	&post.Image,
	 	&post.Title,
	 	&post.Description,
	 )

	if err != nil {

		return nil, err
	}

	return &post, nil
	}

	err := repository.db.QueryRow(
		ctx,
		// "SELECT * FROM public.post where (price = $1 OR $1 IS NULL) and (location like %$2% OR $2 IS NULL) "+
		// 	"and (title = $3 OR $3 IS NULL)",
		"SELECT * FROM public.post" ,
		//getPostRequest.Price,
		//getPostRequest.Location,
		//getPostRequest.Title,
	).Scan(
		&post.Id,
		&post.Date,
		&post.Price,
		&post.Location,
		&post.Image,
		&post.Title,
		&post.Description,
	)

   if err != nil {

	   return nil, err
   }

   return &post, nil
}
