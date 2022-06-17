package fileparse

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func GetVpcIPaddrGroup(file string, addrGroupName string) (ips []string, id string, err error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Errorln(err)
		return nil, "", err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			logrus.Errorln(err)
			return
		}
	}()

	// 逐行读取Excel文件
	rows, err := f.GetRows(addrGroupName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"sheet": addrGroupName,
		}).Errorf("读取中sheet页异常: %v", err)
		return nil, "", err
	}

	for k, row := range rows {
		// 如果第一列的值不是 ID，则赋值给 erd
		if row[0] != "ID" {
			logrus.WithFields(logrus.Fields{
				"k":   k,
				"row": row,
			}).Debugf("检查每一条需要处理的解析记录")
			// 将每一行中的的每列数据赋值到结构体重
			erd := row[0]

			ips = append(ips, erd)
		} else {
			id = row[1]
		}
	}

	return ips, id, nil
}
