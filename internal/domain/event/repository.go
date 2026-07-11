package event

import (
	"errors"

	"gorm.io/gorm"
)

var ErrEventNotFound = errors.New("Event not found")

type Repository interface {
	Create(event *Event) error
	GetAll() ([]*Event, error)
	GetEventByID(eventId uint) (*Event, error)
	GetMyEvents(userId uint) ([]*Event, error)
	Update(event *Event) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) GetAll() ([]*Event, error) {
	var events []*Event

	err := r.db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *repository) GetEventByID(eventId uint) (*Event, error) {
	event := &Event{}

	err := r.db.First(event, eventId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, ErrEventNotFound
		}
		return nil, err
	}
	return event, nil
}

func (r *repository) GetMyEvents(userId uint) ([]*Event, error) {
	var events []*Event

	err := r.db.Where("user_id = ?", userId).Find(&events).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, ErrEventNotFound
		}
		return nil, err
	}
	return events, nil
}

func (r *repository) Update(event *Event) error {
	return r.db.Save(event).Error
}
