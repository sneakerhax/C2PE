# Python Stream C2 (tcp)

I used this TCP echo client/server using streams [example](https://docs.python.org/3/library/asyncio-stream.html#examples) and built a very basic C2

## Starting the server

Must have Python 3.7+

```
python3 server.py
** C2 Serving on ('0.0.0.0', 8888) **
```

Starting the server

## Running the client

Must have Python 3.7+. Make sure to replace the server_ip variable on the client before running.

```
python client.py
```

Running the client with python (no build necessary)

## Building the client with Pyinstaller

```
pyinstaller.exe --onefile --noconsole .\client.py
```

Builds the client into a single file executable using [pyinstaller](https://www.pyinstaller.org/)

## Usage

```powershell
[*] Received callback from '192.168.1.6'
[+] Command to run?
192.168.1.6> powershell.exe $PSVersionTable.PSVersion
[*] Command Sent: powershell.exe $PSVersionTable.PSVersion
[+] Results from 10.0.0.68
Major  Minor  Build  Revision
-----  -----  -----  --------
5      1      19041  610
```

Running a simple powershell command

## References:
* [PyInstaller](https://pyinstaller.org/en/stable/)