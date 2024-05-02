package models

import (
	"time"
)

type Group struct {
    Id int 
    Title string 
    UserId int 
    IsPublic bool 
    Pin string
    CreatedAt time.Time
    UpdatedAt time.Time
}

