document.addEventListener("DOMContentLoaded", function() {
    // Function to fetch votes
    async function fetchVotes() {
        try {
            const response = await fetch('/getVotes'); // Call your API endpoint
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const votes = await response.json();
            displayVotes(votes);
        } catch (error) {
            console.error('Error fetching votes:', error);
        }
    }

    // Function to display votes
    function displayVotes(votes) {
        const container = document.querySelector('.myvotes-container');
        container.innerHTML = ''; // Clear existing content

        if (votes.length === 0) {
            showNoVotesMessage();
        } else {
            // Clear any previous "no votes" message
            const noVotesMessage = document.querySelector('.no-votes');
            if (noVotesMessage) {
                noVotesMessage.remove();
            }

            votes.forEach(vote => {
                const voteItem = document.createElement('div');
                voteItem.classList.add('vote-item');

                const img = document.createElement('img');
                img.src = vote.image.url;
                img.alt = 'Image';
                img.classList.add('vote-image');

                const voteStatus = document.createElement('div');
                voteStatus.classList.add('vote-status');
                voteStatus.textContent = vote.value === 1 ? '👍' : '👎';

                voteItem.appendChild(img);
                voteItem.appendChild(voteStatus);

                container.appendChild(voteItem);
            });
        }
    }

    // Function to show a "no votes" message
    function showNoVotesMessage() {
        let messageDiv = document.createElement('div');
        messageDiv.className = 'no-votes';
        messageDiv.textContent = "You haven't voted anything yet";
        document.querySelector('.myvotes-container').appendChild(messageDiv);
    }

    // Fetch votes when the page loads
    fetchVotes();
});
