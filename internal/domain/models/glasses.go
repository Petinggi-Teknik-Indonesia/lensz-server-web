package models

import (
    "time"
    "gorm.io/gorm"
    "lensz-server-web/internal/domain/enums"
)

type Item struct {
    ID      string          `json:"id" gorm:"primaryKey;type:varchar(100)"`
    Drawer  string          `json:"drawer"`
    Color   string          `json:"color"`
    Type    string          `json:"type"`
    Brand   string          `json:"brand"`
    Company string          `json:"company"`
    Status  enums.ItemStatus `json:"status" gorm:"type:enum('Tersedia','Terjual','Rusak','Terpinjam','Lainnya');default:'Tersedia'"`

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
