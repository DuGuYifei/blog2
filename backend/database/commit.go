package database

import (
	"errors"
	"gorm.io/gorm"
)

type Commit struct {
	gorm.Model
	CommitValue string `gorm:"not null;size:40"`
}

func GetLastCommit() (string, error) {
	// 返回id=1的commit值
	var commit Commit
	if err := DB.First(&commit, 1).Error; err != nil {
		return "", err
	}
	return commit.CommitValue, nil
}

func SetLastCommit(commit string) error {
	// 更新id=1的commit值
	if len(commit) != 40 {
		return errors.New("commit value must be 40 characters")
	}
	return DB.Model(&Commit{}).Where("id = ?", 1).Update("commit_value", commit).Error
}
