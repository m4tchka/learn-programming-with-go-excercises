# mv songs2/favSongs.txt songs.txt
# mv songs.txt songs2/favSongs.txt
# cat songs.txt | sort | uniq

# Using several bash commands on a single line:
# Show the contents of the file with your favorite songs
# Remove duplicates
# Sort the songs alphabetically
# Save the output from this command to a new file songs_sanitized.txt
# Hint: Use bash pipes and output redirection

cat songs2/favSongs.txt | sort | uniq > songs_sanitized.txt