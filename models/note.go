package models

import (
	"time"
)

type Note struct {
    Id int 
    Title string 
    Slug string 
    Body string 
    Archived bool 
    CreatedAt time.Time
    UpdatedAt time.Time
}

