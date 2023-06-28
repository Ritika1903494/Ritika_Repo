package main
 
import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"strings"
)

func main() {

		/* Created an empty file */
		Emptyfile, err := os.Create("./data.txt")
	
		if err != nil {
			panic(err)
		}
		fmt.Printf("Emptyfile is of type  %T\n",Emptyfile)	
		Write_file(Emptyfile); 
		Read_file(Emptyfile);
		Rename_file(Emptyfile);
		Delete_file(Emptyfile);
}

func Read_file(Emptyfile  *os.File){
   /* reading the file content in string format */
	data, err := ioutil.ReadFile("data.txt")  //read the file 
	if err != nil {
		panic(err)
	}
	inputdata := string(data)
	read_lines := strings.Split(inputdata, "\n")
	for _, line := range read_lines {
		fmt.Println(line) //print the content line by line
	 }

	 Emptyfile.Close();
}

func Write_file(Emptyfile  *os.File){
	/* writing the content in file*/
	text:="hello ritika thakur, The most important package that allows us to manipulate files and directories as entities is the os package. The io package has the io.Reader interface to reads and transfers data from a source into a stream of bytes.The io.Writer interface reads data from a provided stream of bytes and writes it as output to a target resource."
	length,err:= io.WriteString(Emptyfile,text)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of file is:",length)
	Emptyfile.Close();
}

func Creating_dir(){
/* Creating a directory*/
_, error:= os.Stat("data")
if os.IsNotExist(error) {
	errDir := os.MkdirAll("data", 0755)
	if errDir != nil {
		panic(errDir)
	}else{
		fmt.Println("Directory created successfully!")
	}
}
}
func Rename_file(Emptyfile  *os.File){
	
	/*renaming the file*/
	oldName := "data.txt"
	newName := "newdata.txt"
	error_file:= os.Rename(oldName, newName)
	if error_file != nil {
		panic(error_file)
	}else{
		fmt.Println("renamed file successfully !")
	}
	Emptyfile.Close();
}

func Delete_file(Emptyfile  *os.File){
	err := os.Remove(`C:\Users\rithakur\goproject\files\newdata.txt`)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("Deleted file successfully !")
	}
	Emptyfile.Close();
}