
Issues -

1) What data type should I use for money / amount?
I've left it at float64 initially but:
In Scala I used BigDecimal
https://github.com/golang/go/issues/12127

googling seems big.Rat might be the quickest, but also until standadisation something like this?
https://pkg.go.dev/github.com/luno/luno-go/decimal

At the moment I just print out the gbpAmount using %v


2) No
Can I just * rate?

3) File size / memory usage:

I'm using csvReader.ReadAll for convenience, pulling the whole file into memory then appending to 
var rows []Customer where date in August 2020 and Description is CARD SPEND

For larger files I would have to re-evaluate. Using csvReader.Read looks like the next step.
I've not got into csv reading with golang before this assignment.

4) If we have multiple rows per-customer then we would probably want to calculate "total transfers in August"
We got into discussion about aggregating amounts according to a unique identifier per customer.
Can you guarantee uniqueness of email address per customer / in this data
Else you need to create some kind of composite primary key

5) I didn't want to spend too much time but it is getting time for unit tests