package model

type Resource struct {
	Id       int     `json:"id" bson:"_id"`
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

type Album struct {
	Id     int    `json:"id" bson:"_id"`
	UserId int16  `json:"userId" bson:"userId" validate:"required"`
	Title  string `json:"title" bson:"title" validate:"required"`
}

type Comment struct {
	Id     int    `json:"id" bson:"_id"`
	PostId int16  `json:"postId" bson:"postId" validate:"required"`
	Name   string `json:"name" bson:"name" validate:"required"`
	Email  string `json:"email" bson:"email" validate:"required"`
	Body   string `json:"body" bson:"body" validate:"required"`
}

type Post struct {
	Id     int    `json:"id" bson:"_id"`
	UserId int16  `json:"userId" bson:"userId" validate:"required"`
	Title  string `json:"title" bson:"title" validate:"required"`
	Body   string `json:"body" bson:"body" validate:"required"`
}

type Photo struct {
	Id           int    `json:"id" bson:"_id"`
	AlbumId      int16  `json:"albumId" bson:"albumId" validate:"required"`
	Title        string `json:"title" bson:"title" validate:"required"`
	Url          string `json:"url" bson:"url" validate:"required"`
	ThumbnailUrl string `json:"thumbnailUrl" bson:"thumbnailUrl" validate:"required"`
}

type Todo struct {
	Id        int    `json:"id" bson:"_id"`
	UserId    int16  `json:"userId" bson:"userId" validate:"required"`
	Title     string `json:"title" bson:"title" validate:"required"`
	Completed bool   `json:"completed" bson:"completed" validate:"required"`
}

type User struct {
	Id       int    `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Username string `json:"username" bson:"username" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required"`
	Address  struct {
		Street  string `json:"street" bson:"street" validate:"required"`
		Suite   string `json:"suite" bson:"suite" validate:"required"`
		City    string `json:"city" bson:"city" validate:"required"`
		ZipCode string `json:"zipcode" bson:"zipcode" validate:"required"`
		Geo     struct {
			Lat string `json:"lat" bson:"lat" validate:"required"`
			Lng string `json:"lng" bson:"lng" validate:"required"`
		} `json:"geo" bson:"geo" validate:"required"`
	} `json:"address" bson:"address" validate:"required"`
	Phone   string `json:"phone" bson:"phone" validate:"required"`
	Website string `json:"website" bson:"website" validate:"required"`
	Company struct {
		Name        string `json:"name" bson:"name" validate:"required"`
		CatchPhrase string `json:"catchPhrase" bson:"catchPhrase" validate:"required"`
		Bs          string `json:"bs" bson:"bs" validate:"required"`
	} `json:"company" bson:"company" validate:"required"`
}
