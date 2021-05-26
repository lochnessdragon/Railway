package main

import (
  "io/ioutil"
  "log"
)

func get_file_list(dir string) []string {
  var files []string 

  var file_list, err = ioutil.ReadDir(dir)
  if err != nil {
    log.Fatal(err.Error())
  }

  for x := 0; x < len(file_list); x++ {
    //log.Printf("Found file: %s", file_list[x].Name())
    if file_list[x].IsDir() {
      files = append(files, get_file_list(dir + file_list[x].Name() + "/")...)
    } else {
      files = append(files, dir + file_list[x].Name())
    }
  }

  return files
}