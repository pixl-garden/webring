# Pixl Garden Webring

A webring connecting the digital gardens of pixl_garden members.

## What's a Webring?

A webring is a collection of websites linked together in a circular structure, allowing visitors to navigate through related sites.

## Members

<!-- MEMBERS_START -->
<!-- This section will be automatically updated -->
<!-- MEMBERS_END -->

How to use the webring:

1. Add your site to the list of members by submitting a pull request.
2. Add the webring navigation to your site (example):

```html
<div class="webring">
    <a href="https://pg-webring.vercel.app/prev?site=YOUR_SITE_URL">
        Previous</a
    >
    <a href="https://pg-webring.vercel.app">Pixl Garden Webring</a>
    <a href="https://pg-webring.vercel.app/next?site=YOUR_SITE_URL">Next</a>
</div>
```

Here's the example of `CURL` request to become a member:

```bash
curl -X POST https://pg-webring.vercel.app/api/members -H "Content-Type: application/json" -d '{
           "githubUsername": "github-username",
           "name": "Your Name",
           "website": "yourwebsite.com"
         }'

```


by [@notalim](https://github.com/notalim)
