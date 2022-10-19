package utils

import (
  "errors"
  "os"
)

//@description: 文件目录是否存在
//@param: path string
//@return: bool, error
func PathExists(path string) (bool, error) {
  fi, err := os.Stat(path)
  if err == nil {
    if fi.IsDir() {
      return true, nil
    }
    return false, errors.New("存在同名文件")
  }
  if os.IsNotExist(err) {
    return false, nil
  }
  return false, err
}

//@description: 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
  for _, v := range dirs {
    exist, err := PathExists(v)
    if err != nil {
      return err
    }
    if !exist {
      if err := os.MkdirAll(v, os.ModePerm); err != nil {
        return err
      }
    }
  }
  return err
}
