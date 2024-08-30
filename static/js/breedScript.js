document.addEventListener("DOMContentLoaded", function() {
    const breedInput = document.getElementById("breedInput");
    const clearDropdown = document.getElementById("clearDropdown");
    const breedList = document.getElementById("breedList");
    const breedDetailsDiv = document.getElementById("breedDetails");
    const carouselImagesDiv = document.getElementById("carouselImages");
    const prevButton = document.getElementById("prevButton");
    const nextButton = document.getElementById("nextButton");
    const imageCounter = document.getElementById("imageCounter");

    let breeds = [];
    let images = [];
    let currentImageIndex = 0;

    // Show/hide the dropdown items when clicking the input field
    breedInput.addEventListener('click', function() {
        if (breedInput.value === '' && breedList.children.length === 0) {
            return; // Do nothing if input is empty and no items in the list
        }
        
        if (breedList.style.display === 'none' || breedList.children.length === 0) {
            populateDropdownList(breeds);
            breedList.style.display = 'block'; // Show dropdown if hidden
        } else {
            breedList.style.display = 'none'; // Hide dropdown if visible
        }
    });

    // Show clear button when input is not empty
    breedInput.addEventListener('input', function() {
        const query = breedInput.value.toLowerCase();
        if (query) {
            clearDropdown.style.display = 'block';
            filterBreeds(query);
        } else {
            clearDropdown.style.display = 'none';
            breedList.style.display = 'none';
            breedInput.placeholder = 'Please select'; // Set placeholder when cleared
        }
    });

    // Clear input and hide the dropdown list
    clearDropdown.addEventListener('click', function() {
        breedInput.value = '';
        clearDropdown.style.display = 'none';
        breedList.style.display = 'none';
        breedInput.placeholder = 'Please select'; // Set placeholder when cleared
    });

    // Function to filter breeds based on input
    function filterBreeds(query) {
        breedList.innerHTML = ''; // Clear previous results
        const filteredBreeds = breeds.filter(breed => breed.name.toLowerCase().includes(query));
        if (filteredBreeds.length > 0) {
            populateDropdownList(filteredBreeds);
            breedList.style.display = 'block'; // Show the dropdown list
        } else {
            breedList.style.display = 'none'; // Hide the list if no matches
        }
    }

    // Populate the dropdown list with breed options
    function populateDropdownList(breedArray) {
        breedList.innerHTML = ''; // Clear the list
        breedArray.forEach(breed => {
            const li = document.createElement("li");
            li.textContent = breed.name;
            li.dataset.id = breed.id; // Store breed ID for later use
            breedList.appendChild(li);
        });
    }

    // Handle clicking on a dropdown item
    breedList.addEventListener('click', function(e) {
        if (e.target.tagName === 'LI') {
            const selectedId = e.target.dataset.id;
            breedInput.value = e.target.textContent;
            breedList.style.display = 'none'; // Hide the dropdown list
            clearDropdown.style.display = 'block'; // Show the clear button
            fetchBreedDetails(selectedId);
            fetchBreedImages(selectedId);
        }
    });

    // Function to fetch breeds
    async function fetchBreeds() {
        try {
            const response = await fetch('/breeds'); // Call your Beego endpoint
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            breeds = await response.json();
            if (Array.isArray(breeds)) {
                populateDropdown(breeds);
            } else {
                console.error('Expected an array of breeds.');
            }
        } catch (error) {
            console.error('Error fetching breeds:', error);
        }
    }

    // Populate dropdown with breeds (hidden initially)
    function populateDropdown(breeds) {
        // Automatically select the first breed and fetch details and images
        if (breeds.length > 0) {
            breedInput.value = breeds[0].name;
            clearDropdown.style.display = 'block';
            fetchBreedDetails(breeds[0].id);
            fetchBreedImages(breeds[0].id);
        }
    }

    // Function to fetch breed details
    async function fetchBreedDetails(breedId) {
        try {
            const response = await fetch(`/breeds/${breedId}`); // Fetch details for the selected breed
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const breedDetails = await response.json();
            displayBreedDetails(breedDetails);
        } catch (error) {
            console.error('Error fetching breed details:', error);
            breedDetailsDiv.innerHTML = '<p>Error fetching breed details. Please try again later.</p>';
        }
    }

    // Function to fetch breed images
    async function fetchBreedImages(breedId) {
        try {
            const response = await fetch(`/breeds/images/${breedId}`); // Fetch images for the selected breed
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const breedImages = await response.json();
            images = breedImages;
            displayImages();
        } catch (error) {
            console.error('Error fetching breed images:', error);
            carouselImagesDiv.innerHTML = '<p>Error fetching breed images. Please try again later.</p>';
        }
    }

    // Function to display breed details
    function displayBreedDetails(breedDetails) {
        breedDetailsDiv.innerHTML = `
            <div class="breed-title-container">
                <h3>${breedDetails.name}</h3>
                <h4 class="breed-details-color-gray">(${breedDetails.origin})</h4>
                <p class="breed-details-color-gray breed-details-font-italic">${breedDetails.id}</p>
            </div>
            <p class="breed-details-color-gray">${breedDetails.description}</p>
            <a class="breed-details-wiki" href="${breedDetails.wikipedia_url}" target="_blank">WIKIPEDIA</a>
        `;
    }

    // Function to display images in the carousel
    function displayImages() {
        carouselImagesDiv.innerHTML = ''; // Clear previous images
        images.forEach((image, index) => {
            const img = document.createElement("img");
            img.src = image;
            img.alt = `Image ${index + 1}`;
            img.style.width = '100%'; // Ensure images fit the container
            img.style.flexShrink = '0'; // Prevent images from shrinking
            carouselImagesDiv.appendChild(img);
        });
        currentImageIndex = 0; // Reset index to 0 for the new breed
        updateCarousel();
    }

    // Function to update carousel display
    function updateCarousel() {
        if (images.length > 0) {
            const imgElements = carouselImagesDiv.getElementsByTagName("img");
            carouselImagesDiv.style.transform = `translateX(-${currentImageIndex * 100}%)`;
            imageCounter.textContent = `${currentImageIndex + 1}/${images.length}`; // Update counter

            prevButton.disabled = currentImageIndex === 0;
            nextButton.disabled = currentImageIndex === imgElements.length - 1;
        } else {
            imageCounter.textContent = '0/0'; // No images
        }
    }

    // Event listeners for carousel controls
    prevButton.addEventListener('click', function() {
        if (currentImageIndex > 0) {
            currentImageIndex--;
            updateCarousel();
        }
    });

    nextButton.addEventListener('click', function() {
        if (currentImageIndex < images.length - 1) {
            currentImageIndex++;
            updateCarousel();
        }
    });

    // Fetch breeds when the page loads
    fetchBreeds();
});
