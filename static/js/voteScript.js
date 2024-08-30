let currentImageId = null; // Global variable to store the current image ID

async function loadCatImage() {
    try {
        const response = await fetch('/getACat');
        const data = await response.json();
        const imageUrl = data.url;
        currentImageId = data.id; // Store image ID in global variable

        const img = document.createElement('img');
        img.src = imageUrl;
        img.alt = 'Random Cat Image';
        img.dataset.id = currentImageId; // Store image ID in a data attribute

        const container = document.getElementById('cat-container');
        container.innerHTML = ''; // Clear previous images
        container.appendChild(img);
    } catch (error) {
        console.error('Error fetching cat image:', error);
    }
}

// General function to handle post requests
async function postRequest(url, body, successMessage) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        });

        if (response.ok) {
            showMessage(successMessage);
            setTimeout(loadCatImage, 1000); // Load new image after 1 second
        } else {
            console.error(`Error in ${url}:`, response.statusText);
        }
    } catch (error) {
        console.error(`Error sending request to ${url}:`, error);
    }
}

function attachEventListeners() {
    // Attach event listeners only once
    const heartIcon = document.querySelector('.heart-icon');
    heartIcon.addEventListener('click', () => {
        postRequest('/createAFavourite', {
            image_id: currentImageId
        }, 'Added to favorites');
    });

    const thumbsUpIcon = document.querySelector('.thumbs-up-icon');
    thumbsUpIcon.addEventListener('click', () => {
        postRequest('/vote', {
            image_id: currentImageId,
            value: 1
        }, 'Up voted!');
    });

    const thumbsDownIcon = document.querySelector('.thumbs-down-icon');
    thumbsDownIcon.addEventListener('click', () => {
        postRequest('/vote', {
            image_id: currentImageId,
            value: -1
        }, 'Down voted!');
    });
}

function showMessage(message) {
    const container = document.getElementById('cat-container');
    const messageDiv = document.createElement('div');
    messageDiv.className = 'message';
    messageDiv.textContent = message;
    container.appendChild(messageDiv);

    setTimeout(() => {
        // Check if the messageDiv is still a child of the container before removing it
        if (container.contains(messageDiv)) {
            container.removeChild(messageDiv);
        }
    }, 2000); // Remove the message after 2 seconds
}


// Load the cat image when the page loads and attach event listeners
window.onload = () => {
    loadCatImage();
    attachEventListeners(); // Attach event listeners once
};
