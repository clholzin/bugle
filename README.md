![mascot](https://github.com/clholzin/bugle/blob/master/assets/gopher.png)

### Bugle - a searchable and refreshable netstat implementation

Simply call bugle with integer second interval and search details and hear the network Bugle

<b>bugle -an 2 80-ESTABLISHED *.80-LISTEN 2000-ESTABLISHED</b>


![howto](https://github.com/clholzin/bugle/blob/master/assets/record.gif)



### How to Search:

##### Searching OR operations are seperated by spaces:

example: ```80-ESTABLISHED *.80-LISTEN```

This will search for connections made on 80** <b>OR</b> 80** listeners


##### Searching AND operations are seperated by hyphen:

By adding a hyphen, you can achieve <b>AND</b> operations, meaning it will search the same line for the information

example: ```*.80-LISTEN```

This will search for ports on 80 wild card <b>AND</b> that they are listeners

<b>Its ok to be more specific or general in the search, give it a couple iterations to nail down what your looking for.</b>

##### Command Argument Example:

1 netstat arguments ex: -an

2 Time In Seconds to sleep between pulling details

3 Search criteria: 80-ESTABLISHED *.80-LISTEN 2000-ESTABLISHED




