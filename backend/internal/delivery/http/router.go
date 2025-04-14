package http

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/handler"
	"github.com/shironxn/inkarya/internal/delivery/http/middleware"
)

type Router struct {
	app     *fiber.App
	version string
	jwksURL string
	handler *Handler
}

type Handler struct {
	Health     handler.HealthHandler
	User       handler.UserHandler
	Forum      handler.ForumHandler
	Course     handler.CourseHandler
	Job        handler.JobHandler
	Post       handler.PostHandler
	Skill      handler.SkillHandler
	Disability handler.DisabilityHandler
}

func NewRouter(app *fiber.App, version string, jwksURL string, handler *Handler) *Router {
	return &Router{
		app:     app,
		version: version,
		jwksURL: jwksURL,
		handler: handler,
	}
}

func (r *Router) Setup() {
	// Setup Swagger
	cfg := swagger.Config{
		BasePath: "/api/v" + r.version,
		FilePath: "./docs/swagger.yaml",
		Path:     "docs",
		Title:    "Inkarya API Documentation",
	}
	r.app.Use(swagger.New(cfg))

	// Base API group
	api := r.app.Group("/api/v" + r.version)

	// Public routes
	public := api.Group("")

	// Health check
	public.Get("/health", r.handler.Health.Check)

	// Setup routes by domain
	r.publicRoutes(public)
	r.privateRoutes(api)
}

func (r *Router) publicRoutes(router fiber.Router) {
	// User routes
	users := router.Group("/users")
	users.Get("/", r.handler.User.GetAllUsers)
	users.Get("/:id", r.handler.User.GetUserByID)

	// Forum routes
	forums := router.Group("/forums")
	forums.Get("/categories", r.handler.Forum.GetAllCategories)
	forums.Get("/", r.handler.Forum.GetAllForums)
	forums.Get("/:id/comments", r.handler.Forum.GetCommentsByForumID)
	forums.Get("/:id", r.handler.Forum.GetForumByID)

	// Course routes
	courses := router.Group("/courses")
	courses.Get("/", r.handler.Course.GetAllCourses)
	courses.Get("/:id/lessons/:lesson_id", r.handler.Course.GetLessonByID)
	courses.Get("/:id/enroll", r.handler.Course.GetEnrollByCourseID)
	courses.Get("/:id", r.handler.Course.GetCourseByID)

	// Job routes
	jobs := router.Group("/jobs")
	jobs.Get("/search", r.handler.Job.SearchJobs)
	jobs.Get("/company/:id", r.handler.Job.GetJobsByCompanyID)
	jobs.Get("/", r.handler.Job.GetAllJobs)
	jobs.Get("/:id", r.handler.Job.GetJobByID)

	// Post routes
	posts := router.Group("/posts")
	posts.Get("/", r.handler.Post.GetAllPosts)
	posts.Get("/:id/comments", r.handler.Post.GetCommentsByPostID)
	posts.Get("/:id", r.handler.Post.GetPostByID)

	// Skill routes
	skills := router.Group("/skills")
	skills.Get("/", r.handler.Skill.GetAll)
	skills.Get("/:id", r.handler.Skill.GetByID)

	// Disability routes
	disabilities := router.Group("/disabilities")
	disabilities.Get("/", r.handler.Disability.GetAll)
	disabilities.Get("/:id", r.handler.Disability.GetByID)
}

func (r *Router) privateRoutes(router fiber.Router) {
	// Apply JWT middleware to all private routes
	private := router.Group("", middleware.JWT(r.jwksURL))

	// Profile routes
	profile := private.Group("/profile")
	profile.Get("/", r.handler.User.GetMe)
	profile.Get("/enroll", r.handler.Course.GetEnrollByUserID)
	profile.Get("/jobs", r.handler.Job.GetJobApplicationsByUserID)

	// User routes
	users := private.Group("/users")
	users.Post("/", r.handler.User.CreateUser)
	users.Put("/", r.handler.User.UpdateUser)
	users.Delete("/", r.handler.User.DeleteUser)

	// Forum routes
	forums := private.Group("/forums")
	forums.Post("/", r.handler.Forum.CreateForum)
	forums.Put("/:id", r.handler.Forum.UpdateForum)
	forums.Delete("/:id", r.handler.Forum.DeleteForum)

	// Forum comments
	comments := forums.Group("/comments")
	comments.Post("/", r.handler.Forum.CreateComment)
	comments.Put("/:id", r.handler.Forum.UpdateComment)
	comments.Delete("/:id", r.handler.Forum.DeleteComment)

	// Course routes
	courses := private.Group("/courses")
	courses.Post("/:id/enroll", r.handler.Course.EnrollCourse)
	courses.Delete("/:id/enroll", r.handler.Course.UnenrollCourse)

	// Job routes
	jobs := private.Group("/jobs")

	// Job applications
	applications := jobs.Group("/applications")
	applications.Get("/", r.handler.Job.GetJobApplications)
	applications.Get("/:id", r.handler.Job.GetJobApplicationByID)
	applications.Post("/:id", r.handler.Job.ApplyForJob)

	// Saved jobs
	saved := jobs.Group("/saved")
	saved.Get("/", r.handler.Job.GetSavedJobs)
	saved.Post("/:id", r.handler.Job.SaveJob)
	saved.Delete("/:id", r.handler.Job.UnsaveJob)

	// Post routes
	posts := private.Group("/posts")
	posts.Post("/", r.handler.Post.CreatePost)
	posts.Put("/:id", r.handler.Post.UpdatePost)
	posts.Delete("/:id", r.handler.Post.DeletePost)
	posts.Post("/:id/like", r.handler.Post.LikePost)
	posts.Post("/:id/unlike", r.handler.Post.UnlikePost)

	// Post comments
	postComments := posts.Group("/comments")
	postComments.Post("/", r.handler.Post.CreateComment)
	postComments.Put("/:id", r.handler.Post.UpdateComment)
	postComments.Delete("/:id", r.handler.Post.DeleteComment)
}
