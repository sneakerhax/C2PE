interval=60

echo "interval is $interval"

while true;
do
	sleep 10
	interval=$(curl 127.0.0.1:8080/sleep.txt 2>/dev/null)
	echo "interval is $interval"
done
	
