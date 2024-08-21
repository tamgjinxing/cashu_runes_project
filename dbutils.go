package main

import (
	"ecash_runes_project/storage"
	"fmt"
	"log"
	"reflect"
)

func QueryDataList(sql string, params []interface{}, model interface{}) (interface{}, error) {
	manager := storage.GetDBManager()
	if manager == nil {
		log.Println("Database manager is not initialized.")
		return nil, nil
	}

	fmt.Println(sql)

	rows, err := manager.DB.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sliceType := reflect.SliceOf(reflect.TypeOf(model))
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		elem := reflect.New(reflect.TypeOf(model)).Elem()
		addresses := make([]interface{}, len(cols))

		for i := range cols {
			addresses[i] = elem.Field(i).Addr().Interface()
		}

		if err := rows.Scan(addresses...); err != nil {
			return nil, err
		}

		sliceValue = reflect.Append(sliceValue, elem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sliceValue.Interface(), nil
}

func SaveData(sql string, params []interface{}) (bool, error) {
	manager := storage.GetDBManager()
	if manager == nil {
		log.Println("Database manager is not initialized.")
		return false, nil
	}

	result, err := manager.DB.Exec(sql, params...)

	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to retrieve rows affected: %v", err)
		return false, err
	}

	// 判断是否确实有行被插入
	if rowsAffected == 0 {
		return false, nil // 没有行被影响，可能是因为插入的数据与现有数据重复等原因
	}

	return true, nil // 插入成功
}

func CheckExist(sql string, params []interface{}) (bool, error) {
	manager := storage.GetDBManager()
	if manager == nil || manager.DB == nil {
		log.Println("Database manager is not initialized.")
		return false, nil
	}

	var exists bool
	rows := manager.DB.QueryRow(sql, params...)
	err := rows.Scan(&exists)

	if err != nil {
		return exists, err
	}

	// 检查记录是否存在
	if exists {
		fmt.Println("记录存在")
	} else {
		fmt.Println("记录不存在")
	}

	return exists, nil
}
