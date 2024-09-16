# pixl garden webring

a webring connecting the digital gardens of pixl_garden members.

## what's a webring?

a webring is a collection of websites linked together in a circular structure, allowing visitors to navigate through related sites.

## members

<!-- MEMBERS_START -->
<!-- this section will be automatically updated -->
<!-- MEMBERS_END -->

how to use the webring:

1. add your site to the list of members by submitting a pull request.
2. add the webring navigation to your site (example):

```html
<div class="webring">
    <a href="https://pg-webring.vercel.app/prev?site=YOUR_SITE_URL">
        prev</a
    >
    <a href="https://pg-webring.vercel.app">pixl garden webring</a>
    <a href="https://pg-webring.vercel.app/next?site=YOUR_SITE_URL">next</a>
</div>
```

> add styling however you like.

here's the example of `curl` request to become a member:

```bash
curl -X POST https://pg-webring.vercel.app/api/members -H "Content-Type: application/json" -d '{
"githubUsername": "github-username",
"name": "your name",
    "website": "yourwebsite.com"
}'
```


by [@notalim](https://github.com/notalim)
