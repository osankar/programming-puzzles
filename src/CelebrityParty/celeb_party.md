# Puzzle: Best time to party #
There is a party to celebrate celebrities that you get to attend because you won a ticket at your office lottery. Because of the high demand for tickets you only get to stay for one hour but you get to pick which one since you received a special ticket. You have access to a schedule that lists when exactly each celebrity is going to attend the party. You want to get as many pictures with celebrities as possible to improve your social standing. This means you wish to go for the hour when you get to hob-nob with the maximum number of celebrities and get selfies with each of them.
We are given a list of intervals that correspond to when each celebrity comes and goes. Assume that these intervals are [i, j), where  i and j correspond to hours, where  and  correspond to hours. That is, the interval is closed on the left hand side and open on the right hand side. This just means that the celebrity will be partying on and through the th hour, but will have left when the th hour begins. So even if you arrive on dot on the th hour, you will miss this particular celebrity.

Here’s an example: 

Celebrity     |Enters          | Leaves  |
------------- | ------------- |-------|
Beyoncé       | 6  |7|
Taylor  | 7   | 9|
Brad|10|11|
Katy|10|12|
Tom|8|10|
Drake|9|11|
Alicia|6|8|

When is the best time to attend the party? That is, which hour should you go to?

# Solution: Approach #
At every hour (or at an instant of time), a few celebrities come in and a few others mey leave. Each celebrity stays in the party for an arbitrary number of hours as allowed by their own schedules. 
As party goer you are allowed to stay for an hour at the maximum. You wanna pick an optimal time so that you can meet maximum number of celebrities. Assume that you have no favorite bias. 

The best way to think about the solution for this puzzle is to put yourself at the party as a celebrity gate keeper and if you keep maintaining the record of celebs coming in and going out, you have the record of 
how many celebrities are in the party at any given point of time. After the party if you look at your records, you can see what's the best time to meet maximum number of celebrities. 
But in reality, we can not wait for the party to complete  because it kills the very purpose of meeting the celebrities. 

If your algorithm can emulate the celebrity gatekeeper record before the party based on schedules, you will be able to determine the best time. How do we you so? With respect to the gatekeeper the events are happening in chronological order whereas the 
schedules given to us are not in the chronological order. And the entrance and exit of a particular celebrity happen at distinct points of times. So, if you treat the entry and exif of a celebrity as separate movements and 
chronologically sort all the movements and increase the celebrity count by 1 for an entrance movement and decrease the celebrity count by 1 for an exit movement you will have the celebreties in the party any given point of time.
One caveat here is that if one celebrity comes and another celebrity leaves at the same time, should you decrement first or increment first? You shall complete the decrements first and then the increments. Let us say you have 10 celebs 
at 8:59 and 7 of them are leaving and 2 more new celebrities are coming in at 9:00. If we count the increments first at 9:00 the count reaches 12 and after considering the decrements at 9:00 the actual count of celebs is 5. When looking for maximum you may see 12 celebs at 9:00
as maximum but in reality there are only 7 celebrities. By the description of teh puzzle, you shall not count the celebs leaving at time T as the celebs present.

Now, the algorithm to solve this puzzle is 
* Split the each schedule into two movements entrance and exit with their corresponding times
* Sort the movements by time to form the chronological order of movements
  * If the an entrance and exit event happens at the same time, make sure exit comes before the entrance in sorted movememts
* set maxCelebs to 0, optimalTime to 0, celebCount to 0
* Iterate over the chronological events
  * if movement is an entrance
    * Increment the celeb count
    * if celebCount > maxCelebs, set the maxCelebs to celebCount and set the movement's time to optimal time
  * if the movement is an exit
    * decrement the celebCount
