package services

import (
	"fmt"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"slices"
	"time"

	"github.com/google/uuid"
)

type EventsService struct {
	eventsRepo    repositories.EventsRepository
	birthdaysRepo repositories.BirthdaysRepository
	usersRepo     repositories.UsersRepository
}

func NewEventsService(
	eventsRepo repositories.EventsRepository,
	birthdaysRepo repositories.BirthdaysRepository,
	usersRepo repositories.UsersRepository) *EventsService {
	return &EventsService{
		eventsRepo:    eventsRepo,
		birthdaysRepo: birthdaysRepo,
		usersRepo:     usersRepo,
	}
}

func toEventDb(req models.EventReq) models.Event {
	return models.Event{
		ID:             uuid.New(),
		EventDateStart: req.EventDateStart,
		EventDateEnd:   req.EventDateEnd,
		RemindMeIn:     req.RemindMeIn,
		Description:    req.Description,
	}
}

func toEventRes(event models.Event, birthdayId uuid.UUID, startDate time.Time, endDate time.Time) models.EventRes {
	return models.EventRes{
		ID:             event.ID,
		BirthdayID:     birthdayId,
		EventDateStart: event.EventDateStart,
		EventDateEnd:   event.EventDateEnd,
		RemindMeIn:     event.RemindMeIn,
		Description:    event.Description,
		StartDate:      startDate,
		EndDate:        endDate,
	}
}

func wallClockInLocation(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

func eventInterval(event models.Event, contactTimezone string, userTimezone string) (time.Time, time.Time, error) {
	contactLoc, err := time.LoadLocation(contactTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("malformed timezone %s", contactTimezone)
	}

	userLoc, err := time.LoadLocation(userTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("malformed timezone %s", userTimezone)
	}

	start := wallClockInLocation(event.EventDateStart, contactLoc)
	end := start
	if event.EventDateEnd != nil {
		end = wallClockInLocation(*event.EventDateEnd, contactLoc)
	}

	return start.In(userLoc), end.In(userLoc), nil
}

func (es EventsService) findOwnedEventBirthday(user models.User, eventId uuid.UUID) (*models.Birthday, error) {
	for _, birthdayIdStr := range user.Birthdays {
		birthdayId, err := uuid.Parse(birthdayIdStr)
		if err != nil {
			continue
		}

		birthdayDb, err := es.birthdaysRepo.GetById(birthdayId)
		if err != nil || birthdayDb == nil {
			continue
		}

		if slices.Contains(birthdayDb.Events, eventId) {
			return birthdayDb, nil
		}
	}

	return nil, fmt.Errorf("event with ID %s not found", eventId)
}

func (es EventsService) List(userId uuid.UUID, includePast bool) (*[]models.EventRes, error) {
	user, err := es.usersRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	var results []models.EventRes
	now := time.Now()

	for _, birthdayIdStr := range user.Birthdays {
		birthdayId, err := uuid.Parse(birthdayIdStr)
		if err != nil {
			continue
		}

		birthdayDb, err := es.birthdaysRepo.GetById(birthdayId)
		if err != nil || birthdayDb == nil {
			continue
		}

		for _, eventId := range birthdayDb.Events {
			eventDb, err := es.eventsRepo.GetById(eventId)
			if err != nil || eventDb == nil {
				continue
			}

			startDate, endDate, err := eventInterval(*eventDb, birthdayDb.Timezone, user.Timezone)
			if err != nil {
				return nil, err
			}

			if !includePast && endDate.Before(now) {
				continue
			}

			results = append(results, toEventRes(*eventDb, birthdayDb.ID, startDate, endDate))
		}
	}

	if results == nil {
		results = []models.EventRes{}
	}

	return &results, nil
}

func (es EventsService) Get(userId uuid.UUID, eventId uuid.UUID) (*models.EventRes, error) {
	user, err := es.usersRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	birthdayDb, err := es.findOwnedEventBirthday(*user, eventId)
	if err != nil {
		return nil, err
	}

	eventDb, err := es.eventsRepo.GetById(eventId)
	if err != nil || eventDb == nil {
		return nil, fmt.Errorf("event with ID %s not found", eventId)
	}

	startDate, endDate, err := eventInterval(*eventDb, birthdayDb.Timezone, user.Timezone)
	if err != nil {
		return nil, err
	}

	res := toEventRes(*eventDb, birthdayDb.ID, startDate, endDate)
	return &res, nil
}

func (es EventsService) Create(userId uuid.UUID, req models.EventReq) error {
	user, err := es.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	if !ownsBirthday(*user, req.BirthdayID) {
		return fmt.Errorf("birthday with ID %s not found", req.BirthdayID)
	}

	birthdayDb, err := es.birthdaysRepo.GetById(req.BirthdayID)
	if err != nil || birthdayDb == nil {
		return fmt.Errorf("birthday with ID %s not found", req.BirthdayID)
	}

	eventDb := toEventDb(req)
	if err := es.eventsRepo.Create(eventDb); err != nil {
		return err
	}

	birthdayDb.Events = append(birthdayDb.Events, eventDb.ID)
	return es.birthdaysRepo.Update(req.BirthdayID, *birthdayDb)
}

func (es EventsService) Update(userId uuid.UUID, eventId uuid.UUID, req models.EventReq) error {
	user, err := es.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	if _, err := es.findOwnedEventBirthday(*user, eventId); err != nil {
		return err
	}

	existing, err := es.eventsRepo.GetById(eventId)
	if err != nil || existing == nil {
		return fmt.Errorf("event with ID %s not found", eventId)
	}

	eventDb := toEventDb(req)
	eventDb.ID = existing.ID

	return es.eventsRepo.Update(eventId, eventDb)
}

func (es EventsService) Delete(userId uuid.UUID, eventId uuid.UUID) error {
	user, err := es.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	birthdayDb, err := es.findOwnedEventBirthday(*user, eventId)
	if err != nil {
		return err
	}

	if err := es.eventsRepo.Delete(eventId); err != nil {
		return err
	}

	birthdayDb.Events = slices.DeleteFunc(birthdayDb.Events, func(id uuid.UUID) bool {
		return id == eventId
	})

	return es.birthdaysRepo.Update(birthdayDb.ID, *birthdayDb)
}
