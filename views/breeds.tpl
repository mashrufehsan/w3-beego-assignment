<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat API | GO Assignment</title>
    <link rel="shortcut icon" href="static/img/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" href="static/css/style.css">
    <link rel="stylesheet" href="static/css/breedStyle.css">
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
                <a class="active-nav" href="/tab_breeds">
                    <img src="static/img/loupe.png" alt="">
                    Breeds
                </a>
            </div>
            <div id="favs" class="nav-item">
                <a href="/tab_favs">
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
        

        <div class="breeds-container">
            <div class="dropdown">
                <div class="dropdown-container">
                    <input type="text" id="breedInput" placeholder="Please select" autocomplete="off">
                    <span id="clearDropdown" class="clear-dropdown">&times;</span>
                    <ul id="breedList" class="dropdown-list"></ul>
                </div>
            </div>
        
            <div class="carousel">
                <div class="carousel-images" id="carouselImages"></div>
                <div class="carousel-controls">
                    <button id="prevButton" class="carousel-button">&lt;</button>
                    <button id="nextButton" class="carousel-button">&gt;</button>
                </div>
                <span id="imageCounter" class="carousel-counter">1/1</span>
            </div>
        
            <div id="breedDetails" class="breed-details"></div>
        </div>
        
        
    </div>

    <script src="static/js/breedScript.js"></script>
</body>

</html>