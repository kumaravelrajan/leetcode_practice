from collections import deque, defaultdict

def bfs(n, m, edges, s):
    # Write your code here
    
    visited = [False]*n
    queue = deque()
    res = [-1] * (n-1)    
    index_in_res_for_node_lookup = {}
    is_s_encountered = False
    edge_hash_table = defaultdict(list)
    distance_of_x_from_s = {}


    # Setting index in res for nodes for easier lookup
    for x in range(1, n+1,1):
        if x == s:
            is_s_encountered = True
            
        if not is_s_encountered:
            index_in_res_for_node_lookup[x] = x - 1
            
        else:
            index_in_res_for_node_lookup[x] = x - 2
    
    queue.append(s)
    
    # Setting up adjacency matrix for easier lookup
    for row in edges:
        edge_hash_table[row[0]].append(row[1])
        edge_hash_table[row[1]].append(row[0])    
    

    while len(queue) > 0:
        
        current_node = queue.popleft()
        
        visited[current_node - 1] = True
        
        for edge_ending in edge_hash_table[current_node]:
            if not visited[edge_ending - 1]:

                if current_node == s:
                    distance_of_x_from_s[edge_ending] = 6
                else:
                    distance_of_x_from_s[edge_ending] = distance_of_x_from_s[current_node] + 6

                res[index_in_res_for_node_lookup[edge_ending]] = distance_of_x_from_s[edge_ending]
                queue.append(edge_ending)
                
    return ' '.join(str(s) for s in res)

print(bfs(4, 2, [[1, 2], [1, 3]], 1))
print(bfs(3, 1, [[2, 3]], 2))