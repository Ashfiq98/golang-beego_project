function showVotingView() {
    // Hide all views
    document.querySelectorAll('.view-content').forEach(view => view.classList.add('hidden'));

    // Show the voting view
    document.getElementById('voting-view').classList.remove('hidden');

    // Update navigation button styles
    const navButtons = document.querySelectorAll('.nav-btn');
    navButtons.forEach(button => {
        button.classList.remove('text-red-500'); // Remove active class
        button.classList.add('text-gray-500');  // Add inactive class

        // Highlight the voting button
        if (button.getAttribute('data-view') === 'voting') {
            button.classList.remove('text-gray-500');
            button.classList.add('text-red-500');
        }
    });
}
let breedsData = [];
let currentImageIndex = 0;
let currentImages = [];
let favorites = JSON.parse(localStorage.getItem('catFavorites')) || [];
let autoSlideInterval;

// Fetch breeds data when the page loads
fetch('/breeds')
    .then(response => response.json())
    .then(data => {
        breedsData = data;
        const breedList = document.getElementById('breeds-list');

        // Populate the datalist with breed names
        data.forEach(breed => {
            const option = document.createElement('option');
            option.value = breed.name;
            breedList.appendChild(option);
        });

        if (data.length > 0) {
            selectBreed(data[0]); // Automatically select the first breed
        }
    })
    .catch(error => console.error('Error fetching breeds:', error));

// Handle breed search
document.getElementById('breed-search').addEventListener('input', function () {
    const selectedBreed = breedsData.find(breed =>
        breed.name.toLowerCase() === this.value.toLowerCase()
    );
    if (selectedBreed) {
        selectBreed(selectedBreed);
    }
});

// Update image slider with dots
function updateImageSlider() {
    const slider = document.getElementById('slider-images');
    const dotsContainer = document.getElementById('slider-dots');
    // console.log(currentImages[0]);
    if (currentImages.length === 0) return;

    // Update the main slider image
    slider.innerHTML = `
        <img src="${currentImages[currentImageIndex]}" 
             alt="Cat" 
             class="w-full h-full object-cover">
    `;

    // Update slider dots
    dotsContainer.innerHTML = currentImages.map((_, index) => `
        <button class="w-3 h-3 rounded-full transition-colors ${index === currentImageIndex ? 'bg-white' : 'bg-white/50'}" data-index="${index}"></button>
    `).join('');

    // Add click handlers to dots
    dotsContainer.querySelectorAll('button').forEach(dot => {
        dot.addEventListener('click', () => {
            currentImageIndex = parseInt(dot.dataset.index);
            updateImageSlider();
            resetAutoSlide();
        });
    });
}

// Slide to the next image
function nextSlide() {
    currentImageIndex = (currentImageIndex + 1) % currentImages.length;
    updateImageSlider();
}

// Reset the auto-slide interval
function resetAutoSlide() {
    clearInterval(autoSlideInterval);
    autoSlideInterval = setInterval(nextSlide, 4000); // Slide every 4 seconds
}

// Fetch and display breed details and images
function selectBreed(breed) {
    fetch(`/breed-images/${breed.id}`)
        .then(response => response.json())
        .then(data => {
            // console.log(data)
            currentImages = data.map(image => image.url);
            // console.log("here....")
            currentImageIndex = 0;
            updateImageSlider();
            resetAutoSlide();

            // Update breed information
            document.getElementById('breed-name').textContent = breed.name;
            document.getElementById('breed-origin').textContent = `(${breed.origin})`;
            document.getElementById('breed-id').textContent = breed.id;
            document.getElementById('breed-description').textContent = breed.description;
            document.getElementById('wiki-link').href = breed.wikipedia_url;

            document.getElementById('breed-info').classList.remove('hidden');
        })
        .catch(error => console.error('Error fetching breed images:', error));
}

// Stop auto-sliding when the page is hidden
document.addEventListener('visibilitychange', () => {
    if (document.hidden) {
        clearInterval(autoSlideInterval);
    } else {
        resetAutoSlide();
    }
});

// Favorites functionality
// function addToFavorites(imageUrl) {
//     if (!favorites.includes(imageUrl)) {
//         favorites.push(imageUrl);
//         localStorage.setItem('catFavorites', JSON.stringify(favorites));
//         updateFavoritesView();
//     }
// }

function updateFavoritesView() {
    const favoritesGrid = document.getElementById('favorites-grid');
    const noFavoritesMessage = document.getElementById('no-favorites-message');

    if (favorites.length === 0) {
        noFavoritesMessage.classList.remove('hidden');
        favoritesGrid.innerHTML = '';
    } else {
        noFavoritesMessage.classList.add('hidden');
        favoritesGrid.innerHTML = favorites.map((imageUrl, index) => `
            <div class="relative group">
                <img src="${imageUrl}" 
                     alt="Favorite Cat" 
                     class="w-full h-48 object-cover rounded-lg">
                <button onclick="removeFavorite(${index})"
                        class="absolute top-2 right-2 bg-white/80 rounded-full p-2 opacity-0 group-hover:opacity-100 transition-opacity">
                    <i class="fas fa-times text-red-500"></i>
                </button>
            </div>
        `).join('');
    }
}

function removeFavorite(index) {
    favorites.splice(index, 1);
    localStorage.setItem('catFavorites', JSON.stringify(favorites));
    updateFavoritesView();
}

// View switching
$('.nav-btn').click(function () {
    const viewToShow = $(this).data('view');
    $('.nav-btn').removeClass('text-red-500').addClass('text-gray-500');
    $(this).removeClass('text-gray-500').addClass('text-red-500');
    $('.view-content').addClass('hidden');
    $(`#${viewToShow}-view`).removeClass('hidden');

    if (viewToShow === 'favs') {
        updateFavoritesView();
    }
});

// Favorite button click handler
$('.fav-btn').click(function () {
    const heartIcon = $(this).find('.fa-heart');
    heartIcon.toggleClass('text-red-500 text-gray-600');

    // if (currentImages.length > 0) {
    //     addToFavorites(currentImages[currentImageIndex]);
    // }
});

// Thumbs up/down handlers
$('.fa-thumbs-up').parent().click(function () {
    $(this).find('.fa-thumbs-up').toggleClass('text-green-500 text-gray-600');
});

$('.fa-thumbs-down').parent().click(function () {
    $(this).find('.fa-thumbs-down').toggleClass('text-red-500 text-gray-600');
});

// Ensure the favorites view updates when the page loads
document.addEventListener('DOMContentLoaded', () => {
    updateFavoritesView();
});


document.addEventListener("DOMContentLoaded", function () {
    const subId = "test123";  // Your static sub_id value

    // Function to fetch cat data and get the image_id
    async function getCatData(data_type = 'id') {
        try {
            const response = await fetch('/getcatdata');
            if (response.ok) {
                const data = await response.json();
                console.log(data)
                // const imageId = data[0].id;  // Assuming `image_id` is in the response
                // console.log(imageId)
                // const catImageElement = document.getElementById('cat-image');
                // catImageElement.src = data.catImage || 'default-image.jpg';  // Fallback image
                if (data_type != 'id')
                    return data[0].url;
                else
                    return data[0].id;
            } else {
                console.error('Failed to fetch cat data');
            }
        } catch (error) {
            console.error('Error fetching cat data:', error);
        }
    }
    async function addToFavorites(imageId) {  // Change parameter name to reflect what it is
        const subId = "test123";
        const response = await fetch('/favourites', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json', // Ensures the backend knows you're sending JSON
            },
            body: JSON.stringify({
                image_id: imageId, // Data in key-value pairs
                sub_id: subId,     // Your fixed sub_id
            }),
        });
        

        const data = await response.json();

        if (response.ok) {
            console.log("Favourite created successfully:", data);
        } else {
            console.error("Failed to create favourite:", data);
        }
    }


    // Function to handle upvote
    async function voteUp(imageId) {
        try {
            const response = await fetch('/vote/up', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `image_id=${imageId}&sub_id=${subId}`,
            });

            if (response.ok) {
                console.log('Vote Up successful');
                // After successful vote, update the vote history
                // getVoteHistory();
            } else {
                console.error('Failed to vote up');
            }
        } catch (error) {
            console.error('Error voting up:', error);
        }
    }

    // Function to handle downvote
    async function voteDown(imageId) {
        try {
            const response = await fetch('vote/down', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `image_id=${imageId}&sub_id=${subId}`,
            });

            if (response.ok) {
                console.log('Vote Down successful');
                // After successful vote, update the vote history
                // getVoteHistory();
            } else {
                console.error('Failed to vote down');
            }
        } catch (error) {
            console.error('Error voting down:', error);
        }
    }

    // Function to fetch and log vote history
    // async function getVoteHistory() {
    //     try {
    //         const response = await fetch(`https://api.thecatapi.com/v1/votes?sub_id=${subId}`, {
    //             headers: {
    //                 'x-api-key': 'live_8Vq87uY7jXkcqmqwhODWVdzEp9iUzbog1G0hxJgh6gphgTP9sjK23Pbnir5Xl5JY',  // Replace with your actual API key
    //             },
    //         });

    //         if (response.ok) {
    //             const data = await response.json();
    //             console.log('Vote history:', data);

    //             // Check if history is returned in the expected format
    //             if (Array.isArray(data) && data.length > 0) {
    //                 // Log the vote history or process it as needed
    //                 data.forEach(vote => {
    //                     console.log('Vote:', vote);
    //                     // Here, you could update the UI if you want to show the history
    //                 });
    //             } else {
    //                 console.log('No votes found in history');
    //             }
    //         } else {
    //             console.error('Failed to fetch vote history');
    //         }
    //     } catch (error) {
    //         console.error('Error fetching vote history:', error);
    //     }
    // }

    // Bind events to upvote and downvote buttons
    document.getElementById('upvote-btn').addEventListener('click', function (event) {
        event.preventDefault();  // Prevent page reload
        getCatData().then(imageId => {
            voteUp(imageId);
        });
    });

    document.getElementById('downvote-btn').addEventListener('click', function (event) {
        event.preventDefault();  // Prevent page reload
        getCatData().then(imageId => {
            voteDown(imageId);
        });
    });
    document.getElementById('favourite-btn').addEventListener('click', function (event) {
        event.preventDefault();  // Prevent page reload
        console.log("Pressed favourite...")
        getCatData().then(imageId => {
            addToFavorites(imageId);  // Call the function to add the image to favourites
            // console.log(imageUrl);
        });
    });
    // Optionally, you can fetch and display the vote history when the page loads
    // getVoteHistory();
});
