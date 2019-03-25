package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

const tab string = "    "
const closeFor = "}  //For"
const closeIf = "}  //If"
const closemain = "}  //func main"

type Checkpoint struct {
    Weight float32
    label  int
}

/**
 * @brief      Change a character of a string
 *
 * @param      str        The string
 * @param      char       The char
 * @param      indexChar  The position of the char to change
 *
 * @return     { The string with the character changed }
 */
func ChangeChar(str string, letter string, indexChar int) string {
    return str[:indexChar] + letter + str[indexChar+1:]
}

func MakeTabs(i int, file io.Writer, err error) int {
    var WhatIWrite int
    for j := 0; j < i; j++ {
        WhatIWrite, err = io.WriteString(file, tab) //tab
    }
    return WhatIWrite
}

func CheckError(err error) {
    if err != nil {
        fmt.Print(err)
    }
}

func main() { //BestPath(this Graph, Base int) {
    filename := "SolveTravel.go"
    file, err := os.Create(filename)
    var Index = []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75}
    ByteHeadImport, err := ioutil.ReadFile("CodeBricks/HeadImport.txt") // just pass the file name
    CheckError(err)
    ByteOpenFor, err := ioutil.ReadFile("CodeBricks/OpenFor.txt") // just pass the file name
    CheckError(err)
    ByteOpenIf, err := ioutil.ReadFile("CodeBricks/OpenIf.txt") // just pass the file name
    CheckError(err)
    ByteIfStatement, err := ioutil.ReadFile("CodeBricks/IfStatement.txt") // just pass the file name
    CheckError(err)

    //Path := [5040]Checkpoint{}
    var count int = 0

    HeadImport := string(ByteHeadImport)
    OpenFor := string(ByteOpenFor)
    OpenIf := string(ByteOpenIf)
    IfStatement := string(ByteIfStatement)

    // Path[0].label = Base
    //Path[(this.size - 1)].label = Base
    WhatIWrite, err := io.WriteString(file, HeadImport) //HeadImport
    if err != nil {
        fmt.Println(WhatIWrite, err)
    }
    fmt.Println(WhatIWrite, err)
    //WhatIWrite, err = io.WriteString(file, tab)     //tab
    WhatIWrite, err = io.WriteString(file, OpenFor) //For
    if err != nil {
        fmt.Println(WhatIWrite, err)
    }
    var aux int = 0
    index := string("i")
    for i := 1; i < 8; i++ {
        WhatIWrite = MakeTabs(i, file, err)

        //var letter byte = Index[i]
        letter := string([]byte{Index[i]})
        OpenFor = ChangeChar(OpenFor, letter, 4)
        OpenFor = ChangeChar(OpenFor, letter, 12)
        OpenFor = ChangeChar(OpenFor, letter, 27)

        OpenIf = ChangeChar(OpenIf, index, 8)
        OpenIf = ChangeChar(OpenIf, letter, 3)

        WhatIWrite, err = io.WriteString(file, "\r\n")
        WhatIWrite = MakeTabs((i + 1), file, err)
        WhatIWrite, err = io.WriteString(file, OpenFor) //For

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, OpenIf) //If

        WhatIWrite, err = io.WriteString(file, tab) //tab
        IfStatement = ChangeChar(IfStatement, letter, 0)

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, IfStatement) //IfStatement

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, closeIf) // "}"
        aux = i
    }
    WhatIWrite, err = io.WriteString(file, "\r\n")
    WhatIWrite = MakeTabs((aux + 2), file, err)
    WhatIWrite, err = io.WriteString(file, "//AHHHH QUE LINDU CARAAAAA\n") //SHOULD BE CODE
    for i := 0; i < 8; i++ {
        WhatIWrite = MakeTabs((8 - i), file, err)
        WhatIWrite, err = io.WriteString(file, closeFor) // "}"
        WhatIWrite, err = io.WriteString(file, "\r\n")

    }
    WhatIWrite, err = io.WriteString(file, closemain)
    fmt.Println(WhatIWrite, err)
    file.Close()
    fmt.Printf(" I found an amount of %d paths \n ", count)
}
