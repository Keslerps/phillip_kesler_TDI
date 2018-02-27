# phillip_kesler_TDI
Internship Application challenge by Phillip Kesler for TDI


<h2><strong>If you do not work for TDI or are not affiliated with the TDI internship application process, please do not issue a pull request for this code. 
Thank you!</strong></h2>

<b>Files Found in this repository include: </b>
<h3>
main.go
</h3>
<p>
Main entry point for the program. Run this file to start the server up on port 8080.
</p>
<h3>
API.go
</h3>
<p>
Contains structs for all the .json data we will be collecting. As the .json data is tiered for each reddit.com/r/subreddit.json request, we must construct multiple structs to break the data up:
</p>
<p>
  Reddit_API exists to grab the first two values returned. The first value should be the string "listing", and the second should be the rest of the data of the Front page. The front page data is stored in the topData struct.
  </p>
  <p>
  TopData simply stores an array of childs, or posts, called Children. Other fields present in the .json are ignored, as they are not of interest to us.
  </p>
  <p>
  Child struct is the struct of a post. Each post has data and a kind field. The kind field for a post should always be the string "t3", though this will be of little to no use for us and should only be checked for debugging purposes, or for sanity checks.
  </p>
  <p>
  data struct contains the actual data that we care about for an individual post. It should be of note that there contains much more data for a single post than what is covered, but that data is of no interest and is thus excluded. It would however be trivial to add more. For a definition of each field, check the comments of the API.go file.
</p>
<h3>
HTTP_Handler.go
</h3>
<p>
There are two functions in the HTTP_Handler.go file, startUp() and handler(). The startUp() function will be called by main to construct the mux router (from github/gorilla/mux) and specify what request it should handle. Requests that are currently handled by the server are of the form localhost:8080/redditer/:{search}, where search is some user defined parameter. This function also returns the router for any future needs (such as testing).

The other function, handler, is the function that handles the specific request. It will issue a GET request for the .json data for the front page of whatever subreddit is specified in the subreddit constant, read in the data, and then UnMarshall it into the structs, before calling output for the data format and printing. Error checking is provided, but currently the program will only exit if an error is encountered.
</p>

<h3>
Output.go
</h3>
<p>
Formats and prints the data received by HTTP_Handler. Output will post the search term and date to the response writer, and then list out the data for each post that has the search term somewhere in the title. If there are no posts that match, the user is currently told so.

Format prints out the posts data. This is made to reduce the amount of code in the output function, and for code style purposes. Examples of the output are listed below in the examples section.
</p>
