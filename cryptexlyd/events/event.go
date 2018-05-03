/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
package events

import (
	"fmt"
	"time"
)

type Event struct {
	Name        string
	Time        time.Time
	Title       string
	Description string
}

func New(name, title, description string) Event {
	return Event{
		Name:        name,
		Time:        time.Now(),
		Title:       title,
		Description: description,
	}
}

func (e Event) String() string {
	return fmt.Sprintf("%s{when:%s what:'%s'}", e.Name, e.Time, e.Title)
}
