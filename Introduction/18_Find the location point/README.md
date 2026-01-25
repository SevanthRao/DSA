Q. Find the location point
time limit per test1 second
memory limit per test256 megabytes
You are given the coordinates of a point (X,Y)
 on a Cartesian plane. Your task is to determine the location of the point.

The possible locations are:

Origin — if X=0
 and Y=0
X axis — if Y=0
 and X≠0
Y axis — if X=0
 and Y≠0
1st Quadrant — if X>0
 and Y>0
2nd Quadrant — if X<0
 and Y>0
3rd Quadrant — if X<0
 and Y<0
4th Quadrant — if X>0
 and Y<0
Input
A single line containing two integers X
 and Y
 (−109≤X,Y≤109)
.

Output
Print a single line indicating the location of the point.

Examples
InputCopy
1 10
OutputCopy
1st Quadrant
InputCopy
0 0
OutputCopy
Origin

