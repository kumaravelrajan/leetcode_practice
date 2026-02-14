def insertionSort2(n, arr):
    # Write your code here
    
    for index in range(1, len(arr)):
        
        possibly_unsorted_elem = arr[index]
        i = index
        
        while i > 0 and arr[i - 1] > possibly_unsorted_elem:
            arr[i] = arr[i - 1]
            i -= 1
            
        arr[i] = possibly_unsorted_elem
        
        print((" ").join(str(xyz) for xyz in arr))

insertionSort2(6, [1, 4, 3, 5, 6, 2])