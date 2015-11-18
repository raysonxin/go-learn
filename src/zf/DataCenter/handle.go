package main

import (
	"fmt"
)

type DbServiceImpl struct {
}

//前置端登陆
func (d DbServiceImpl) FrontLogin(frontCode string, password string) (r bool, err error) {
	fmt.Println("登录")
	return true, nil

}

//获取前置端
func (d DbServiceImpl) GetFronts() (r []*FrontStationInfo, err error) {
	frontStations := make([]FrontStation, 0)
	err = dbContext.FrontStations.Find(nil).All(&frontStations)
	fmt.Println(&frontStations)
	result := make([]*FrontStationInfo, 0)
	for i := 0; i < len(frontStations); i++ {
		frontStationInfo := FrontStationInfo{frontStations[i].Latitude, frontStations[i].Longitude, frontStations[i].Id.Hex()}
		result = append(result, &frontStationInfo)
	}

	if err == nil && len(result) > 0 {
		return result, nil
	}
	return nil, nil
}

//添加告警事件数据
func (d DbServiceImpl) InsertEventAlarm(r EventAlarm) (err error) {
	eventAlarm := EventAlarm{}
	err = dbContext.EventAlarms.Insert(eventAlarm)
	return err
}

//添加告警事件
func (d DbServiceImpl) InsertFrontEvent(r FrontAlarm) (err error) {
	frontAlarm := FrontAlarm{}
	err = dbContext.FrontAlarms.Insert(frontAlarm)
	return err
}
