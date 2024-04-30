package jsonplaceholder

type Resource struct {
	Id       int16   `json:"id" bson:"id"`
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`
	Username *string `json:"username,omitempty" bson:"username,omitempty"`
	Email    *string `json:"email,omitempty" bson:"email,omitempty"`
	Address  *struct {
		Street  *string `json:"street,omitempty" bson:"street,omitempty"`
		Suite   *string `json:"suite,omitempty" bson:"suite,omitempty"`
		City    *string `json:"city,omitempty" bson:"city,omitempty"`
		ZipCode *string `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
		Geo     *struct {
			Lat *string `json:"lat,omitempty" bson:"lat,omitempty"`
			Lng *string `json:"lng,omitempty" bson:"lng,omitempty"`
		} `json:"geo,omitempty" bson:"geo,omitempty"`
	} `json:"address,omitempty" bson:"address,omitempty"`
	Phone   *string `json:"phone,omitempty" bson:"phone,omitempty"`
	Website *string `json:"website,omitempty" bson:"website,omitempty"`
	Company *struct {
		Name        *string `json:"name,omitempty" bson:"name,omitempty"`
		CatchPhrase *string `json:"catchPhrase,omitempty" bson:"catchPhrase,omitempty"`
		Bs          *string `json:"bs,omitempty" bson:"bs,omitempty"`
	} `json:"company,omitempty" bson:"company,omitempty"`
	PostId       *int16  `json:"postId,omitempty" bson:"postId,omitempty"`
	Body         *string `json:"body,omitempty" bson:"body,omitempty"`
	UserId       *int16  `json:"userId,omitempty" bson:"userId,omitempty"`
	Title        *string `json:"title,omitempty" bson:"title,omitempty"`
	AlbumId      *int16  `json:"albumId,omitempty" bson:"albumId,omitempty"`
	Url          *string `json:"url,omitempty" bson:"url,omitempty"`
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty" bson:"thumbnailUrl,omitempty"`
	Completed    *bool   `json:"completed,omitempty" bson:"completed,omitempty"`
}

// type Album struct {
// 	Resource
// }

// type Comment struct {
// 	Resource
// }

// type Photo struct {
// 	Resource
// }

// type Post struct {
// 	Resource
// }

// type Todo struct {
// 	Resource
// }

// type User struct {
// 	Resource
// }

type Album struct {
	Id     int16  `json:"id" bson:"id"`
	UserId int16  `json:"userId" bson:"userId"`
	Title  string `json:"title" bson:"title"`
}

type Comment struct {
	PostId int16  `json:"postId" bson:"postId"`
	Id     int16  `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
	Body   string `json:"body" bson:"body"`
}

type Post struct {
	UserId int16  `json:"userId" bson:"userId"`
	Id     int16  `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	Body   string `json:"body" bson:"body"`
}

type Photo struct {
	AlbumId      int16  `json:"albumId" bson:"albumId"`
	Id           int16  `json:"id" bson:"id"`
	Title        string `json:"title" bson:"title"`
	Url          string `json:"url" bson:"url"`
	ThumbnailUrl string `json:"thumbnailUrl" bson:"thumbnailUrl"`
}

type Todo struct {
	Id        int16  `json:"id" bson:"id"`
	UserId    int16  `json:"userId" bson:"userId"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

type User struct {
	Id       int16  `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Address  struct {
		Street  string `json:"street" bson:"street"`
		Suite   string `json:"suite" bson:"suite"`
		City    string `json:"city" bson:"city"`
		ZipCode string `json:"zipcode" bson:"zipcode"`
		Geo     struct {
			Lat string `json:"lat" bson:"lat"`
			Lng string `json:"lng" bson:"lng"`
		} `json:"geo" bson:"geo"`
	} `json:"address" bson:"address"`
	Phone   string `json:"phone" bson:"phone"`
	Website string `json:"website" bson:"website"`
	Company struct {
		Name        string `json:"name" bson:"name"`
		CatchPhrase string `json:"catchPhrase" bson:"catchPhrase"`
		Bs          string `json:"bs" bson:"bs"`
	} `json:"company" bson:"company"`
}
