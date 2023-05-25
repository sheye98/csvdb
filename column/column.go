package column

import (
	"fmt"
	"gycsv/file"
)

type Column struct {
	Path string
}

func New(path string) *Column {
	return &Column{
		Path: path,
	}
}

func (c *Column) Add(field string) error {
	rwf := file.New(c.Path)

	// 读取表头
	err := rwf.ReadHeader()
	if err != nil {
		return err
	}

	// 读取表格
	err = rwf.ReadTable()
	if err != nil {
		return err
	}

	// 添加新的列名
	rwf.Header = append(rwf.Header, field)

	// 写入新表头
	err = rwf.WriteHeader()
	if err != nil {
		return err
	}

	// 写入表格
	err = rwf.WriteTable()
	if err != nil {
		return err
	}

	return nil
}

func (c *Column) Alter(before string, after string) error {
	rwf := file.New(c.Path)

	// 读取表头
	err := rwf.ReadHeader()
	if err != nil {
		return err
	}

	fmt.Println(rwf.Header)

	// 读取表格
	err = rwf.ReadTable()
	if err != nil {
		return err
	}

	for i := 0; i < len(rwf.Header); i++ {
		if rwf.Header[i] == before {
			rwf.Header[i] = after
		}
	}

	// 写入新表头
	err = rwf.WriteHeader()
	if err != nil {
		return err
	}

	// 写入表格
	err = rwf.WriteTable()
	if err != nil {
		return err
	}

	return nil
}

// TODO 删除列
