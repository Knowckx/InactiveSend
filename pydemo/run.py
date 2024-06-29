import ctypes
""" For more information about parameters, please refer to this Link 
    https://www.autoitscript.com/autoit3/docs/functions/ControlSend.htm
"""

autoItXDLL = ctypes.WinDLL('dll/AutoItX3_x64.dll')
autoItXDLL.AU3_ControlSend.argtypes = [ctypes.c_wchar_p, ctypes.c_wchar_p, ctypes.c_wchar_p, ctypes.c_wchar_p, ctypes.c_int]
autoItXDLL.AU3_ControlSend.restype = ctypes.c_int
appTitle = "test - Notepad"
content = "Hello InactiveSend"
autoItXDLL.AU3_ControlSend(appTitle, "", "Edit1", content, 0)
