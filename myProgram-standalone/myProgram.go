
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

      // I'm adding code for Box here
        mybox := MyType{}
        mybox1 := MyType{}

     // Assigning values for mybox    
        mybox.x1 = 0.0
        mybox.x2 = 1.1
        mybox.y1 = 3.3
        mybox.y2 = 4.4

        fd,status := ethos.OpenDirectoryPath(path)
        mybox.Write(fd)
        mybox.WriteVar(path)
        efmt.Fprint(syscall.Stderr,"\n This will print in the log\n")
        efmt.Println("\n The coordinates for the first instance are  ",mybox)

     // Computing distnace between the two coordinates
        a := (mybox.x2-mybox.x1)*(mybox.x2-mybox.x1)
        b := (mybox.y2-mybox.y1)*(mybox.y2-mybox.y1)
        efmt.Println("The distance between the two coordinates is ",math.Sqrt(a+b))

 
     // Computing Square root and writing into the new instance 'mybox1'
        mybox1.x1 = math.Sqrt(mybox.x1)
        mybox1.y1 = math.Sqrt(mybox.y1)
        mybox1.x2 = math.Sqrt(mybox.x2)
        mybox1.y2 = math.Sqrt(mybox.y2)
        mybox1.Write(fd)
        mybox1.WriteVar(path)
        
        efmt.Println("\n The coordinates for the new instance are ",mybox1)
        efmt.Println("\n You have successed completed the program\n")
  
}
