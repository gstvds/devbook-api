package models

type Follower struct {
	UserId     uint64 `json:"user_id,omitempty"`
	User       User   `gorm:"primaryKey,constraint:OnDelete:CASCADE" json:"user,omitempty"`
	FollowerId uint64 `json:"follower_id,omitempty"`
	Follower   User   `gorm:"primaryKey,constraint:OnDelete:CASCADE" json:"follower,omitempty"`
}
