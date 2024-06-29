# InactiveSend

Currently, many open source projects provide some limited automation capabilities. These limitations are that they can only send mouse and keyboard events to activated windows, so I have recently been researching how to implement this capability for inactive windows.

This project provides two language demos, which respectively show how to call the dll provided by autoit3 through Python and Golang to implement the ability to send mouse and keyboard events to inactive windows(Only Support Windows OS).

If you want to run this demo locally: 
1. You need to create a plain text file locally, name it "test.txt"
2. Open it with the editor that comes with Windows, then make it inactive
3. Use Python or Go to run the file in this project.  
  


## demo display
This project demonstrates the capabilities of two Autolt APIs, corresponding to mouse clicks on inactive windows and sending keyboard events to inactive windows.

with golang, run `go run main.go`

![image](https://github.com/Knowckx/InactiveSend/blob/main/pic/gif%20go-demo.gif)


with python, run `python pydemo\run.py`

![image](https://github.com/Knowckx/InactiveSend/blob/main/pic/gif%20py-demo.gif)



According to my tests, this capability is also effective for some game windows.


You don't need to download autoit3 to use this project. I have copied the corresponding dll file. You may need to change the path of the dll file in your local computer.

# About autoit3
The autoit3 dll also provides many other APIs. If you have more complex needs, you can study them further. 

Thanks to the autoit3 project, it is very powerful.

https://www.autoitscript.com/site/autoit/