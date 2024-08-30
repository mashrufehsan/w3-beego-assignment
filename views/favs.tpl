<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat API | GO Assignment</title>
    <link rel="shortcut icon" href="static/img/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" href="static/css/style.css">
    <link rel="stylesheet" href="static/css/favStyle.css">
</head>

<body>
    <div class="container">
        <div class="navbar">
            <div id="voting" class="nav-item">
                <a href="/">
                    <img src="static/img/two-way.png" alt="">
                    Voting
                </a>
            </div>
            <div id="breeds" class="nav-item">
                <a href="/tab_breeds">
                    <img src="static/img/loupe.png" alt="">
                    Breeds
                </a>
            </div>
            <div id="favs" class="nav-item">
                <a class="active-nav" href="/tab_favs">
                    <img src="static/img/heart.png" alt="">
                    Favs
                </a>
            </div>
            <div id="myVotes" class="nav-item">
                <a href="/tab_myVotes">
                    <img src="static/img/positive-vote.png" alt="">
                    My Votes
                </a>
            </div>
        </div>
        

        <div class="favs-container">
        </div>

    </div>

    <script src="static/js/favScript.js"></script>
</body>

</html>