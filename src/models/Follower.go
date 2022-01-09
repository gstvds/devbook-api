package models

type Follower struct {
	UserId     uint64 `gorm:"primaryKey" json:"user_id,omitempty"`
	User       User   `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	FollowerId uint64 `gorm:"primaryKey" json:"follower_id,omitempty"`
	Follower   User   `gorm:"constraint:OnDelete:CASCADE" json:"follower,omitempty"`
}
