import sys

dequeue_index = 0
queue = []

counter = int(sys.stdin.readline().strip())

for _ in range(counter):
    queue_command = sys.stdin.readline().strip()
    
    if queue_command[0] == "1":
        enqueue_command_and_value = queue_command.split(" ")
        queue.append(enqueue_command_and_value[1])

        continue
        
    if queue_command[0] == "2":
        if (dequeue_index + 1) <= len(queue):
            dequeue_index += 1
            
        continue
        
    if queue_command[0] == "3":
        print(queue[dequeue_index])