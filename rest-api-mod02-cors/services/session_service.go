package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// SessionService handles session-related logic
type SessionService struct {
	store *session.Store
}

// NewSessionService initializes a new SessionService
func NewSessionService() *SessionService {
	// Initialize session store with default options
	store := session.New()
	return &SessionService{store: store}
}

// SaveSession saves data in the session
func (s *SessionService) SaveSession(c *fiber.Ctx, key string, value interface{}) error {
	// Get session
	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}

	// Set session value
	sess.Set(key, value)

	// Save session
	return sess.Save()
}

// GetSession retrieves data from the session
func (s *SessionService) GetSession(c *fiber.Ctx, key string) interface{} {
	// Get session
	sess, err := s.store.Get(c)
	if err != nil {
		return nil
	}

	// Return session value
	return sess.Get(key)
}

// ClearSession clears a specific session key
func (s *SessionService) ClearSession(c *fiber.Ctx, key string) error {
	// Get session
	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}

	// Delete the session key
	sess.Delete(key)

	// Save session
	return sess.Save()
}
