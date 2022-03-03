package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)
type MSG struct {
	FILE 	[]byte `json:"FILE"`
	NAME	string `json:"NAME"`
	CHANNEL string `json:"CHANNEL"`
}

func sendMode(){

	app := &cli.App{}
	app.UseShortOptionHandling =true
	app.Commands  = []*cli.Command{
		{
			
			Name:  "send",
			Usage: "send a file mode",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "file", Aliases:[]string{"s"}},
				&cli.StringFlag{Name: "channel", Aliases:[]string{"ch"}},
			},
			
			Action: func (c *cli.Context) error  {
				handleSend(c.String("file"), c.String("channel"))
				return nil
			},
		},
	}
	  
	
	err := app.Run(os.Args)
	if err != nil {
	log.Fatal(err)
	}
}

func handleSend(arh string, channel string) {
	/// preparar json
	path := "./docs/" + arh
	file,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("sorry, the file isn't open: ", err)
	}
	// fmt.Println("ARCHIVO",string(file))
	data := MSG{file, arh, channel}
	// data, err := json.Marshal(msg)
	if err != nil{
		log.Fatal("sorry, the data isn't convert to json", err)
	}

	/// 
	// fmt.Print(string(data))

	/// enviar json al server
	var input string
	go clientToServer(data)
	fmt.Scanln(&input)
}

func clientToServer(data MSG){
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		fmt.Println("Error al intentar conectarse al servidor ", err)
		return
	}

	err = json.NewEncoder(conn).Encode(data)
	if err != nil {
		fmt.Println("Error al codificar el msg ", err)
	}
	defer conn.Close();
}