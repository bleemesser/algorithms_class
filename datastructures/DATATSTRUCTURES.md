# Data structures
> Option 3

I implemented an AVL tree (with guidance on how they work & how to do rotations from ChatGPT).
I also implemented a standard unbalanced binary tree.

Timing results:
(Averages of 20,000 loops * 10,000 executions of each function -> 200M timing values per function) \
The input array is randomized for each of the 20,000 loops, and is 10,000 items long\
Insert:  146ns\
AVL Insert:  197ns\
Find Val:  103ns\
AVL Find Val:  86ns\
Find Min:  32ns\
AVL Find Min:  34ns

We can see that the AVL tree is slower at inserting values by a small amount, about the same when finding the minimum, and somewhat faster at finding a value. This is because it self-balances after every insertion, so takes a few extra operations to place each node in the correct spot when given an average (not ideal, not worst-case) input. In contrast, the binary tree has fewer operations per insertion, but—being unabalanced—is harder to binary search to find a value.



