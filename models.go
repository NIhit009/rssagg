package main

import (
	"time"

	"github.com/NIhit009/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	API_KEY   string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		API_KEY:   dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedstoFeeds(DBfeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range DBfeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbfeed))
	}
	return feeds
}

type feedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_id"`
	UpdatedAt time.Time `json:"updated_id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databasefeedFollowtoFeedFollow(feedFollows database.FeedFollow) feedFollow {
	return feedFollow{
		ID:        feedFollows.ID,
		CreatedAt: feedFollows.CreatedAt,
		UpdatedAt: feedFollows.UpdatedAt,
		UserID:    feedFollows.UserID,
		FeedID:    feedFollows.FeedID,
	}
}

func databasefeedFollowstoFeedFollows(feedFollows []database.FeedFollow) []feedFollow {
	feed_follows := []feedFollow{}

	for _, feed_follow := range feedFollows {
		feed_follows = append(feed_follows, databasefeedFollowtoFeedFollow(feed_follow))
	}
	return feed_follows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePosttoPost(post database.Post) Post {
	var description *string
	if post.Description.Valid {
		description = &post.Description.String
	}
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Description: description,
		PublishedAt: post.PublishedAt,
		Url:         post.Url,
		FeedID:      post.FeedID,
	}
}

func databasePoststoPosts(posts []database.Post) []Post {
	newPosts := []Post{}
	for _, post := range posts {
		newPosts = append(newPosts, databasePosttoPost(post))
	}
	return newPosts
}
