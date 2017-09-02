package v2handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/die-net/lrucache"
	"github.com/xenu256/qprob_goapi/api_server/database"
	"github.com/xenu256/qprob_goapi/api_server/models"
)

var cache = lrucache.New(104857600*3, 60*60*24) //300 Mb, 24 hours
var postsPerPage = 20
var catsPerPage = 40

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := url.QueryEscape(strings.Split(r.RequestURI, "/")[3])
	p, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	cached, isCached := cache.Get("posts_" + page)
	if isCached == false {
		db := database.Connect()
		defer db.Close()

		query := fmt.Sprintf(`SELECT 
			posts.title, 
			posts.slug, 
			posts.url, 
			posts.summary, 
			posts.date, 
			posts.sentiment, 
			COALESCE(posts.image, ""), 
			COALESCE(posts.wordcloud, ""), 
			posts.category_id, 
			cats.slug, 
			COALESCE(cats.thumbnail, ""), 
			posts.hits, 
			(SELECT 
				COUNT(*) 
				FROM aggregator_post) 
			FROM aggregator_post as posts
			INNER JOIN aggregator_category AS cats ON posts.category_id = cats.title 
			ORDER BY posts.date DESC 
			LIMIT %[1]d,%[2]d;`, postsPerPage*p, postsPerPage)

		rows, err := db.Query(query)
		if err != nil {
			return
		}
		defer rows.Close()

		posts := make([]models.Post, 0)
		for rows.Next() {
			post := models.Post{}
			err := rows.Scan(&post.Title, &post.Slug, &post.URL, &post.Summary, &post.Date,
				&post.Sentiment, &post.Image, &post.Wordcloud, &post.CategoryID.Title, &post.CategoryID.Slug,
				&post.CategoryID.Thumbnail, &post.Hits, &post.TotalPosts)
			if err != nil {
				return
			}
			posts = append(posts, post)
		}
		if err = rows.Err(); err != nil {
			return
		}

		j, err := json.Marshal(posts)
		if err != nil {
			return
		}

		cache.Set("posts_"+page, j)
		w.Write(j)
	}
	w.Write(cached)
}

func PostsByCatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cat := url.QueryEscape(strings.Split(r.RequestURI, "/")[3])
	if len(cat) == 0 {
		return
	}
	page := url.QueryEscape(strings.Split(r.RequestURI, "/")[4])
	p, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	cached, isCached := cache.Get("posts_cat_" + cat + "_" + page)
	if isCached == false {
		db := database.Connect()
		defer db.Close()

		query := fmt.Sprintf(`SELECT 
			posts.title, 
			posts.slug, 
			posts.url, 
			posts.summary, 
			CASE posts.dead 
				WHEN 0 THEN "" 
				WHEN 1 THEN posts.content 
			END AS content, 
			posts.date AS dt, 
			posts.sentiment, 
			COALESCE(posts.image, ""), 
			COALESCE(posts.wordcloud, ""), 
			posts.category_id, 
			cats.slug, 
			COALESCE(cats.thumbnail, ""), 
			posts.hits, 
			(SELECT 
				COUNT(*) 
				FROM aggregator_post AS posts 
				INNER JOIN aggregator_category AS cats ON posts.category_id = cats.title 
				WHERE cats.slug='%[1]s'), 
			posts.dead 
			FROM aggregator_post AS posts 
			INNER JOIN aggregator_category AS cats ON posts.category_id = cats.title 
			WHERE cats.slug='%[1]s' 
			ORDER BY dt DESC 
			LIMIT %[2]d,%[3]d;`, cat, postsPerPage*p, postsPerPage)
		rows, err := db.Query(query)
		if err != nil {
			return
		}
		defer rows.Close()

		posts := make([]models.Post, 0)
		for rows.Next() {
			post := models.Post{}
			err := rows.Scan(&post.Title, &post.Slug, &post.URL, &post.Summary, &post.Content, &post.Date,
				&post.Sentiment, &post.Image, &post.Wordcloud, &post.CategoryID.Title, &post.CategoryID.Slug,
				&post.CategoryID.Thumbnail, &post.Hits, &post.TotalPosts, &post.Status)
			if err != nil {
				return
			}
			posts = append(posts, post)
		}
		if err = rows.Err(); err != nil {
			return
		}

		j, err := json.Marshal(posts)
		if err != nil {
			return
		}

		cache.Set("posts_cat_"+cat+"_"+page, j)
		w.Write(j)
	}
	w.Write(cached)
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := url.QueryEscape(strings.Split(r.RequestURI, "/")[3])
	p, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	cached, isCached := cache.Get("cats_" + page)
	if isCached == false {
		db := database.Connect()
		defer db.Close()

		query := fmt.Sprintf(`SELECT 
			cats.title, 
			cats.slug, 
			COALESCE(cats.thumbnail, ""), 
			COUNT(posts.title) AS cnt, 
			(SELECT 
				COUNT(*) 
				FROM aggregator_category) 
			FROM aggregator_category AS cats 
			INNER JOIN aggregator_post AS posts ON posts.category_id = cats.title 
			GROUP BY cats.title 
			LIMIT %[1]d,%[2]d;`, catsPerPage*p, catsPerPage)

		rows, err := db.Query(query)
		if err != nil {
			return
		}
		defer rows.Close()

		categories := make([]models.Category, 0)
		for rows.Next() {
			category := models.Category{}

			err := rows.Scan(&category.Title, &category.Slug, &category.Thumbnail,
				&category.PostCnt, &category.TotalCats)
			if err != nil {
				return
			}

			categories = append(categories, category)
		}
		if err = rows.Err(); err != nil {
			return
		}

		j, err := json.Marshal(categories)
		if err != nil {
			return
		}

		cache.Set("cats_"+page, j)
		w.Write(j)
	}
	w.Write(cached)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postSlug := url.QueryEscape(strings.Split(r.RequestURI, "/")[3])
	if len(postSlug) == 0 {
		return
	}

	cached, isCached := cache.Get("post_" + postSlug)
	if isCached == false {
		db := database.Connect()
		defer db.Close()

		query := fmt.Sprintf(`SELECT 
			posts.title, 
			posts.slug, 
			posts.url, 
			posts.summary, 
			CASE posts.dead 
				WHEN 0 THEN "" 
				WHEN 1 THEN posts.content 
			END AS content, 
			posts.date, 
			posts.sentiment, 
			COALESCE(posts.image, ""), 
			COALESCE(posts.wordcloud, ""), 
			posts.category_id, 
			cats.slug, 
			COALESCE(cats.thumbnail, ""), 
			posts.hits, 
			posts.dead 
			FROM aggregator_post as posts 
			INNER JOIN aggregator_category as cats ON posts.category_id = cats.title 
			WHERE posts.slug='%s';`, postSlug)
		rows, err := db.Query(query)
		if err != nil {
			return
		}
		defer rows.Close()

		posts := make([]models.Post, 0)
		for rows.Next() {
			post := models.Post{}
			err := rows.Scan(&post.Title, &post.Slug, &post.URL, &post.Summary,
				&post.Content, &post.Date, &post.Sentiment, &post.Image, &post.Wordcloud,
				&post.CategoryID.Title, &post.CategoryID.Slug, &post.CategoryID.Thumbnail,
				&post.Hits, &post.Status)
			if err != nil {
				return
			}
			posts = append(posts, post)
		}
		if err = rows.Err(); err != nil {
			return
		}

		j, err := json.Marshal(posts)
		if err != nil {
			return
		}

		cache.Set("post_"+postSlug, j)
		w.Write(j)
	}
	w.Write(cached)

	db := database.Connect()
	defer db.Close()

	prep, err := db.Prepare(`UPDATE aggregator_post 
		SET hits = hits + 1 
		WHERE slug= ? ;`)
	if err != nil {
		return
	}
	defer prep.Close()
	rows, err := prep.Query(postSlug)
	if err != nil {
		return
	}
	defer rows.Close()
}
