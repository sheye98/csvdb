package file

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/fs"
	"os"
	"sync"
)

type Data struct {
	Path   string
	Header []string
	Table  [][]string
	rw     *sync.RWMutex
}

func New(path string) *Data {
	return &Data{
		rw:   new(sync.RWMutex),
		Path: path,
	}
}

func (d *Data) ReadHeader() error {
	d.rw.RLock()
	defer d.rw.RUnlock()

	rf, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer rf.Close()

	// 从文件开始，偏移0
	_, err = rf.Seek(0, 0)
	if err != nil {
		return err
	}

	// 读取csv文件
	reader := csv.NewReader(rf)
	column, err := reader.Read()
	if err != nil {
		return err
	}

	d.Header = column
	return nil
}

func (d *Data) ReadTable() error {
	d.rw.RLock()
	defer d.rw.RUnlock()

	rf, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer rf.Close()

	br := bufio.NewReader(rf)
	line, _, err := br.ReadLine()
	if err != nil {
		return err
	}

	_, err = rf.Seek(int64(len(line))+1, 0)
	if err != nil {
		return err
	}

	// 读取csv文件
	cr := csv.NewReader(rf)
	table, err := cr.ReadAll()
	if err != nil {
		return err
	}

	d.Table = table
	return nil
}

func (d *Data) WriteHeader() error {
	d.rw.Lock()
	defer d.rw.Unlock()

	if len(d.Header) == 1 {
		d.Header = append(d.Header, "")
	}

	wf, err := os.OpenFile(d.Path, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}
	defer wf.Close()

	// 偏移到第一行
	_, err = wf.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// 写入列名
	writer := csv.NewWriter(wf)
	defer writer.Flush()

	//TODO 当表头数组的长度为1会出现bug，目测官方包底层bug，没能力修改
	if len(d.Header) == 1 {

	}

	err = writer.Write(d.Header)
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) WriteTable() error {
	d.rw.Lock()
	defer d.rw.Unlock()

	wf, err := os.OpenFile(d.Path, os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(wf)
	line, _, err := reader.ReadLine()
	if err != nil {
		return err
	}

	// 偏移一行
	_, err = wf.Seek(int64(len(line))+1, io.SeekStart)
	if err != nil {
		return err
	}

	// 写入表格
	writer := csv.NewWriter(wf)
	defer writer.Flush()

	for _, t := range d.Table {
		err = writer.Write(t)
		if err != nil {
			return err
		}
	}

	return nil
}
