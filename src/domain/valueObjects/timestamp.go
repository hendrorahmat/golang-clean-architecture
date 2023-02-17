package valueObjects

import "time"

type TimestampValueObject struct {
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func (t *TimestampValueObject) CreatedAt() time.Time {
	return t.createdAt
}

func (t *TimestampValueObject) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *TimestampValueObject) DeletedAt() *time.Time {
	return t.deletedAt
}

func NewTimestampValueObject(createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) TimestampValueObject {
	return TimestampValueObject{createdAt: createdAt, updatedAt: updatedAt, deletedAt: deletedAt}
}

func (t *TimestampValueObject) IsDeleted() bool {
	return t.deletedAt != nil
}
