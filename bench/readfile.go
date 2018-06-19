package main

import (
    "os"
    "io"
    "bufio"
    "io/ioutil"
    "time"
    "log"
)

func readCommon(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }


    buf := make([]byte,0)
    for {
        readNum, err := file.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if 0 == readNum {
            break
        }
    }
    defer file.Close()
}

func readBufio(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    bufReader := bufio.NewReader(file)
    buf := make([]byte, 1024)

    for {
        readNum, err := bufReader.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if 0 == readNum {
            break
        }
    }
}

func readIOUtil(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    _, err = ioutil.ReadAll(file)
}

func main() {
    //size is 26MB
    pathName := "benchmarkFile.txt"
    start := time.Now()
    readCommon(pathName)
    timeCommon := time.Now()
    log.Printf("read common cost time %v\n", timeCommon.Sub(start))

    readBufio(pathName)
    timeBufio := time.Now()
    log.Printf("read bufio cost time %v\n", timeBufio.Sub(timeCommon))

    readIOUtil(pathName)
    timeIOUtil := time.Now()
    log.Printf("read ioutil cost time %v\n", timeIOUtil.Sub(timeBufio))
}
