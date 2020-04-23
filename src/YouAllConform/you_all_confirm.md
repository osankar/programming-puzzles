#Puzzle:
There is a queue of people standing to enter an event. All the people are wearing caps. Some of therm are wearing the cap with visor/bill facing forward and the others are wearing the with visor/bill facing backwards. 
Every person in the queue holds a token that tells their position in the queue. The gatekeeper want to let the people in after everyone in the queue wears the cap in the same direction; either visor backwards or forwards.
You want to generate a minimum set of commands asking the people in the queue so that all the people conform to the same way of wearing the cap. 

The commands can be like,
* People with token [1-3] flip your cap
* People with [7-7] flip your cap.

For example, if the cap visor direction of the people in queue are F, B, B, B, F,F,F,B,F,B,F,F where F denotes forwards and B denoted backwards.
Assuming the token numbers start at 0, the solution here to flip [1-3], [7-7], [9-9]

#Solution Approach
The brute force approach is to find out how many ranges are there for the forwards and backwards by iterating through the queue. Whichever has rge less number of ranges, use those ranges to give flip cap command. 

In the given example, F, B, B, B, F,F,F,B,F,B,F,F the ranges for forwards are [0-0], [4-6], [8-8], [10-11] and backwards are [1-3], [7-7], [9-9] 
. This solution requires a pass through the entire queue once and then iterating through the ranges to give commands. Not a bad solution, time complexity wise O(N) for pass through the queue and O(N/2) in worst case for iteration through the ranges. That makes the time complexity O(N+N/2).
O(N) space complexity for the worst case. The worst case for this problem is alternate visor directions like F,B,F,B,F or B,F,B.

Lets look at a few more example and see if we can find any patterns that can help us to find a better solution. 
* B, F -> the answer here is either either flip [0,0] or [1,1]
* B, F, B -> the answer here is [1,1]
* F, F, B, F, B, B -> the answer here is [2,2], [4,5]
* B, F, B, F, F, F, B, F -> the answer could either be [0,0], [2,2], [6,6] or [1,1], [3,5], [7,7]

A close observation to the data sets and the answers, the ranges with the visor direction that appears second are the ranges you want to flip the caps. 
Algorithmically, 
* Use the first direction [0] as the pivot
* Iterate over the queue from 1 to size-1
    * if the direction of the current visor is same as the previous visor, continue
    * if the direction of the visor is not same as the visor, start the range
    * else give the command to flip from start of the range to previous index
    
