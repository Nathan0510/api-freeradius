package app

import (
    "api-freeradius/models"
    "api-freeradius/db"
)

func CreateGroup(groupname, attribute, value string) error {
    group := &models.Radgroupreply{
        Groupname:  groupname,
        Attribute: attribute,
        Op: ":=",
        Value:     value,
    }

    return db.DB.Create(group).Error
}

func GetAllGroup() ([]models.Radgroupreply, error) {
    var groupList []models.Radgroupreply
    return groupList, db.DB.Find(&groupList).Error
}

func GetGroup(groupname string) (*models.Radgroupreply, error) {
    var group models.Radgroupreply
    return &group, db.DB. Where("groupname = ?", groupname).First(&group).Error
}

func DeleteGroup(groupname string) error {
    return db.DB.Where("groupname = ?", groupname).Delete(&models.Radgroupreply{}).Error
}

func UpdateGroup(groupname string, updates *models.Radgroupreply) error {
    return db.DB.Model(&models.Radgroupreply{}).Where("groupname = ?", groupname).Updates(updates).Error
}
