# Reddit comment parser
The idea here to simply to have some fun with go and a large data set. There are approximately 54 million records to play with, which you can downoad and find at [kaggle datasets](https://www.kaggle.com/reddit/reddit-comments-may-2015).

First step is connecting to the db, making my query.

I'm planning next on putting the data in elastic search, since doing LIKE searching on a text field is painfully slow*.

*Note that sqlite does have full-text search modules [FTS3,4,5](https://www.sqlite.org/fts5.html).

Ideally, I'd like to use this to play with go channels, testing out how quickly and efficiently I can process and eventually display some metadata about those months comments.

Stuff like:
- What subreddits do users with 'PICKLE' in their username comment the most in? (replace Pickle with term of your choice)
- What time of day are people most likely to curse?
- What sort of comment earns people gold? Does the subreddit they post in have any correlation?

And whatever other weird stuff I can think to find.
