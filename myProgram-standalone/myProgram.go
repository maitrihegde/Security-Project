package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
        "ethos/efmt"
        "math"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	if status != syscall.StatusOk {
		efmt.Fprint(syscall.Stderr,"Error opening %v: %v\n", path, status)
        syscall.Exit(syscall.StatusOk)
	}


//	data    := MyType { "foo", "bar" }
//        fd, status := ethos.OpenDirectoryPath(path)
//	data.Write(fd)
//	data.WriteVar(path +"foobar")
//        efmt.Fprint(syscall.Stderr,"this will print in the log")
//        efmt.Println(data)
//        efmt.Println(fd)

        // I'm adding code for Box here
        mybox := MyType{}
        mybox.w = 0.0
        mybox.x = 1.1
        mybox.y = 3.3
        mybox.z = 4.4
        fd,status := ethos.OpenDirectoryPath(path)
        mybox.Write(fd)
        mybox.WriteVar(path)
        efmt.Fprint(syscall.Stderr,"this will print in the log")
        efmt.Println(mybox)
     //   efmt.Println(fd)
        a := (1.1-0.0)*(1.1-0.0)
        b := (4.4-3.3)*(4.4-3.3)
        efmt.Println("The distance between the two coordinates are",math.Sqrt(a+b))
        efmt.Println("You have successed completed the program")
  
}
