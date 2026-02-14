def icecreamParlor(m, arr):
    # Write your code here
    
    total_money = m
    
    lookup_table = {}
    
    for index, ice_cream_cost in enumerate(arr):
        if lookup_table.get(total_money - ice_cream_cost, None) is not None:
            
            lesser = min(index, lookup_table.get(total_money - ice_cream_cost))
            
            greater = max(index, index, lookup_table.get(total_money - ice_cream_cost))
            
            return [lesser, greater]      
            
        elif lookup_table.get(ice_cream_cost, None) is None:
            lookup_table[ice_cream_cost] = index

        else: 
            continue