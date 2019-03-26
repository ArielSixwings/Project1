package mygraph

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strconv"
)

//___________Code Bricks const________//
//                                    //
const tab string = "    "             //
const closeFor = "}  //for"           //
const closesIf = "}  //if"            //
const closemainfunc = "}  //funcmain" //
//                                    //
//____________________________________//

var Index = []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
    75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
    85, 86, 87, 88, 89, 90, 97, 98, 99, 100,
    101, 102, 103, 104, 105, 106, 106, 108, 109, 110,
    111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
    121, 122}

/**
 * @brief      Class for checkpoint, an "stop" in the path.
 */
type Checkpoint struct {
    Weight      []float32
    Label       []int
    TotalWeight float32
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

/**
 * @brief      Makes tabs.
 *
 * @param      i     { Number of tabs }
 * @param      file  The file
 * @param      err   The error
 *
 * @return     { The new string }
 */
func MakeTabs(i int, file io.Writer, err error) int {
    var WhatIWrite int
    for j := 0; j < i; j++ {
        WhatIWrite, err = io.WriteString(file, tab) //tab
    }
    return WhatIWrite
}

/**
 * @brief      { Check is the file was correctly readed }
 *
 * @param      err   The error
 *
 * @return     { void }
 */
func CheckError(err error) {
    if err != nil {
        fmt.Print(err)
    }
}

/**
 * @brief      Builds paths.
 *
 * @param      size  The size
 *
 * @return     The paths.
 */
func BuildPaths(size int) {
    numberfile := strconv.Itoa(size)
    filename := "SolveTravel .go"
    filename = (filename[:11] + numberfile + filename[11+1:])
    fmt.Println("the file name", filename)
    file, err := os.Create(filename)

    /*Reads all the Code Bricks*/
    ByteHeadImport, err := ioutil.ReadFile("graph/CodeBricks/HeadImport.txt") // just pass the file name
    CheckError(err)
    ByteOpenFor, err := ioutil.ReadFile("graph/CodeBricks/OpenFor.txt")
    CheckError(err)
    ByteOpenIf, err := ioutil.ReadFile("graph/CodeBricks/OpenIf.txt")
    CheckError(err)
    ByteIfStatement, err := ioutil.ReadFile("graph/CodeBricks/IfStatement.txt")
    CheckError(err)
    ByteStatement, err := ioutil.ReadFile("graph/CodeBricks/Statement.txt")
    CheckError(err)
    ByteReturn, err := ioutil.ReadFile("graph/CodeBricks/Return.txt")
    CheckError(err)

    /*Convert all byte type Code Bricks to string type Code Bricks*/
    HeadImport := string(ByteHeadImport)
    OpenFor := string(ByteOpenFor)
    OpenIf := string(ByteOpenIf)
    IfStatement := string(ByteIfStatement)
    Statement := string(ByteStatement)
    Return := string(ByteReturn)

    /*Set the begning of the file*/
    WhatIWrite, err := io.WriteString(file, HeadImport) //HeadImport
    WhatIWrite, err = io.WriteString(file, OpenFor)     //For
    fmt.Println("okay?", WhatIWrite)
    CheckError(err)

    var aux int = 0
    pastletter := string("i")

    for i := 1; i < size; i++ {
        WhatIWrite = MakeTabs(i, file, err)

        letter := string([]byte{Index[i]})
        /*Change the for Open*/
        OpenFor = ChangeChar(OpenFor, letter, 4)
        OpenFor = ChangeChar(OpenFor, letter, 12)
        OpenFor = ChangeChar(OpenFor, letter, 27)

        /*Change the if Open*/
        OpenIf = ChangeChar(OpenIf, pastletter, 8)
        pastletter = letter
        OpenIf = ChangeChar(OpenIf, letter, 3)

        WhatIWrite, err = io.WriteString(file, "\r\n")
        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, OpenFor) //For

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, OpenIf) //If

        WhatIWrite, err = io.WriteString(file, tab) //tab
        IfStatement = ChangeChar(IfStatement, letter, 0)

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, IfStatement) //IfStatement

        WhatIWrite = MakeTabs((i + 2), file, err)
        WhatIWrite, err = io.WriteString(file, closesIf) // "}"
        aux = i
    } //For

    WhatIWrite, err = io.WriteString(file, "\r\n")
    WhatIWrite = MakeTabs((aux + 2), file, err)
    WhatIWrite, err = io.WriteString(file, Statement) //the code to build the path

    /*Close the fors*/
    for i := 0; i < 8; i++ {
        WhatIWrite = MakeTabs((8 - i), file, err)
        WhatIWrite, err = io.WriteString(file, closeFor) // "}"
        WhatIWrite, err = io.WriteString(file, "\r\n")

    }
    WhatIWrite, err = io.WriteString(file, Return)
    WhatIWrite, err = io.WriteString(file, closemainfunc)
    file.Close()
}
