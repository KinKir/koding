package models

import (
	"fmt"
	"time"

	"github.com/koding/bongo"
)

func (c Channel) GetId() int64 {
	return c.Id
}

func (c Channel) TableName() string {
	return "api.channel"
}

func (c *Channel) AfterCreate() {
	bongo.B.AfterCreate(c)
}

func (c *Channel) AfterUpdate() {
	bongo.B.AfterUpdate(c)
}

func (c Channel) AfterDelete() {
	bongo.B.AfterDelete(c)
}

func (c *Channel) BeforeCreate() error {
	c.CreatedAt = time.Now().UTC()
	c.UpdatedAt = time.Now().UTC()
	c.DeletedAt = ZeroDate()

	return c.MarkIfExempt()
}

func (c *Channel) BeforeUpdate() error {
	c.UpdatedAt = time.Now()

	return c.MarkIfExempt()
}


func (c *Channel) Update() error {
	if c.Name == "" || c.GroupName == "" {
		return fmt.Errorf("Validation failed %s - %s", c.Name, c.GroupName)
	}

	return bongo.B.Update(c)
}

func (c *Channel) Delete() error {
	return bongo.B.Delete(c)
}

func (c *Channel) ById(id int64) error {
	return bongo.B.ById(c, id)
}

func (c *Channel) One(q *bongo.Query) error {
	return bongo.B.One(c, c, q)
}

func (c *Channel) Some(data interface{}, q *bongo.Query) error {
	return bongo.B.Some(c, data, q)
}
