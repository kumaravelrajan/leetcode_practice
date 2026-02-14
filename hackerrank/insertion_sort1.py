def insertionSort1(n, arr):
    # Write your code here
    
    unsorted_num = arr[-1]
    index = len(arr) - 1
    
    while index >= 0:
        if index - 1 >= 0:
            if arr[index - 1] >= unsorted_num:
                arr[index] = arr[index - 1]
                print(" ".join(str(elem) for elem in arr))

                
            else:
                
                arr[index] = unsorted_num
                print(" ".join(str(elem) for elem in arr))
                
                break

            
                
        index -= 1


        

insertionSort1(3, [2,3,4,5,1])