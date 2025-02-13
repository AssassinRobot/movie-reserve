package model

type Reserve struct {
    ID        uint `gorm:"primary_key"`
    SeatID    uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    UserID    uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    CreatedAt string `gorm:"not null"`
}

func NewReserve(seatID, userID uint, createdAt string) *Reserve {
    return &Reserve{
        SeatID:    seatID,
        UserID:    userID,
        CreatedAt: createdAt,
    }
}