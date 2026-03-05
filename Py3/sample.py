class TrieNode:
    def __init__(self): 
        self.children = {} 

class Trie:

    def __init__(self): 
        self.root = TrieNode()

    def search(self, word): 
        currentNode = self.root
        
        for char in word:
            # If the current node has child key with current character:
            if currentNode.children.get(char):
                # Follow the child node:
                currentNode = currentNode.children[char]
            else:
                # If the current character isn't found among
                # the current node's children, our search word 
                # must not be in the trie:
                return None
        return currentNode
    
    def insert(self, word): 
        currentNode = self.root
        for char in word:
            # If the current node has child key with current character:
            if currentNode.children.get(char):
                # Follow the child node:
                currentNode = currentNode.children[char]
            else:
                # If the current character isn't found among
                # the current node's children, we add 
                # the character as a new child node: 
                newNode = TrieNode() 
                currentNode.children[char] = newNode
                # Follow this new node: 
                currentNode = newNode
        # After inserting the entire word into the trie, 
        # we add a * key at the end: 
        currentNode.children["*"] = None

    def collectAllWords(self, node=None, word="", words=[]): 
        # This method accepts three arguments. The first is the
        # node whose descendants we're collecting words from.
        # The second argument, word, begins as an empty string,
        # and we add characters to it as we move through the trie.
        # The third argument, words, begins as an empty array,
        # and by the end of the function will contain all the words 
        # from the trie.

        # The current node is the node passed in as the first parameter, 
        # or the root node if none is provided:
        currentNode = node or self.root
        # We iterate through all the current node's children: 
        for key, childNode in currentNode.children.items():
            # If the current key is *, it means we hit the end of a 
            # complete word, so we can add it to our words array: 
            if key == "*":
                words.append(word)
            else: # If we're still in the middle of a word:
                # We recursively call this function on the child node. 
                self.collectAllWords(childNode, word + key, words)
        return words
    
    def autocomplete(self, prefix): 
        currentNode = self.search(prefix)
        if not currentNode:
            return None
        return self.collectAllWords(currentNode)
    
    ###################### Addition by Kumar ###########################
    
    def print_all_keys(self, current_node=None):
        # Initialize to root only if this is the very first call
        if current_node is None:
            current_node = self.root

        for key, node in current_node.children.items():
            print(key)
            
            # Only recurse if the node exists (is not None).
            # This prevents us from following the "*" key, which points to None.
            if node: 
                self.print_all_keys(current_node=node)

    def autocorrect(self, user_str = None):
        if not user_str :
            return
        
        current_node = self.root
        word = ""
        is_autocorrect_used = False
        
        for user_char in user_str:
            
            if current_node.children.get(user_char): 
                # current_node.children has user_char key. Proceed as usual

                current_node = current_node.children[user_char]
                word += user_char
            
            else:
                #current_node.children does not have user_char key. Start autocorrect procedure

                is_autocorrect_used = True

                return word + self.collectAllWords(current_node)[0]

        if(is_autocorrect_used):
            return f"autocorrect applied for {user_str}: {word}"
        else:
            return f"autocorrect NOT applied for {user_str}"


trie = Trie()

trie.insert("cat")
trie.insert("catnap")
trie.insert("catnip")

print(trie.autocorrect("catnar"))
print(trie.autocorrect("ca"))
print(trie.autocorrect("caxasfdij"))