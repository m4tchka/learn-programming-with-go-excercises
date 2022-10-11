# Using bash commands:
# Make a new directory called “documents”
# Create three files in it - name.txt, age.txt, address.txt
# In each of them, write some data about you, based on the doc’s name
# name.txt - write your full name
# age.txt - write your age
# address.txt - write your home address
# List the contents of the new directory & verify the three files exist
# Example output:
# address.txt     age.txt         name.txt

# Output the contents of the files on the terminal:
# Rocky Balboa
# 99
# the boxing ring

# After you’re done, remove the directory with all the files inside


mkdir documents
cd documents; touch name.txt age.txt address.txt 
echo "name: m4tchka" > name.txt
echo "age: 22" > age.txt
echo "address: London, United Kingdom" > address.txt
ls
cat name.txt age.txt address.txt
cd ..
rm -rf documents/