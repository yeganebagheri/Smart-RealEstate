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
) ([]*domain.Post, error) {
//) (res []string , error) {
	ctx := context.Background()
	//post := domain.Post{}
	//var res []interface{}

	// rows, err := repository.db.Query(ctx,"SELECT date,price FROM public.post where location = $1", getPostRequest.Location)
	// if err != nil {

	// 	return nil, err
	// }
    // defer rows.Close()

    // var tweet string
    // for rows.Next() {
    //     err := rows.Scan(&post.Date,&post.Price)
    //     if err != nil {

	// 		return nil, err
	// 	}
	// 	//fmt.Printf("this %s", tweet)
    //     res = append(res, tweet)
	// 	fmt.Printf("this %s", &post)
    // }

//3
	users := make([]*domain.Post, 0)
	rows, err := repository.db.Query(ctx,"SELECT date,price FROM public.post where location = $1",getPostRequest.Location)
	if err != nil {
    //log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
    user := new(domain.Post)
    if err := rows.Scan(&user.Date, &user.Price); err != nil {
        panic(err)
    }       
    users = append(users, user)
	}
	if err := rows.Err(); err != nil {
    panic(err)
	}
//end3

	//  err := repository.db.QueryRow(
	//  	ctx,
	//  	// "SELECT * FROM public.post where (price = $1 OR $1 IS NULL) and (location like %$2% OR $2 IS NULL) "+
	//  	// 	"and (title = $3 OR $3 IS NULL)",
	// 	 "SELECT date,price FROM public.post where location = $1" ,
	//  	//getPostRequest.Price,
	//  	getPostRequest.Location,
	//  	//getPostRequest.Title,
	//  ).Scan(
	//  	//&post.Id,
	//  	&post.Date,
	//  	&post.Price,
	//  	// &post.Location,
	//  	// &post.Image,
	//  	// &post.Title,
	//  	// &post.Description,
	//  )

	if err != nil {

		return nil, err
	}

	// return &post, nil
	// }

// 	err := repository.db.QueryRow(
// 		ctx,
// 		// "SELECT * FROM public.post where (price = $1 OR $1 IS NULL) and (location like %$2% OR $2 IS NULL) "+
// 		// 	"and (title = $3 OR $3 IS NULL)",
// 		"SELECT * FROM public.post" ,
// 		//getPostRequest.Price,
// 		//getPostRequest.Location,
// 		//getPostRequest.Title,
// 	).Scan(
// 		&post.Id,
// 		&post.Date,
// 		&post.Price,
// 		&post.Location,
// 		&post.Image,
// 		&post.Title,
// 		&post.Description,
// 	)

//    if err != nil {

// 	   return nil, err
//    }

   //return &post, nil
   return users, nil
}
