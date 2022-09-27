package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type PostsRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewPostsRepository(db infrastructure.Database, logger infrastructure.Logger) PostsRepository {
	return PostsRepository{
		db:     db,
		logger: logger,
	}
}

func (c PostsRepository) WithTrx(trxHandle *gorm.DB) PostsRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context.")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// CreatePosts -> Post
func (c PostsRepository) CreatePosts(posts models.Post) error {
	return c.db.DB.Create(&posts).Error
}

// UpdatePosts -> Post
func (c PostsRepository) UpdatePosts(Posts models.Post) error {
	return c.db.DB.Model(&models.Post{}).
		Where("id = ?", Posts.ID).
		Updates(map[string]interface{}{
			"title":    Posts.Title,
			"caption":  Posts.Caption,
			"user_id":  Posts.UserId,
			"audience": Posts.Audience,
		}).Error
}

// DeletePosts -> Post
func (c PostsRepository) DeletePosts(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.Post{}).Error
}

// GetOneUser -> gets one post of postId
func (c PostsRepository) GetOnePost(postId int64, userId string) (Posts models.UserPost, err error) {
	return Posts, c.db.DB.
		Model(&models.Post{}).
		Select(`posts.*,(SELECT COUNT(post_id)
		FROM post_likes JOIN posts p ON p.id = post_likes.post_id) like_count,
	   IF((SELECT c.user_id FROM post_likes c WHERE user_id = ?) = ?, TRUE, FALSE) has_liked`, userId, userId).
		Where("id = ?", postId).
		First(&Posts).
		Error
}

func (c PostsRepository) GetPost(postId int64) (Posts models.Post, err error) {
	return Posts, c.db.DB.
		Model(&models.Post{}).
		Where("id = ?", postId).
		First(&Posts).
		Error
}

// GetAllPosts -> Get All Posts
func (c PostsRepository) GetAllPosts(pagination utils.Pagination) ([]models.Post, int64, error) {
	var Posts []models.Post
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Post{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`Posts`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&Posts).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return Posts, totalRows, err
}

// GetCreatorPosts-> Get Creator Posts
func (c PostsRepository) CreatorPosts(cursorPagination utils.CursorPagination, userId string) (Posts []models.Post, err error) {

	parsedCursor, _ := time.Parse(time.RFC3339, cursorPagination.Cursor)
	queryBuilder := c.db.DB.Model(&models.Post{}).Select(`posts.*`).Where("user_id= ? ", userId).Limit(cursorPagination.PageSize)
	if cursorPagination.Cursor != "" {
		queryBuilder = queryBuilder.Where("created_at < ?", parsedCursor)
	}

	return Posts, queryBuilder.Order("created_at desc").Find(&Posts).
		Error

}

//GetUserFeed => Get Users Feeds

func (c PostsRepository) GetUserFeed(cursorPagination utils.CursorPagination, userId string) (Posts []models.FeedPost, err error) {
	parsedCursor, _ := time.Parse(time.RFC3339, cursorPagination.Cursor)
	queryBuilder := c.db.DB.Model(&models.FeedPost{}).Select(`posts.*,(SELECT COUNT(post_id) FROM post_likes WHERE posts.id = post_likes.post_id) like_count,
	(SELECT COUNT(comment_id) FROM comment_likes JOIN comments p ON p.id = comment_likes.comment_id) comment_like_count,
	IF((SELECT c.user_id FROM post_likes c WHERE user_id = ?) = ?, TRUE, FALSE) has_liked`, userId, userId).Joins(`join followers on followers.follow_user_id=posts.user_id`).
		Where(`posts.audience != 'private' and followers.user_id= ?`, userId).
		Limit(cursorPagination.PageSize)
	if cursorPagination.Cursor != "" {
		queryBuilder = queryBuilder.Where("created_at < ?", parsedCursor)
	}

	return Posts, queryBuilder.Order("created_at desc").Preload("PostContents").Preload("User").Find(&Posts).
		Error
}

func (c PostsRepository) UploadFile(fileName string) {
	ctx := context.Background()
	config := &firebase.Config{
		StorageBucket: "flutterproject-31436.appspot.com",
	}
	sa := option.WithCredentialsFile("./serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	wc := bucket.Object(fileName).NewWriter(ctx)
	if err := wc.Close(); err != nil {
		log.Println(err)
	}
}
