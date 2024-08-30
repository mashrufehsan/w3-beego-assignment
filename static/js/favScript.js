document.addEventListener("DOMContentLoaded", function () {
    // Function to fetch favorites
    async function fetchFavorites() {
        try {
            const response = await fetch('/getFavourites'); // Call your API endpoint
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const favorites = await response.json();
            displayFavorites(favorites);
        } catch (error) {
            console.error('Error fetching favorites:', error);
        }
    }

    // Function to display favorites
    function displayFavorites(favorites) {
        const container = document.querySelector('.favs-container');
        container.innerHTML = ''; // Clear existing content

        favorites.forEach(favorite => {
            const favItem = document.createElement('div');
            favItem.classList.add('fav-item');

            const img = document.createElement('img');
            img.src = favorite.image.url;
            img.alt = 'Favorite Image';
            img.classList.add('fav-image');

            const heartIcon = document.createElement('div');
            heartIcon.classList.add('fav-icon');

            const heartEmoji = document.createElement('span');
            heartEmoji.textContent = '💔'; // Heart emoji
            heartEmoji.title = 'Remove from Favorites'; // Tooltip text
            heartIcon.appendChild(heartEmoji);

            // Attach event listener to heart emoji
            heartIcon.addEventListener('click', async () => {
                try {
                    const response = await fetch(`/deleteAFavourite/${favorite.id}`, { // Correct endpoint with ID
                        method: 'DELETE',
                    });

                    if (response.ok) {
                        showMessage('Removed from favorites');
                        fetchFavorites(); // Refresh the list after deletion
                    } else {
                        console.error('Error removing favorite:', response.statusText);
                    }
                } catch (error) {
                    console.error('Error during delete request:', error);
                }
            });

            favItem.appendChild(img);
            favItem.appendChild(heartIcon);

            container.appendChild(favItem);
        });
    }

    // Function to show a temporary message
    function showMessage(message) {
        let messageDiv = document.querySelector('.message');
        if (!messageDiv) {
            messageDiv = document.createElement('div');
            messageDiv.className = 'message';
            document.querySelector('.favs-container').appendChild(messageDiv);
        }
        messageDiv.textContent = message;

        setTimeout(() => {
            if (messageDiv.parentNode) {
                messageDiv.parentNode.removeChild(messageDiv);
            }
        }, 2000); // Remove the message after 2 seconds
    }

    // Fetch favorites when the page loads
    fetchFavorites();
});
