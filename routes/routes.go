package routes

import (
	"maftuhin/kpop-api/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("KPOP APP")
	})
	app.Get("app-ads.txt", func(c *fiber.Ctx) error {
		return c.SendString("google.com, pub-9691140516799861, DIRECT, f08c47fec0942fa0")
	})
	// app.Get("update/artist", controller.UpdateArtist)
	// app.Get("update/song", controller.UpdateSong)
	// app.Post("update/track", controller.UpdateSoundTrack)
	// app.Post("update/link", controller.UpdateLink)
	// // Auth
	// app.Post("auth/login", controller.SocialLogin)
	// // Post Route
	// app.Get("post/latest", controller.LatestPost)
	// app.Get("post/search", controller.PostSearch)
	// app.Get("post/detail/:id", controller.PostDetail)
	// app.Post("post/mine", controller.MyPost)
	// app.Post("post/create", controller.PostCreate)
	// app.Post("post/update", controller.PostUpdate)
	// app.Post("post/delete", controller.PostDelete)

	// Song
	app.Get("song", controller.SearchSong)
	app.Get("song/artist/:code", controller.SongByCode)
	app.Get("song/latest", controller.LastUpdate)
	app.Get("song/most-visited", controller.MostVisited)
	app.Get("song/counter/:id", controller.Counter)
	app.Get("song/:id", controller.SongDetail)

	// Artist
	app.Get("artist", controller.ArtistSearch)
	app.Get("artist/:code", controller.ArtistDetail)

	// // User Route
	// app.Get("user/search", controller.UserSearch)
	// app.Get("user/profile/:uid", controller.UserProfile)
	// app.Get("user/post/:uid", controller.UserPost)
	// app.Post("user/update", controller.UserUpdate)
	// app.Post("user/detail", controller.UserDetail)
	// Other
	app.Get("language", controller.Language)
	app.Post("request", controller.SendRequest)
}
