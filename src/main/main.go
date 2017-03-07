package main

import (
    "fmt"
    "flag"
    "header"
    "log"
)

func main(){

    var sStrValue string
    flag.StringVar(&sStrValue, "s"," string values","")
    flag.StringVar(&sStrValue, "string"," string values","")

    var sIntValue int

    flag.IntVar(&sIntValue, "i",0,"")
    flag.IntVar(&sIntValue, "int",0,"")


    var sHelpValue bool
    flag.BoolVar(&sHelpValue,"h",false,"")
    flag.BoolVar(&sHelpValue,"help",false,"")

    settingFlag(flag.CommandLine)

    flag.Parse()

    if sHelpValue {
        showHelp()
        return
    }

    header.GetGsocWSO2Header()

        if sStrValue != "" {
            log.Printf("This is user string input : %s\n",sStrValue)
        }

        if sIntValue != 0 {
            log.Printf("This is integer input : %d\n",sIntValue)
        }
    }

    func showHelp(){
        fmt.Println(`
            Usage of our CLI Template OPTIONS :)

            OUR OPTIONS

            --S, => string   print a string input. ( e=>ult='Any string').
            --i, => int      print a int input. (defaut value => =0).
            --w, => warning  print the warning you en => red
            --h, => help     print for help info.

            Good Luck :D
        `)
    }

    func settingFlag(flag *flag.FlagSet){
        flag.Usage = func(){
            showHelp()
        }
    }
