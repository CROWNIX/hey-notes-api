package models

import (
	"time"
)

type Group struct {
    Id int 
    UserId int 
    Title string 
    IsPublic bool 
    Pin string
    CreatedAt time.Time
    UpdatedAt time.Time
}

