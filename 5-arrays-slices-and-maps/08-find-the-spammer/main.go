package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Write a program which reads a set of lines which shows all your email messages, with date as first column, sender as second column & subject as final column.
Find what’s the average emails per day you receive by sender.
The total number of days is based on the range of days which you have given. For example, if you’re given emails from 2021-01-01 to 2021-01-03 only, the total number of days is three.
Finally, print the list of senders which are likely to be spammers - those which send you more than 3 emails per day on average.
Print the results sorted based on the average emails per day in descending order.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	// Gets the first input (the number of emails to be inputted)
	minTime := time.Unix(-2208988800, 0)
	maxTime := minTime.Add(1<<63 - 1)
	minDate, maxDate := maxTime, minTime
	avgEmailFreq := map[string]float64{} // Map of emails (strings) against the average frequency over the time period
	for i := 0; i < num; i++ {           // Loop through all of the emails (which are on separate lines)
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		lineSli := strings.Split(strings.TrimSpace(line), " ") // From each line, create a slice of strings
		email := lineSli[1]                                    // Set each line's email to be the second entry in the slice
		date, err := time.Parse("2006-05-02", lineSli[0])      // Set each line's date to be the parsed date (in the format "YYYY-MM-DD") which is the first entry in the slice
		if err != nil {
			panic(err)
		}
		avgEmailFreq[email]++     // Within the avgEmailFreq slice, increment that email's count
		if date.Before(minDate) { // If the date of an email is earlier than the min date, move the mindate back to that date
			minDate = date
		}
		if date.After(maxDate) { // If the date of an email is later than the max date, move the max date back to that date
			maxDate = date
		}

	}
	totalDays := maxDate.Add(24*time.Hour).Sub(minDate).Hours() / 24 //Find the total number of days between the newest and oldest emails
	var emails []string                                              // A slice of strings of each unique email from the input, used to sort the addresses by their average frequency
	for email := range avgEmailFreq {                                // Loop through each email in the average frequency map
		avgEmailFreq[email] /= totalDays // Set the values to be the average emails per day
		emails = append(emails, email)   // Append each email to the slice
	}
	sort.Slice(emails, func(i, j int) bool { // Sort the slice of emails by their average frequency using what is basically a callbacks function
		return avgEmailFreq[emails[i]] > avgEmailFreq[emails[j]]
	})
	var spammers []string          // Define a slice of emails that will be classed as spammers
	for _, email := range emails { // Loop through the now sorted slice of emails
		fmt.Printf("%s -> %.2f\n", email, avgEmailFreq[email]) // Print each email (in descending order of average frequency, and that email's corresponding frequency in the map)
		if avgEmailFreq[email] >= 3 {                          // If that email has at least 3 emails per day, append that email to the spammers slice
			spammers = append(spammers, email)
		}
	}
	fmt.Println("\nThe spammers are: ")
	for _, spEmail := range spammers { // Loop through the spammers slice and print each
		fmt.Println(spEmail)
	}
}
