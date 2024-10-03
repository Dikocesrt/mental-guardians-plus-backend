package video

type Video struct {
	ID        string `gorm:"primaryKey;not null;type:varchar(100);name:id"`
	VideoID   string `gorm:"type:text;name:video_id"`
	Title     string `gorm:"type:text;name:title"`
	Author    string `gorm:"type:text;name:author"`
	Views     int    `gorm:"type:bigint(20);name:views"`
	Likes     int    `gorm:"type:double;name:likes"`
	Comments  int    `gorm:"type:double;name:comments"`
	Labels    string `gorm:"type:text;name:labels"`
	Metadata  string `gorm:"type:text;name:metadata"`
	Thumbnail string `gorm:"type:varchar(100);name:thumbnail"`
}