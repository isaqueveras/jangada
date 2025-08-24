package newapp

const tmplWebLayoutsPage string = `package layouts

templ Page(title string) {
  <!DOCTYPE html>
  <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>{title} - Jangada Framework</title>
      <link rel="icon" type="image/png" href="public/background.png" />
      <script src="https://unpkg.com/htmx.org"></script>
      <script src="https://cdn.tailwindcss.com"></script>
			<style>
        body {
          background-image: url("public/background.png");
          background-repeat: repeat;
          background-size: 220px 220px;
          background-position: top left;
        }
      </style>
    </head>
    <body>
      { children... }
    </body>
  </html>
}
`
