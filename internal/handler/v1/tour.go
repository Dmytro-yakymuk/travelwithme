package v1

import (
	"net/http"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func (h *Handler) initTourRoutes(api *gin.RouterGroup) {
	tours := api.Group("/tours")
	{
		tours.GET("", h.getAll)
		tours.GET("/:slug", h.getBySlug)

		toursPermissionAdmin := tours.Group("", h.userIdentity, h.admin)
		{
			toursPermissionAdmin.POST("/public/:slug", h.public)
		}

		toursPermissionOwner := tours.Group("", h.userIdentity, h.owner)
		{
			toursPermissionOwner.POST("/activ/:slug", h.activ)
		}
		toursPermission := tours.Group("", h.userIdentity, h.ownerAndAdmin)
		{
			toursPermission.POST("", h.create)
			toursPermission.PUT("/:slug", h.update)
			toursPermission.DELETE("/:slug", h.delete)
		}

	}
}

func (h *Handler) getAll(c *gin.Context) {
	urlQuery := c.Request.URL.Query()

	tours, err := h.services.Tour.GetAll(urlQuery)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	var newTours []interface{}
	for _, tour := range tours {

		images, err := h.services.Image.GetAllImageByTourId(tour.Id)
		if err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}

		trips, err := h.services.Trip.GetAllTripByTourId(c, rdb, tour.Id)
		if err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}

		user, err := h.services.User.GetOneByID(tour.UserId)
		if err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}
		// userJson, err := json.Marshal(user)
		// if err != nil {
		// 	newErrorResponse(c, http.StatusOK, err.Error())
		// 	return
		// }

		newTours = append(newTours, map[string]interface{}{
			"id":          tour.Id,
			"title":       tour.Title,
			"slug":        tour.Slug,
			"description": tour.Description,
			"public":      tour.Public,
			"activ":       tour.Activ,
			"minPrice":    trips["minPrice"],
			"countGroup":  trips["countGroup"],
			"images":      images,
			"trips":       trips["mapTrips"],
			"user":        user,
		})

	}

	c.JSON(http.StatusOK, newTours)
}

func (h *Handler) getBySlug(c *gin.Context) {

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusOK, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	images, err := h.services.Image.GetAllImageByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	trips, err := h.services.Trip.GetAllTripByTourId(c, rdb, tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	category, err := h.services.Category.GetById(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	region, err := h.services.Region.GetById(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	events, err := h.services.Event.GetAllEventByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	var eventIds []int
	for _, v := range events {
		eventIds = append(eventIds, v.Id)
	}

	comments, err := h.services.Comment.GetAllCommentsByTourId(tour.Id)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	generalStar := 0
	mapComments := []map[string]interface{}{}
	allCountStar := 0

	if len(comments) > 0 {
		for _, v := range comments {
			user, err := h.services.User.GetOneByID(v.UserId)
			if err != nil {
				newErrorResponse(c, http.StatusOK, err.Error())
				return
			}

			allCountStar += v.Star
			mapComment := map[string]interface{}{
				"id":         v.Id,
				"message":    v.Message,
				"star":       v.Star,
				"tour_id":    v.TourId,
				"user_name":  user.Name,
				"created_at": v.CreatedAt,
			}
			mapComments = append(mapComments, mapComment)
		}
	}

	generalStar = allCountStar / len(mapComments)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":              tour.Id,
		"title":           tour.Title,
		"slug":            tour.Slug,
		"description":     tour.Description,
		"text":            tour.Text,
		"category_id":     tour.CategoryId,
		"region_id":       tour.RegionId,
		"x":               tour.X,
		"y":               tour.Y,
		"minPrice":        trips["minPrice"],
		"countGroup":      trips["countGroup"],
		"trips":           trips["mapTrips"],
		"images":          images,
		"category":        category,
		"region":          region,
		"events":          events,
		"selected_events": eventIds,
		"comments":        mapComments,
		"generalStar":     generalStar,
	})
}

func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	var input models.CreateTourInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	tour := models.Tour{
		Title:       input.Title,
		Slug:        slug.Make(input.Title),
		Description: input.Description,
		Text:        input.Text,
		Public:      false,
		Activ:       input.Activ,
		X:           input.X,
		Y:           input.Y,
		CategoryId:  input.CategoryId,
		RegionId:    input.RegionId,
		UserId:      userId,
	}

	tourID, err := h.services.Tour.Create(tour)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	countToursEnvent := len(input.SelectedEvents)
	if countToursEnvent > 0 {
		toursEvents := make([]*models.ToursEvent, 0, countToursEnvent)
		for _, eventID := range input.SelectedEvents {
			toursEvent := &models.ToursEvent{}
			toursEvent.EventId = eventID
			toursEvent.TourId = tourID

			toursEvents = append(toursEvents, toursEvent)
		}
		if err = h.services.ToursEvent.Create(toursEvents); err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}
	}

	newSuccessResponse(c, http.StatusCreated,
		"Тур '"+tour.Title+"' створено успішно!",
		map[string]interface{}{
			"slug": slug.Make(input.Title),
		},
	)
}

func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusOK, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if userId != tour.UserId {
		newErrorResponse(c, http.StatusOK, "permission denied. not your record")
		return
	}

	var input models.UpdateTourInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if err := h.services.Tour.Update(slug, input); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	tourID := tour.Id
	countToursEnvent := len(input.SelectedEvents)
	if countToursEnvent > 0 {
		toursEvents := make([]*models.ToursEvent, 0, countToursEnvent)
		for _, eventID := range input.SelectedEvents {
			toursEvent := &models.ToursEvent{}
			toursEvent.EventId = eventID
			toursEvent.TourId = tourID

			toursEvents = append(toursEvents, toursEvent)
		}

		if err = h.services.ToursEvent.Delete(tourID); err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}

		if err = h.services.ToursEvent.Create(toursEvents); err != nil {
			newErrorResponse(c, http.StatusOK, err.Error())
			return
		}
	}

	newSuccessResponse(c, http.StatusCreated,
		"Тур '"+tour.Title+"' оновлено успішно!",
		nil,
	)
}

func (h *Handler) public(c *gin.Context) {
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusOK, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if userRole != "admin" {
		newErrorResponse(c, http.StatusOK, "permission denied. not your record")
		return
	}

	if err := h.services.Tour.Public(slug, !tour.Public); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if tour.Public {
		newSuccessResponse(c, http.StatusCreated,
			"Tур '"+tour.Title+"' не опубліковано!",
			nil,
		)
	} else {
		newSuccessResponse(c, http.StatusCreated,
			"Tур '"+tour.Title+"' опубліковано успішно!",
			nil,
		)
	}
}

func (h *Handler) activ(c *gin.Context) {
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusOK, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if userRole != "owner" || userId != tour.UserId {
		newErrorResponse(c, http.StatusOK, "permission denied. not your record")
		return
	}

	if err := h.services.Tour.Activ(slug, !tour.Activ); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if tour.Activ {
		newSuccessResponse(c, http.StatusCreated,
			"Tур '"+tour.Title+"' не активовано!",
			nil,
		)
	} else {
		newSuccessResponse(c, http.StatusCreated,
			"Tур '"+tour.Title+"' активовано успішно!",
			nil,
		)
	}
}

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}
	userRole, err := getUserRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	slug := c.Param("slug")
	if slug == "" {
		newErrorResponse(c, http.StatusOK, "empty slug param")
		return
	}

	tour, err := h.services.Tour.GetBySlug(slug)
	if err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	if userId != tour.UserId && userRole != "admin" {
		newErrorResponse(c, http.StatusOK, "permission denied. not your record")
		return
	}

	if err := h.services.Tour.Delete(slug); err != nil {
		newErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	newSuccessResponse(c, http.StatusCreated,
		"Tур '"+tour.Title+"' видалено успішно!",
		nil,
	)
}
