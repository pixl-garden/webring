package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/pixl-garden/webring/pkg/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler called")
	log.Printf("Go version: %s", runtime.Version())
	
	log.Println("FIREBASE_DATABASE_URL:", os.Getenv("FIREBASE_DATABASE_URL"))
	log.Println("FIREBASE_CREDENTIALS length:", len(os.Getenv("FIREBASE_CREDENTIALS")))

	db := database.GetDBClient()
	if db == nil {
		log.Println("Database client is nil")
	} else {
		log.Println("Database client initialized successfully")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>pixl garden</title>
</head>
<body>
    <h1>pixl garden</h1>
    <h2>code collective</h2>
    <h3>members list</h3>
    <ul id="members-list">
        <!-- Members will be populated here -->
    </ul>

    <script>
        // Fetch members from the API and populate the list
        fetch('/api/members')
            .then(response => response.json())
            .then(data => {
                const membersList = document.getElementById('members-list');
                data.forEach(member => {
                    const listItem = document.createElement('li');
                    listItem.innerHTML = '<a href="' + member.website + '">' + member.name + '</a>';
                    membersList.appendChild(listItem);
                });
            })
            .catch(error => console.error('Error fetching members:', error));
    </script>
</body>
</html>
	`)
}