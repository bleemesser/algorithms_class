# Data structures
> Option 3

I implemented an AVL tree (with guidance on how they work & how to do rotations from ChatGPT).
I also implemented a standard unbalanced binary tree.

Timing results:
Binary Tree FindMin n=10000 94ns\
Binary Tree FindMin n=20000 130ns\
Binary Tree FindMin n=40000 148ns\
Binary Tree FindMin n=80000 194ns\
Binary Tree Insertion n=1000 157.139µs\
Binary Tree Insertion n=2000 243.015µs\
Binary Tree Insertion n=4000 520.117µs\
Binary Tree Insertion n=8000 1.127149ms

These values were calculated from 10000 iterations of the function. For FindMin, the number 
represents the number of nodes in the tree. For Insertion, the number represents the number
of nodes to insert.









