package stores

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/iamjeremylim/myadmin-service/db/sqlc"
)

type Service struct {
	queries *db.Queries
}

type getStoreReq struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
}

type idParameters struct {
	ID int64 `uri:"id" binding:"required"`
}

type listStoreReq struct {
	Username string `form:"username" binding:"required"`
}

func NewService(queries *db.Queries) *Service {
	return &Service{queries: queries}
}

func fromDB(store db.Store) *getStoreReq {
	return &getStoreReq{
		ID:    store.ID,
		Name:  store.Name,
		Owner: store.Owner,
	}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/stores", s.List)
	router.POST("/stores", s.Create)
	router.GET("/stores/:id", s.Get)
	router.PATCH("/stores/:id", s.Update)
	router.DELETE("/stores/:id", s.Delete)
}

func (s *Service) List(c *gin.Context) {
	// Parse request
	var req listStoreReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	// List stores
	stores, err := s.queries.ListStores(context.Background(), req.Username)
	if err != nil {
		log.Println("Error parsing query:", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	var response []*getStoreReq

	if len(stores) == 0 {
		response = []*getStoreReq{}
	}

	for _, store := range stores {
		response = append(response, fromDB(store))
	}
	c.JSON(http.StatusOK, response)
}

func (s *Service) Create(c *gin.Context) {
	// Parse request
	var request getStoreReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := db.CreateStoreParams{
		Name:  request.Name,
		Owner: request.Owner,
	}

	// Create store
	store, err := s.queries.CreateStore(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(store)
	c.JSON(http.StatusCreated, response)
}

func (s *Service) Get(c *gin.Context) {
	// Parse request
	var pathParams idParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error from get": err.Error()})
		return
	}

	// Get store
	store, err := s.queries.GetStore(context.Background(), pathParams.ID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(store)
	c.JSON(http.StatusOK, response)
}

func (s *Service) Delete(c *gin.Context) {
	// Parse request
	var pathParams idParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete store
	err := s.queries.DeleteStore(context.Background(), pathParams.ID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	c.Status(http.StatusOK)
}

func (s *Service) Update(c *gin.Context) {
	// Parse request
	var request getStoreReq
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update store
	params := db.UpdateStoreParams{
		ID:   request.ID,
		Name: request.Name,
	}

	store, err := s.queries.UpdateStore(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(store)
	c.JSON(http.StatusOK, response)
}
