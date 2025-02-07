type Class struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=50,excludesall=|"`
	Grade     string    `json:"grade" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}