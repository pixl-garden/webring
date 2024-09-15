package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Pixl Garden Webring</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        h1 {
            color: #2c3e50;
        }
        a {
            color: #3498db;
        }
    </style>
</head>
<body>
    <h1>Welcome to the Pixl Garden Webring</h1>
    <p>A webring connecting the digital gardens of pixl_garden members.</p>
    <h2>What's a Webring?</h2>
    <p>A webring is a collection of websites linked together in a circular structure, allowing visitors to navigate through related sites.</p>
    <h2>How to Use</h2>
    <p>To navigate the webring, look for these links on member sites:</p>
    <ul>
        <li><a href="/prev?site=CURRENT_SITE_URL">Previous</a></li>
        <li><a href="/">Webring</a> (this page)</li>
        <li><a href="/next?site=CURRENT_SITE_URL">Next</a></li>
    </ul>
    <p>Replace CURRENT_SITE_URL with the URL of the site you're on.</p>
    <h2>Join the Webring</h2>
    <p>To join the webring, please visit our <a href="https://github.com/notalim/pixl_garden_webring">GitHub repository</a> for instructions.</p>
</body>
</html>
	`)
}