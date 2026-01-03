package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/umrzoq-toshkentov/social/internal/store"
)

func Seed(store *store.Storage) {
	ctx := context.Background()
	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			fmt.Println("error creating users")
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			fmt.Println("error creating posts")
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			fmt.Println("error creating comments")
		}
	}

	log.Println("Completed seed")
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)
	for i := 0; i < count; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    fmt.Sprintf("%s%d@example.com", usernames[i%len(usernames)], i),
			Password: "123123",
		}
	}
	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)
	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]

		// Pick random tags (2-4 tags per post)
		numTags := rand.Intn(3) + 2
		postTags := make([]string, numTags)
		for j := 0; j < numTags; j++ {
			postTags[j] = tags[rand.Intn(len(tags))]
		}

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags:    postTags,
		}
	}
	return posts
}

func generateComments(count int, users []*store.User, posts []*store.Post) []*store.Comment {
	result := make([]*store.Comment, count)
	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		result[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return result
}
