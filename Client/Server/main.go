package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	ts "github.com/architagr/taskScheduler"
	"github.com/architagr/workflow/models"
	routerVersion1 "github.com/architagr/workflow/router/Version1"
	"github.com/gin-gonic/gin"
)

var configuration models.Configuration

//var ch = make(chan time.Time)
//var ch = &time.Ticker{}

func main(){
//get Environment for which the api will run
	//this is to help identifu with config file is to be read
	env := flag.String("env", "dev", "environment to be used")
	flag.Parse()

	 fmt.Printf("Envirnment used is : %s\n", *env)

	//get config data according to falg been set
	configuration, err := configuration.Init(*env)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("config file issue for environment %s\n", *env)
		os.Exit(0)
	}

	//start gin server with default middleware
	r := gin.Default()
	//configure routes begin

	//confugure routes for api version1
	routerVersion1.RouterVersion1(r)

	//configure route end

	//run the api server on the mentioned id and port in the config file
	r.Run(configuration.IP + ":" + configuration.Port)

	ts.Init()
	main1()
}
// type Task struct {
// 	t     *time.Timer
// 	time  time.Time
// 	index int
// }

func main1() {
	f, err := os.OpenFile("test.txt", 1, 0644)
	check(err)

	defer func() {
		f.Close()
	}()

	secondsEastOfUTC := int((5 * time.Hour).Seconds()) + int((30 * time.Minute).Seconds())
	india := time.FixedZone("IST Time", secondsEastOfUTC)

	//go WriteFile()
	//timmers := make(chan ts.Task)
	timeTask := make(map[int]time.Time)
	now := time.Now()
	hour := now.Hour()
	minutes := now.Minute()
	second := now.Second()
	date:= now.Day()
	month:= now.Month()
	index := 1

	timeTask[index] = time.Date(2020, month, date, hour, minutes, second, 00, india)
	index = index + 1
	//minutes = minutes + 1
	second = second + 1
	timeTask[index] = time.Date(2020, month, date, hour, minutes, second, 00, india)
	index = index + 1
	//minutes = minutes + 1
	second = second + 1
	timeTask[index] = time.Date(2020, month, date, hour, minutes, second, 00, india)
	fmt.Println("Start")
	fmt.Println(timeTask)
	// go func() {
	// 	for {
	// 		select {
	// 		case x := <-timmers:
	// 			func(task ts.Task) {
	// 				<-task.T.C
	// 				fmt.Println("task : ",task.Data)
	// 				WriteFile(f, task.Data)
	// 				if _, ok := timeTask[task.Index]; ok {
	// 					delete(timeTask, task.Index)
	// 				}
	// 			}(x)
	// 		}
	// 	}
	// }()
	for _, val := range timeTask {
		ts.AddTask(ts.Task{
			T:     time.NewTimer(time.Now().Sub(val)),			
		}, ts.TaskData{
			Data: val,
			CallBack : func(data interface{}, index int){
				fmt.Println("task : ",data)
					WriteFile(f, data)
					ts.DeleteTask(index)
			},
		})
		// timmers <- ts.Task{
		// 	T:     time.NewTimer(time.Now().Sub(val)),
		// 	Data:  val,
		// 	Index: key,
		// }
	}

	defer func() {
		fmt.Println("End")
		fmt.Println(timeTask)
	}()

	time.AfterFunc(10*time.Second, func(){
		timeTask[4] =time.Date(2020, month, date, hour, minutes, second+10, 00, india)
		// timmers <- ts.Task{
		// 	T:     time.NewTimer(time.Now().Sub(time.Date(2020, 10, 18, hour, minutes, second+10, 00, india))),
		// 	Data:  time.Date(2020, 10, 18, hour, minutes, second+10, 00, india),
		// 	Index: 4,
		// }
		newData := time.Date(2020, month, date, hour, minutes, second+10, 00, india)
		ts.AddTask(ts.Task{
			T:     time.NewTimer(time.Now().Sub(newData,)),			
		}, ts.TaskData{
			Data: newData,
			CallBack : func(data interface{}, index int){
				fmt.Println("task : ",data)
					WriteFile(f, data)
					ts.DeleteTask(index)
			},
		})

	})
	// d:= time.NewTimer(15*time.Second)
	// <-d.C
	// 	os.Exit(0);
	//time.Sleep(15 * time.Second)
}

// WriteFile asd
func WriteFile(f *os.File, x interface{}) {

	d2 := fmt.Sprintf("%v \n", x)
	_, err := f.WriteString(d2)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
