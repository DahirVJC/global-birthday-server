package services

import (
	"fmt"
	"global-birthday-server/models"
	"global-birthday-server/repositories"
	"slices"
	"time"

	"github.com/google/uuid"
)

type BirthdaysService struct {
	birthdaysRepo repositories.BirthdaysRepository
	usersRepo     repositories.UsersRepository
}

func NewBirthdaysService(birthdaysRepo repositories.BirthdaysRepository, usersRepo repositories.UsersRepository) *BirthdaysService {
	return &BirthdaysService{
		birthdaysRepo: birthdaysRepo,
		usersRepo:     usersRepo,
	}
}

func normalizeBirthdate(birthdate time.Time) time.Time {
	return time.Date(2000, birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.UTC)
}

func toBirthdayDb(req models.BirthdayReq) models.Birthday {
	contactLinks := req.ContactLinks
	if contactLinks == nil {
		contactLinks = []string{}
	}

	wishlists := req.Wishlists
	if wishlists == nil {
		wishlists = []string{}
	}

	return models.Birthday{
		ID:           uuid.New(),
		Name:         req.Name,
		Birthdate:    normalizeBirthdate(req.Birthdate),
		Timezone:     req.Timezone,
		RemindMeIn:   req.RemindMeIn,
		ContactLinks: contactLinks,
		Wishlists:    wishlists,
		Events:       []uuid.UUID{},
	}
}

func emptyStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

func emptyUUIDs(values []uuid.UUID) []uuid.UUID {
	if values == nil {
		return []uuid.UUID{}
	}
	return values
}

func toBirthdayRes(birthday models.Birthday, startDate time.Time, endDate time.Time) models.BirthdayRes {
	return models.BirthdayRes{
		ID:           birthday.ID,
		Name:         birthday.Name,
		Birthdate:    birthday.Birthdate,
		Timezone:     birthday.Timezone,
		RemindMeIn:   birthday.RemindMeIn,
		ContactLinks: emptyStrings(birthday.ContactLinks),
		Wishlists:    emptyStrings(birthday.Wishlists),
		Events:       emptyUUIDs(birthday.Events),
		StartDate:    startDate,
		EndDate:      endDate,
	}
}

func ownsBirthday(user models.User, id uuid.UUID) bool {
	return slices.Contains(user.Birthdays, id.String())
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func birthdayDayForYear(birthdate time.Time, year int) (time.Month, int) {
	month := birthdate.Month()
	day := birthdate.Day()

	if month == time.February && day == 29 && !isLeapYear(year) {
		day = 28
	}

	return month, day
}

func birthdayInterval(birthday models.Birthday, year int, userTimezone string) (time.Time, time.Time, error) {
	contactLoc, err := time.LoadLocation(birthday.Timezone)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("malformed timezone %s", birthday.Timezone)
	}

	userLoc, err := time.LoadLocation(userTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("malformed timezone %s", userTimezone)
	}

	month, day := birthdayDayForYear(birthday.Birthdate, year)
	start := time.Date(year, month, day, 0, 0, 0, 0, contactLoc)
	end := start.AddDate(0, 0, 1)

	return start.In(userLoc), end.In(userLoc), nil
}

func (bs BirthdaysService) ListByYear(userId uuid.UUID, year int) (*[]models.BirthdayRes, error) {
	user, err := bs.usersRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	var results []models.BirthdayRes

	for _, birthdayIdStr := range user.Birthdays {
		birthdayId, err := uuid.Parse(birthdayIdStr)
		if err != nil {
			continue
		}

		birthdayDb, err := bs.birthdaysRepo.GetById(birthdayId)
		if err != nil || birthdayDb == nil {
			continue
		}

		startDate, endDate, err := birthdayInterval(*birthdayDb, year, user.Timezone)
		if err != nil {
			return nil, err
		}

		results = append(results, toBirthdayRes(*birthdayDb, startDate, endDate))
	}

	if results == nil {
		results = []models.BirthdayRes{}
	}

	return &results, nil
}

func (bs BirthdaysService) GetByYear(userId uuid.UUID, year int, birthdayId uuid.UUID) (*models.BirthdayRes, error) {
	user, err := bs.usersRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	if !ownsBirthday(*user, birthdayId) {
		return nil, fmt.Errorf("birthday with ID %s not found", birthdayId)
	}

	birthdayDb, err := bs.birthdaysRepo.GetById(birthdayId)
	if err != nil || birthdayDb == nil {
		return nil, fmt.Errorf("birthday with ID %s not found", birthdayId)
	}

	startDate, endDate, err := birthdayInterval(*birthdayDb, year, user.Timezone)
	if err != nil {
		return nil, err
	}

	res := toBirthdayRes(*birthdayDb, startDate, endDate)
	return &res, nil
}

func (bs BirthdaysService) Create(userId uuid.UUID, req models.BirthdayReq) error {
	user, err := bs.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	if _, err := time.LoadLocation(req.Timezone); err != nil {
		return fmt.Errorf("malformed timezone %s", req.Timezone)
	}

	birthdayDb := toBirthdayDb(req)

	if err := bs.birthdaysRepo.Create(birthdayDb); err != nil {
		return err
	}

	user.Birthdays = append(user.Birthdays, birthdayDb.ID.String())
	return bs.usersRepo.Update(userId, *user)
}

func (bs BirthdaysService) Update(userId uuid.UUID, birthdayId uuid.UUID, req models.BirthdayReq) error {
	user, err := bs.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	if !ownsBirthday(*user, birthdayId) {
		return fmt.Errorf("birthday with ID %s not found", birthdayId)
	}

	if _, err := time.LoadLocation(req.Timezone); err != nil {
		return fmt.Errorf("malformed timezone %s", req.Timezone)
	}

	existing, err := bs.birthdaysRepo.GetById(birthdayId)
	if err != nil || existing == nil {
		return fmt.Errorf("birthday with ID %s not found", birthdayId)
	}

	birthdayDb := toBirthdayDb(req)
	birthdayDb.ID = existing.ID
	birthdayDb.Events = existing.Events

	return bs.birthdaysRepo.Update(birthdayId, birthdayDb)
}

func (bs BirthdaysService) Delete(userId uuid.UUID, birthdayId uuid.UUID) error {
	user, err := bs.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	if !ownsBirthday(*user, birthdayId) {
		return fmt.Errorf("birthday with ID %s not found", birthdayId)
	}

	if err := bs.birthdaysRepo.Delete(birthdayId); err != nil {
		return err
	}

	user.Birthdays = slices.DeleteFunc(user.Birthdays, func(id string) bool {
		return id == birthdayId.String()
	})

	return bs.usersRepo.Update(userId, *user)
}
