def check_if_100_within_range(start_position):
    if 100 - start_position <= 6:
        return True
    else:
        return False

def quickestWayUp(ladders, snakes):
    # Write your code here
    
    # Idea:
    # Focus on the ladders since they offer the biggest jump in moving upwards.
    # Find the diff b/w start and end of each ladder row and the difference between end of one ladder and start of another ladder.
    
    # Meanwhile while jumping from ladder to ladder also consider starting indices of snakes. Make, use hash table to store these indices for fast lookup?
    
    ladders_lookup_table = {}
    snakes_lookup_table = {}
    number_of_steps = 0
    current_player_position = 1
    
    for row in ladders:
        ladders_lookup_table[row[0]] = {row[1]}
        
    for row in snakes:
        snakes_lookup_table[row[0]] = row[1]
        
        
    sorted_ladders_lookup_table = {k: ladders_lookup_table[k] for k in sorted(ladders_lookup_table)}
    
    sorted_snakes_lookup_table = {k: snakes_lookup_table[k] for k in sorted(snakes_lookup_table)}
    
    for ladder_starting_position, ladder_ending_position in sorted_ladders_lookup_table.items():
        if ladder_starting_position <= (current_player_position + 6):
            number_of_steps += 1
            current_player_position = ladder_ending_position
            
            if check_if_100_within_range(current_player_position) :
                # player is in range to reach 100 with a single dice role. 
                
                number_of_steps += 1
                return number_of_steps
                
        else: 
            # Next ladder farther than current_position + 6. Hence, check if snakes occur b/w current_position and next ladder starting.
            
            snakes_in_between_player_position_and_next_ladder = []
            
            for snake_starting_position in sorted_snakes_lookup_table:
                if current_player_position <= snake_starting_position <= ladder_starting_position:
                    snakes_in_between_player_position_and_next_ladder.append(snake_starting_position)
                    
            # Avoid snake starting positions while trying to reach next ladder.
            
            for snake_starting_position in snakes_in_between_player_position_and_next_ladder:
                if snake_starting_position - current_player_position > 6:
                    current_player_position += 6
                    number_of_steps += 1
                    
                    if check_if_100_within_range(current_player_position):
                        return number_of_steps + 1
                    
                elif (current_player_position + 6) == snake_starting_position:
                    current_player_position += 5
                    number_of_steps += 1
                    
                    if check_if_100_within_range(current_player_position):
                        return number_of_steps + 1
                    
                elif (snake_starting_position - current_player_position) < 6:
                    list_snake_starting_positions = list(sorted_snakes_lookup_table.keys())
                    
                    index_of_current_snake_starting_position = list_snake_starting_positions.index(snake_starting_position)
                    
                    if ((list_snake_starting_positions[index_of_current_snake_starting_position + 1] - current_player_position) > 6) or ((list_snake_starting_positions[index_of_current_snake_starting_position + 1] - current_player_position) < 6):
                        current_player_position += 6
                        number_of_steps += 1

                        if check_if_100_within_range(current_player_position):
                            return number_of_steps + 1
                        
                    elif (list_snake_starting_positions[index_of_current_snake_starting_position + 1] - current_player_position) == 6:
                        
