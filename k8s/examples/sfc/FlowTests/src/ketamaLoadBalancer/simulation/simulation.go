package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strings"

    "strconv"
)

func main() {
    csvFile, _ := os.Open("../tmp/testFlows.xls")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    //var lineFlowID,lineOutPut []string
    ///////////////////////////////////////////////////////////////////////////
    ///               prepare output outputFiles                            ///
    ///////////////////////////////////////////////////////////////////////////
    // Jitter file
    fileJitter, errJitter := os.Create("../sim/jitter.txt")
    checkError("Cannot create file", errJitter)
    defer fileJitter.Close()
    // throughput
    fileThroughput, errThroughtput := os.Create("../sim/throughput.txt")
    checkError("Cannot create file", errThroughtput)
    defer fileThroughput.Close()
    // errorRate
    fileErrorRate, errErrorRate := os.Create("../sim/errorRate.txt")
    checkError("Cannot create file", errErrorRate)
    defer fileErrorRate.Close()

    for {
        lineFlowID, errorFlowID := reader.Read()
        if errorFlowID == io.EOF {
            break
        } else if errorFlowID != nil {
            log.Fatal(errorFlowID)
        }
        lineOutPut, errorOutPut := reader.Read()
        if errorOutPut == io.EOF {
            break
        } else if errorOutPut != nil {
            log.Fatal(errorOutPut)
        }
        amoutOfData, _:=strconv.ParseFloat(lineOutPut[7],64)
        splitTimeInterval:=strings.Split(lineOutPut[6], "-")
        timeIntervalEnd, _:=strconv.ParseFloat(splitTimeInterval[1], 64)
        timeIntervalBegin, _:=strconv.ParseFloat(splitTimeInterval[0], 64)
        timeInterval:= timeIntervalEnd-timeIntervalBegin // - strconv.ParseFloat(splitTimeInterval[0], 32)
        //fmt.Println(string(lineFlowID[0])+" "+fmt.Sprintf("%f", timeInterval))
        //timeInterval, _:=strconv.Atoi(lineOutPut[6])
        throughput:=amoutOfData/timeInterval
        fmt.Println(string(lineFlowID[0])+" "+ string(lineOutPut[9])+" "+string(lineOutPut[12])+" "+fmt.Sprintf("%f",throughput/125000))
        // Write results in outputFiles
        // Jitter
        lineJitter:=make([]string, 2)
        lineJitter[0]=string(lineFlowID[0])
        lineJitter[1]=string(lineOutPut[9])
        writerJitter := csv.NewWriter(fileJitter)
        writerJitter.Comma=' '
        defer writerJitter.Flush()
        errWrite:= writerJitter.Write(lineJitter)
        checkError("Cannot write to file", errWrite)
        // throughput
        lineThroughput:=make([]string, 2)
        lineThroughput[0]=string(lineFlowID[0])
        lineThroughput[1]=fmt.Sprintf("%f",throughput/125000)
        writerThroughtput := csv.NewWriter(fileThroughput)
        writerThroughtput.Comma=' '
        defer writerThroughtput.Flush()
        errWrite= writerThroughtput.Write(lineThroughput)
        checkError("Cannot write to file", errWrite)
        // errorRate fileErrorRate
        lineErrorRate:=make([]string, 2)
        lineErrorRate[0]=string(lineFlowID[0])
        lineErrorRate[1]=string(lineOutPut[12])
        writerErrorRate := csv.NewWriter(fileErrorRate)
        writerErrorRate.Comma=' '
        defer writerErrorRate.Flush()
        errWrite= writerErrorRate.Write(lineErrorRate)
        checkError("Cannot write to file", errWrite)
        // End Writes in outputFiles
    }

}
func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
