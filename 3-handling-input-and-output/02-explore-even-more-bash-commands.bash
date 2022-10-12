# Using bash commands:
# Create a new directory with a file which contains your top 10 favorite songs
# Write some of the songs multiple times so that there are duplicates
# Using a bash command, output the contents of a file, sorted alphabetically
# Show the contents of the file but remove duplicates using yet another bash command
# Use a command to search the file for your favorite song
# Show the first 5 songs in the file
# Show the last 5 songs in the file
# Make a copy of your file into a different directory (create that one as well if need be) & print its contents to verify the copy succeeded
# Move your file into another directory. The old file location should be empty now

mkdir songs; cd songs; touch favSongs.txt
#echo "xx" > favSongs.txt
nano favSongs.txt 

Ariana Grande - side to side
Rocky - going the distance
Lady Gaga - shallow
Lordi - Hard Rock Hallelujah
Rocky - going the distance
Shocking Blue - Venus
Lady Gaga - shallow
Shocking Blue - Venus
The Vogues - youre the one

sort favSongs.txt | uniq

grep Lordi favSongs.txt
head -n 5 favSongs.txt 
tail -n 5 favSongs.txt

mkdir copy; cp favSongs.txt copy/favSongs_copy.txt
cat copy/favSongs_copy.txt
mkdir songs2; mv favSongs.txt songs2/
cat songs2/favSongs.txt