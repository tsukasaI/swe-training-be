package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type QueryParam struct {
	// userIdはintで指定してほしい テスト仕様書の変更したほうがよい。 取り急ぎコメント
	UserId int `form:"userId" binding:"required"`
}

func main() {
	engine := setUpRouter()
	engine.Run(":8080")
}

func setUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.GET("/home", getHome)

	return engine
}

func getHome(c *gin.Context) {
	var queryParam QueryParam
	if c.ShouldBind(&queryParam) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userを指定してください。"})
		return
	}
	userId := c.Query("userId")

	db, err := connectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := findUser(db, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "userが存在しません。"})
		return
	}
	posts, err := getHomeData(db, user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func findUser(db *gorm.DB, userId string) (User, error) {
	var user User
	if err := db.Preload("Follows").First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

// database.go
func connectDb() (*gorm.DB, error) {
	dsn := "docker:docker@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

// user.go

// jsonのfield変更のために以下の書き方したらマイグレーション通らなくなった
// type commonGormModel struct {
// 	ID        uint           `gorm:"primaryKey" json:"id"`
// 	CreatedAt time.Time      `json:"createdAt"`
// 	UpdatedAt time.Time      `json:"updatedAt"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
// }

type Post struct {
	gorm.Model
	Comment string `gorm:"type:varchar(200) not null" json:"comment"`
	UserID  uint   `json:"-"`
	User    User   `json:"writer"`
}

type CommonResponseField struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostResponse struct {
	CommonResponseField
	Comment string `json:"comment"`
	User    UserResponse
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50) not null" json:"name"`
	Email    string `gorm:"type:varchar(100) not null unique" json:"email"`
	Password string `gorm:"type:varchar(255) not null" json:"-"`
	Posts    []Post `json:"posts"`
	Follows  []User `gorm:"many2many:user_follows"`
}

type UserResponse struct {
	CommonResponseField
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Posts   *[]PostResponse `json:"posts"`
	Follows *[]UserResponse `json:"follows"`
}

// home.go
func getHomeData(db *gorm.DB, user User) ([]PostResponse, error) {
	followIds := getFollowIds(user)

	var posts []Post
	db.Where("`user_id` in ?", followIds).Or("user_id = ?", user.ID).Preload("User").Find(&posts)
	var postsResponse []PostResponse
	for _, post := range posts {
		response := post.CreatePostResponse()
		postsResponse = append(postsResponse, response)
	}

	return postsResponse, nil
}

func getFollowIds(user User) []uint {
	var followIds []uint
	for _, followUser := range user.Follows {
		followIds = append(followIds, followUser.ID)
	}
	return followIds
}

func (post *Post) CreatePostResponse() PostResponse {
	postResponse := PostResponse{}
	postResponse.Id = post.ID
	postResponse.Comment = post.Comment
	postResponse.CreatedAt = post.CreatedAt.Format("2006/01/02/15/04/05")
	postResponse.UpdatedAt = post.UpdatedAt.Format("2006/01/02/15/04/05")
	postResponse.User = post.User.CreateUserResponse()
	return postResponse
}

func (user *User) CreateUserResponse() UserResponse {
	userResponse := UserResponse{}
	userResponse.Id = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.CreatedAt = user.CreatedAt.Format("2006/01/02/15/04/05")
	userResponse.UpdatedAt = user.UpdatedAt.Format("2006/01/02/15/04/05")
	return userResponse
}
