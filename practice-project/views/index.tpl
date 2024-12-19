<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat API</title>
    <link rel="stylesheet" href="/static/css/styles.css">
</head>
<body>
    <div class="cat-container">
        <h1>Here is your random cat!</h1>
        <img src="{{.CatImage}}" alt="Random Cat" class="cat-image">
        
        <div class="cat-info">
            <h2>Breed: {{.CatBreedName}}</h2>
            <p><strong>Origin:</strong> {{.CatBreedOrigin}}</p>
            <p><strong>Description:</strong> {{.CatBreedDescription}}</p>
            <p><a href="{{.CatBreedURL}}" target="_blank">Learn more about this breed</a></p>
        </div>

        <button onclick="window.location.reload()">Get Another Cat</button>
    </div>
    <script src="/static/js/scripts.js"></script>
</body>
</html>
